// Copyright © 2020 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ttnmage

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/magefile/mage/target"
)

// EndToEnd namespace
type EndToEnd mg.Namespace

var (
	testDatabaseName = "ttn_lorawan_test"
	databaseURI      = fmt.Sprintf("postgresql://root@localhost:26257/%s?sslmode=disable", testDatabaseName)
)

func getTestURL() string {
	switch v := os.Getenv("NODE_ENV"); v {
	case "", "production":
		return "http://localhost:1885"

	case "development":
		return "http://localhost:8080"

	default:
		if mg.Verbose() {
			fmt.Printf("Unknown `NODE_ENV` value `%s`, assuming production mode\n", v)
		}
		return "http://localhost:1885"
	}
}

func (e EndToEnd) runCypress(command string, args ...string) error {
	return runYarnCommand("cypress", append([]string{command, fmt.Sprintf("--config-file=%s", filepath.Join("config", "cypress.json")), "--config", fmt.Sprintf("baseUrl=%s", getTestURL())}, args...)...)
}

func (EndToEnd) prepareDB() error {
	isCI := os.Getenv("CI") == "true"
	dumpExists := pathExists(filepath.Join(".cache", "sqldump.sql"))
	ok, err := target.Dir(
		filepath.Join(".cache", "/sqldump.sql"),
		filepath.Join("pkg", "identityserver", "store"),
	)

	if err != nil {
		return targetError(err)
	}
	if !ok || (isCI && dumpExists) {
		mg.SerialDeps(EndToEnd.DBRestore)
		return nil
	}
	if err = os.MkdirAll(".cache", 0755); err != nil {
		return err
	}
	os.Setenv("TTN_LW_IS_DATABASE_URI", databaseURI)
	mg.SerialDeps(Dev.DBErase, Dev.DBStart, Dev.InitStack, EndToEnd.DBDump)
	return nil
}

func (EndToEnd) prepareBuild() {
	mg.SerialDeps(Js.Deps, Js.Build)
}

// Prepare prepares the server for running end to end tests.
func (EndToEnd) Prepare() {
	mg.Deps(EndToEnd.prepareDB, EndToEnd.prepareBuild)
}

// CypressHeadless runs the Cypress end-to-end tests in the headless mode.
func (e EndToEnd) CypressHeadless() error {
	mg.Deps(Js.Deps, EndToEnd.WaitUntilReady)
	if mg.Verbose() {
		fmt.Println("Running Cypress E2E tests in headless mode")
	}
	return e.runCypress("run")
}

// CypressInteractive runs the Cypress end-to-end tests in interactive mode.
func (e EndToEnd) CypressInteractive() error {
	mg.Deps(Js.Deps, EndToEnd.WaitUntilReady)
	if mg.Verbose() {
		fmt.Println("Running Cypress E2E tests in interactive mode")
	}
	return e.runCypress("open")
}

// StartTestStack starts TTS in end-to-end test configuration.
func (EndToEnd) StartTestStack() error {
	mg.Deps(Js.Build)
	if mg.Verbose() {
		fmt.Println("Starting stack in end-to-end test configuration")
	}
	os.Setenv("TTN_LW_IS_DATABASE_URI", databaseURI)
	return runGoMute("./cmd/ttn-lw-stack", "start")
}

// WaitUntilReady waits until the web endpoints become available. For CI use.
func (EndToEnd) WaitUntilReady() error {
	pingURL := getTestURL() + "/oauth"
	if mg.Verbose() {
		fmt.Printf("Waiting for the stack to be available on %s\n", pingURL)
	}
	for i := 0; i < 100; i++ {
		resp, _ := http.Get(pingURL)
		if resp != nil && resp.StatusCode == 200 {
			return nil
		}
		time.Sleep(1000 * time.Millisecond)
	}
	return errors.New("Could not connect to server")
}

// DBDump performs an SQL database dump of the test database to the .cache folder.
func (EndToEnd) DBDump() error {
	if mg.Verbose() {
		fmt.Println("Saving database dump")
	}
	if err := os.MkdirAll(".cache", 0755); err != nil {
		return err
	}
	return execDockerComposeWithOutput(filepath.Join(".cache", "sqldump.sql"), "exec", "-T", "cockroach", "./cockroach", "dump", testDatabaseName, "--insecure")
}

// DBRestore restores the dev database using a previously generated dump.
func (EndToEnd) DBRestore() error {
	mg.Deps(Js.Deps, Dev.DBStart)
	if mg.Verbose() {
		fmt.Println("Restoring database from dump")
	}
	return sh.Run("node", filepath.Join("tools", "mage", "scripts", "restore-db-dump.js"), "--db", testDatabaseName)
}
