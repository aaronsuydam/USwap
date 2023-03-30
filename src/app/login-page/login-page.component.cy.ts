import { HttpClientTestingModule } from "@angular/common/http/testing";
import { TestBed } from "@angular/core/testing";
import { AuthService } from "../services/auth.service";
import { LoginPageComponent } from "./login-page.component";

describe('LoginPageComponent', () => {

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
      providers: [AuthService, HttpClientTestingModule]
    });
  })

  it('mounts', () => {
    cy.mount(LoginPageComponent), {
      imports: [HttpClientTestingModule]
    }
  })
})