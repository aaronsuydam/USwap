import { Component } from '@angular/core';
import { AuthService } from '../auth.service';

@Component({
  selector: 'app-login-page',
  templateUrl: './login-page.component.html',
  styleUrls: ['./login-page.component.css']
})
export class LoginPageComponent {
    loginField = document.querySelector("input");

    constructor(private authService: AuthService) {}

    /**
     * Ok so we are going to have some kind of loginService. this login service will need to
     * take in user credentials, interface with the backend, determine if the login attempt 
     * is valid, and return the result. Something like the following:
     * 
     * login_attempt(): boolean {
     *      return true if loginService.login()===true;
     * }
     */
    loginAttempt(userName: string, userPassword: string): boolean {
        console.log("attempt");
        console.log(userName, userPassword);
        if(this.authService.login(userName, userPassword))
        {
            return true;
        }
        else
        {
            return false;
        }
    }

    onClick(userName: string = "", userPassword: string = "") {
        this.loginAttempt(userName, userPassword);
    }

}
