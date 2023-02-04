import { Component } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-homepage',
  templateUrl: './homepage.component.html',
  styleUrls: ['./homepage.component.css']
})
export class HomepageComponent {
    swapFor: string = "something";
    swapWith: string = "better thing";

    userSwapFor: string = "";
    userSwapWith: string = "";

    constructor(private router: Router, private route: ActivatedRoute) {}

    swap(): void {
        console.log("attempt");
        this.router.navigate(['../swapNarrow'], {relativeTo: this.route});
    }
}
