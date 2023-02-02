import { Component } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
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
    loginSuccess: boolean = false;
    value: string = 'Clear me';

    constructor(private authService: AuthService, private router: Router, private route: ActivatedRoute ) {}

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
        if(this.authService.login(this.userName, this.userPassword) === true) {
            this.router.navigate(['home'], {relativeTo: this.route});
            return true;
        } else {
            return false;
        }
    }

    onClick() {
        this.loginSuccess = this.loginAttempt();
        this.userName = "";
        this.userPassword = "";
        if(this.loginSuccess) {
            console.log("Logging you in");
            // Navigate to the new page.
        }
    }

}
