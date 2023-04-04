import { Component } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';

// export interface Tile {
//   color: string;
//   cols: number;
//   rows: number;
//   text: string;
// }

@Component({
  selector: 'app-homepage',
  templateUrl: './homepage.component.html',
  styleUrls: ['./homepage.component.scss']
})
export class HomepageComponent {
  // tiles: Tile[] = [
  //   {text: 'One', cols: 3, rows: 1, color: 'lightblue'},
  //   {text: 'Two', cols: 1, rows: 2, color: 'lightgreen'},
  //   {text: 'Three', cols: 1, rows: 1, color: 'lightpink'},
  //   {text: 'Four', cols: 2, rows: 1, color: '#DDBDF1'},
  // ];
    swapFor: string = "something";
    swapWith: string = "better thing";

    userSwapFor: string = "";
    userSwapWith: string = "";

    constructor(private router: Router, private route: ActivatedRoute) {}

    login(): void {
        console.log("attempt");
        this.router.navigate(['../login'], {relativeTo: this.route});
    }
}
