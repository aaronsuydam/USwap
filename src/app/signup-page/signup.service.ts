import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse, HttpHeaders } from '@angular/common/http';
import { catchError, Observable, tap, of } from 'rxjs';
import { User } from '../interfaces/UserInterface';

@Injectable({
  providedIn: 'root'
})
export class SignupService {
  baseUrl: string = 'signup';
  constructor(private http: HttpClient) {}
  
  readonly headers = new HttpHeaders().set('Content-Type', 'application/json');

  addUser(user: User): Observable<User> {
    return this.http.post<User>(this.baseUrl, user, {headers: this.headers})
    .pipe(
      catchError(this.handleError(user))
    );
  };
  private handleError<T>(result = {} as T) {
    return (error: HttpErrorResponse) : Observable<T> => {
      console.log(error);
      return of(result);
    };
  };
}
