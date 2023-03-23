import { Component, isStandalone, OnInit } from '@angular/core';
import { AuthService } from '../services/auth.service';
import { ThemeServiceService } from '../theme-service.service';

@Component({
  selector: 'app-top-bar',
  templateUrl: './top-bar.component.html',
  styleUrls: ['./top-bar.component.css']
})
export class TopBarComponent implements OnInit {
  constructor(private authService: AuthService, private themeService: ThemeServiceService) {};
  isAuthenticated: boolean = false;
  
  ngOnInit(): void {
    this.isAuthenticated = this.authService.isAuthenticated();
  }

  changeTheme(): void {
    this.themeService.toggleDarkTheme();
  }
  
}
