import {Component, OnInit} from '@angular/core';
import {HelloWorldService} from './hello-world.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  title: any;
  field: string = 'title';

  constructor(private hw: HelloWorldService) {}

  ngOnInit(): void {
    this.hw.get().subscribe((data) => {
      console.log(data);
      this.title = data[this.field as keyof Object];
    })
  }
}

