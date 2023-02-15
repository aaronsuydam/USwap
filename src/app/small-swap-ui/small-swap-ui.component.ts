import { Component } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-small-swap-ui',
  templateUrl: './small-swap-ui.component.html',
  styleUrls: ['./small-swap-ui.component.css']
})
export class SmallSwapUiComponent {


    swappingFor: boolean = true;
    constructor(private router: Router, private route: ActivatedRoute) {
        if(this.router.url !== "/swap-narrow")
        {
            this.swappingFor = false;
        }
    }


    onClickSwapFor(): void{
        this.router.navigate(['../user-profile'], {relativeTo: this.route});
    }

    onClickSwapWith(): void{
        this.router.navigate(['../swap-final'], {relativeTo: this.route});
    }
}
