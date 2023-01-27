import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  constructor() { }

  login(userName: string, userPassword: string): boolean
  {
    console.log(userName, userPassword);
    return false;
  }
}
