import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { catchError, retry } from 'rxjs/operators';
import { ServerRequestLogin } from './server-request-login';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  constructor(private httpClient: HttpClient) { }
  loginSuccess: ServerRequestLogin | undefined;
  loginUrl: string = "wahtever the heck the login url is.";
  

  login(userName: string, userPassword: string): boolean
  {
    console.log(userName, userPassword);
    this.getLoginSuccess(userName, userPassword).subscribe(data => this.loginSuccess =
         {loginSuccess : (data as any).loginSuccess}); 
    if(this.loginSuccess?.loginSuccess)
    {
        return true;
    }
    else
    {
        return false;
    }
  }

  getLoginSuccess(userName: string, userPassword: string) {
    return this.httpClient.get<ServerRequestLogin>(this.loginUrl);
  }


  interpret() {

  }
}
