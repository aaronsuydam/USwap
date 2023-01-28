import { Component } from '@angular/core';
import { AuthService } from '../auth.service';

@Component({
  selector: 'app-login-page',
  templateUrl: './login-page.component.html',
  styleUrls: ['./login-page.component.css']
})
export class LoginPageComponent {
    loginField = document.querySelector("input");
    userName: string = "";
    userPassword: string = "";

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
    loginAttempt(): boolean {
        console.log("attempt");
        console.log(this.userName, this.userPassword);
        if(this.authService.login(this.userName, this.userPassword))
        {
            return true;
        }
        else
        {
            return false;
        }
    }

    onClick() {
        this.loginAttempt();
    }

}
