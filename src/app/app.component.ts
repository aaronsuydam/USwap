import {Component, OnInit} from '@angular/core';
import { Observable } from 'rxjs';
import { ThemeServiceService } from './theme-service.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {

    isDarkTheme: Observable<boolean> = new Observable<boolean>;

    constructor(
      private themeService: ThemeServiceService,
      ) {}
  
    ngOnInit() {
      this.isDarkTheme = this.themeService.isDarkTheme;
    }
}
