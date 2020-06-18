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

import { defineSmokeTest } from '../utils'

const applicationCreate = defineSmokeTest('should be able to create application', () => {
  cy.visit(Cypress.config('consoleRootPath'))

  const user = {
    user_id: 'app-create-test-user',
    password: '123456QWERTY!',
    email: 'app-create-test-user@example.com',
    name: 'Test App',
  }

  // Register
  cy.findByRole('button', { name: 'Create an account' }).click()
  cy.findByLabelText('User ID').type(user.user_id)
  cy.findByLabelText('Name').type(user.name)
  cy.findByLabelText('Email').type(user.email)
  cy.findByLabelText('Password').type(user.password)
  cy.findByLabelText('Confirm password').type(user.password)
  cy.findByRole('button', { name: 'Register' }).click()

  cy.findByTestId('notification')
    .should('be.visible')
    .should('contain', 'You have successfully registered and can login now')

  // Login
  cy.visit(Cypress.config('consoleRootPath'))
  cy.findByLabelText('User ID').type(user.user_id)
  cy.findByLabelText('Password').type(`${user.password}`)
  cy.findByRole('button', { name: 'Login' }).click()

  // Create application.
  const application = {
    application_id: 'app-create-test-app',
    name: 'Application Create Test',
    description: 'Application used in smoke test to verify application creation',
  }
  cy.get('header').within(() => {
    cy.findByRole('link', { name: /Applications/ }).click()
  })
  cy.findByRole('link', { name: /Add application/ }).click()
  cy.findByLabelText('Application ID').type(application.application_id)
  cy.findByLabelText('Application name').type(application.name)
  cy.findByLabelText('Description').type(application.description)
  cy.findByRole('button', { name: 'Create application' }).click()

  cy.location('pathname').should(
    'eq',
    `${Cypress.config('consoleRootPath')}/applications/${application.application_id}`,
  )
})

export default [applicationCreate]
