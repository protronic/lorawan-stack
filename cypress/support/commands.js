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

import stringToHash from '../../pkg/webui/lib/string-to-hash'

// Helper function to quickly login to the oauth app programmatically.
Cypress.Commands.add('loginOAuth', credentials => {
  const baseUrl = Cypress.config('baseUrl')

  // Obtain csrf token.
  cy.request({
    method: 'GET',
    url: `${baseUrl}/oauth/login`,
  }).then(({ headers }) => {
    cy.request({
      method: 'POST',
      url: `${baseUrl}/oauth/api/auth/login`,
      body: { user_id: credentials.user_id, password: credentials.password },
      headers: {
        'X-CSRF-Token': headers['x-csrf-token'],
      },
    })
  })
})

// Helper function to quickly login to the console programmatically.
Cypress.Commands.add('loginConsole', credentials => {
  const baseUrl = Cypress.config('baseUrl')

  // Obtain csrf token.
  cy.request({
    method: 'GET',
    url: `${baseUrl}/oauth/login`,
  }).then(({ headers }) => {
    // Login to OAuth provider.
    cy.request({
      method: 'POST',
      url: `${baseUrl}/oauth/api/auth/login`,
      body: { user_id: credentials.user_id, password: credentials.password },
      headers: {
        'X-CSRF-Token': headers['x-csrf-token'],
      },
    })

    // Do OAuth round trip.
    cy.request({
      method: 'GET',
      url: `${baseUrl}/console/login/ttn-stack?next=/`,
    })

    // Obtain access token.
    cy.request({
      method: 'GET',
      url: `${baseUrl}/console/api/auth/token`,
    }).then(resp => {
      window.localStorage.setItem(
        // We store local storage values with a hashed key based on the mount path
        // to prevent clashes with other apps on the same domain.
        `accessToken-${stringToHash('/console')}`,
        JSON.stringify(resp.body),
      )
    })
  })
})

// Helper function to register a new user programmatically.
Cypress.Commands.add('registerUser', user => {
  const baseUrl = Cypress.config('baseUrl')

  // Obtain csrf token.
  cy.request({
    method: 'GET',
    url: `${baseUrl}/oauth/login`,
  }).then(({ headers }) => {
    // Register user.
    cy.request({
      method: 'POST',
      url: `${baseUrl}/api/v3/users`,
      body: { user },
      headers: {
        'X-CSRF-Token': headers['x-csrf-token'],
      },
    })
  })
})

// Helper function to quickly seed the database to a fresh state using a
// previously generated sql dump.
Cypress.Commands.add('dropAndSeedDatabase', () => {
  cy.exec('node tools/mage/scripts/restore-db-dump.js')
    .its('code')
    .should('eq', 0)
})

// Selectors

const getFieldDescriptorByLabel = label => {
  cy.findByLabelText(label).as('field')
  return cy
    .get('@field')
    .invoke('attr', 'aria-describedby')
    .then(describedBy => {
      return cy.get(`[id=${describedBy}]`)
    })
}

// Helper function to select field error.
Cypress.Commands.add('findErrorByLabelText', label => {
  getFieldDescriptorByLabel(label).as('error')

  cy.get('@error')
    .children()
    .first()
    .should('contain', 'error')
    .and('be.visible')

  return cy.get('@error')
})

// Helper function to select field warning.
Cypress.Commands.add('findWarningByLabelText', label => {
  getFieldDescriptorByLabel(label).as('warning')

  cy.get('@warning')
    .children()
    .first()
    .should('contain', 'warning')
    .and('be.visible')

  return cy.get('@warning')
})

// Helper function to select field description.
Cypress.Commands.add('findDescriptionByLabelText', label => {
  return getFieldDescriptorByLabel(label)
})
