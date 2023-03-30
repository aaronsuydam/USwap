import { StorageService } from "src/app/services/storage.service";

describe('AuthService', () => {
    let username = 'erob';
    let password = 'test';

    it("should login", () => {
        cy.request({
            method: 'POST',
            url: 'http://localhost:4201/login',
            body: {username, password},
        }).as('login')
        cy.get('@login').its('status').should('eq', 200)
        cy.get('@login').getCookie('token').should('exist')
    })
})