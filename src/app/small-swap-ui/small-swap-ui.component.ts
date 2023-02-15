import { Component } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-small-swap-ui',
  templateUrl: './small-swap-ui.component.html',
  styleUrls: ['./small-swap-ui.component.css']
})
export class SmallSwapUiComponent {

    constructor(private router: Router, private route: ActivatedRoute) {}


    onClick(): void{
        this.router.navigate(['../user-profile'], {relativeTo: this.route});
    }
}
