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
  loginSuccess: boolean = false;
  loginUrl: string = '${environment.serverUrl}/login';
  

  login(userName: string, userPassword: string): boolean
  {
    console.log(userName, userPassword);
    //this.getLoginSuccess(userName, userPassword).subscribe(data => this.loginSuccess = {loginSuccess : (data as any).loginSuccess}); 
    this.loginSuccess = true;
    if(this.loginSuccess)
    {
        console.log("AuthService: Credentials Accepted, Logging In...")
        return true;
    }
    else
    {
        console.log("AuthService: Credentials Refused, Please Enter Valid Credentials")
        return false;
    }
  }

  getLoginSuccess(userName: string, userPassword: string) {
    return this.httpClient.get<ServerRequestLogin>(this.loginUrl);
  }

}
