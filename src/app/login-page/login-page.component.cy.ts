import { HttpClientTestingModule } from "@angular/common/http/testing";
import { TestBed } from "@angular/core/testing";
import { AuthService } from "../services/auth.service";
import { LoginPageComponent } from "./login-page.component";
import { StorageService } from "../services/storage.service";
import { TestScheduler } from "rxjs/testing";
import { mount } from 'cypress/angular';
import { JsonPipe } from "@angular/common";

describe('LoginPageComponent', () => {

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
      providers: [
        AuthService,
        HttpClientTestingModule,
        StorageService
      ]
    });
  })

  let username = 'erob';
  let password = 'test';

  it('mounts', () => {
    mount(LoginPageComponent), {
      imports: [HttpClientTestingModule]
    }
    cy.get('input[name=username]').type(username)
    cy.get('input[name=password]').type(password)
    cy.get('[data-cy=login]').click()
  })

  it("should contain auth-user", () => {
    let storageService: StorageService = TestBed.inject(StorageService);
    storageService.saveUser({username, password})
    cy.getAllSessionStorage().then((res) => {
        expect(res['http://localhost:8080']['auth-user']).to.equal(
          JSON.stringify({
            "username": username, "password": password
          })
        )
    })
  })
})