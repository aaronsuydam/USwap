import { Component, OnInit } from '@angular/core';
import { StorageService } from '../services/storage.service';

@Component({
  selector: 'app-top-bar',
  templateUrl: './top-bar.component.html',
  styleUrls: ['./top-bar.component.scss']
})
export class TopBarComponent implements OnInit {
  constructor(private storageService: StorageService) {};
  isLoggedIn: boolean = false;
  
  ngOnInit(): void {
    this.isLoggedIn = this.storageService.isLoggedIn();
  }

}
