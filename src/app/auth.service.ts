import { Injectable } from '@angular/core';
import { HttpClient, HttpInterceptor} from '@angular/common/http';
import { environment } from '../environments/environment.development';
import { Observable, startWith, throwError } from 'rxjs';
import { ServerRequestLogin } from './server-request-login';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  constructor(private httpClient: HttpClient) { }
  loginAuthorized: ServerRequestLogin = {loginSuccess : false};
  loginUrl: string = `/login`;
  

  login(userName: string, userPassword: string): boolean
  {
    console.log(userName, userPassword);
    this.getLoginSuccess(userName, userPassword).subscribe(data => this.loginAuthorized = {loginSuccess : (data as any).loginSuccess}); 
    this.loginAuthorized.loginSuccess = true;
    if(this.loginAuthorized)
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
    return this.httpClient.get(this.loginUrl);
  }
}
