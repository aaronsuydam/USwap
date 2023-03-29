import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, startWith, throwError } from 'rxjs';
import { ServerRequestLogin } from './server-request-login';
import { Router } from '@angular/router';
import { inject } from '@angular/core';

// This function should access the local cached information of whether or not the user is logged in,
// and return true or false accordingly
export const authGuard = () => {
    const authService = inject(AuthService);
    const router = inject(Router);
    console.log('authGuard#canActivate called');
    if(authService.isLoggedIn) {
        return true;
    } else {
        return router.parseUrl('/login');
    }
};

@Injectable({
  providedIn: 'root'
})

export class AuthService {

  constructor(private httpClient: HttpClient) { }
  loginAuthorized: ServerRequestLogin = {loginSuccess : false};
  loginUrl: string = `login`;
  isLoggedIn: boolean = false;
  

  login(userName: string, userPassword: string): boolean {
    console.log(userName, userPassword);
    this.getLoginSuccess(userName, userPassword).subscribe(data => this.loginAuthorized = {loginSuccess : (data as any).loginSuccess}); 

    this.getLoginSuccess(userName, userPassword).subscribe(data => console.log(data));

    this.loginAuthorized.loginSuccess = true;
    if(this.loginAuthorized) {
        console.log("AuthService: Credentials Accepted, Logging In...");
        localStorage.setItem('isLoggedIn', 'true');
        this.isLoggedIn = true;
        return true;
    }
    else {
        console.log("AuthService: Credentials Refused, Login Failed");
        console.log("AuthService: Please Enter Valid Credentials");
        return false;
    }
  }

  getLoginSuccess(userName: string, userPassword: string) {
    console.log(this.httpClient.get(this.loginUrl));
    return this.httpClient.get(this.loginUrl);
  }

  isAuthenticated(): boolean {
    if (localStorage.getItem('isLoggedIn')) {
      return true;
    } else {
      return false;
    }
  }

}
