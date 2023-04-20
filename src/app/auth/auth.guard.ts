import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, CanActivate, RouterStateSnapshot, Router } from '@angular/router';
import { Observable } from 'rxjs';
import { StorageService } from '../services/storage.service';

@Injectable({
  providedIn: 'root'
})
export class AuthGuard implements CanActivate {
  constructor(private storageService: StorageService, private router: Router) {};

  canActivate(
    next: ActivatedRouteSnapshot,
    state: RouterStateSnapshot) {
      console.log('CanActivate called');
      // let isLoggedIn = this.storageService.isLoggedIn();
      let isLoggedIn = true;
      if (isLoggedIn) {
        return true;
      } else {
        this.router.navigate(['../login']);
      }
      return false;
    }
}
