import { User } from "src/app/interfaces/UserInterface";

describe('SignupService', () => {
  let newUser: User = {
    username: "erob",
    email: "evan.robinson@ufl.edu",
    password: "test"
  }
  it("should create a user", () => {
    cy.request({
      method: 'POST',
      url: 'http://localhost:4201/signup',
      body: newUser,
    }).as('new-user')
    cy.get('@new-user').its('status').should('eq', 200)
  })
})