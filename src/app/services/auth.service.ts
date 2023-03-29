import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { User } from '../interfaces/UserInterface';
import { Observable, startWith, throwError, shareReplay, tap } from 'rxjs';
import { ServerRequestLogin } from './server-request-login';
import { Router } from '@angular/router';
import { inject } from '@angular/core';

// This function should access the local cached information of whether or not the user is logged in,
// and return true or false accordingly
// export const authGuard = () => {
//     const authService = inject(AuthService);
//     const router = inject(Router);
//     console.log('authGuard#canActivate called');
//     if(authService.isLoggedIn) {
//         return true;
//     } else {
//         return router.parseUrl('/login');
//     }
// };

@Injectable({
  providedIn: 'root'
})

export class AuthService {

  constructor(private http: HttpClient) { }
  
  login(username: string, password: string): Observable<any> {
    return this.http.post<any>('login', {username, password}).pipe(
      shareReplay()
    );
  }

}
