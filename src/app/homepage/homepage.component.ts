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
