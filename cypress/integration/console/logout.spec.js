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

const logout = userName => {
  cy.findByTestId('profile-dropdown')
    .should('contain', userName)
    .as('profileDropdown')

  cy.get('@profileDropdown').click()
  cy.get('@profileDropdown')
    .findByText('Logout')
    .click()
}

describe('Console logout', () => {
  before(() => {
    cy.dropAndSeedDatabase()
  })

  it('succeeds when logged in properly', () => {
    const user = {
      ids: { user_id: 'test-user' },
      primary_email_address: 'test-user@example.com',
      password: 'ABCDefg123!',
      password_confirm: 'ABCDefg123!',
    }
    cy.registerUser(user)
    cy.loginConsole({ user_id: user.ids.user_id, password: user.password })
    cy.visit(Cypress.config('consoleRootPath'))

    logout(user.ids.user_id)

    cy.location('pathname').should('eq', `${Cypress.config('oauthRootPath')}/login`)
  })

  it('obtains a new CSRF token and succeeds when CSRF token not present', () => {
    const user = {
      ids: { user_id: 'test-user2' },
      primary_email_address: 'test-user2@example.com',
      password: 'ABCDefg123!',
      password_confirm: 'ABCDefg123!',
    }
    cy.server()
    cy.route({
      method: 'POST',
      url: 'http://localhost:8080/console/api/auth/logout',
      onRequest: req => {
        expect(req.request.headers).to.have.property('X-CSRF-Token')
      },
    })

    cy.registerUser(user)
    cy.loginConsole({ user_id: user.ids.user_id, password: user.password })
    cy.visit(Cypress.config('consoleRootPath'))
    cy.clearCookie('_console_csrf')

    logout(user.ids.user_id)

    cy.location('pathname').should('eq', `${Cypress.config('oauthRootPath')}/login`)
  })
})
