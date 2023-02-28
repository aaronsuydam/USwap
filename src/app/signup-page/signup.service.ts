import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { User } from '../interfaces/UserInterface';

@Injectable({
  providedIn: 'root'
})
export class SignupService {
  baseUrl: string = 'test';
  readonly headers = new HttpHeaders().set('Content-Type', 'application/json');

  constructor(private http: HttpClient) {}

  addUser(user: User): Observable<User> {
    return this.http.post<User>(this.baseUrl, user, {headers: this.headers});
  }
}
