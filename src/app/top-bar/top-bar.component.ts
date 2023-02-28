import { Component, isStandalone, OnInit } from '@angular/core';
import { AuthService } from '../services/auth.service';

@Component({
  selector: 'app-top-bar',
  templateUrl: './top-bar.component.html',
  styleUrls: ['./top-bar.component.css']
})
export class TopBarComponent implements OnInit {
  constructor(private authService: AuthService) {};
  isAuthenticated: boolean = false;
  
  ngOnInit(): void {
    this.isAuthenticated = this.authService.isAuthenticated();
  }
  
}
