import { User } from "src/app/interfaces/UserInterface";

describe('SignupService', () => {
  let newUser: User = {
    username: "erob",
    email: "evan.robinson@ufl.edu",
    password: "test"
  }
  it("should create a user and return it", () => {
    cy.request({
      method: 'POST',
      url: 'http://localhost:4201/test',
      body: newUser,
    }).as('new-user')
    cy.get('@new-user').its('status').should('eq', 200)
    cy.get('@new-user').its('body').then((res) => {
      let data = JSON.stringify(res);
      expect(data).to.eq(JSON.stringify(newUser));
    })
  })
})