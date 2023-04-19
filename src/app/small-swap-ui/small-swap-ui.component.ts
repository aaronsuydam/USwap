import { Component, Input } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-small-swap-ui',
  templateUrl: './small-swap-ui.component.html',
  styleUrls: ['./small-swap-ui.component.scss']
})
export class SmallSwapUiComponent {
    swappingFor: boolean = true;

    @Input() itemName : string = "Placeholder";
    @Input() itemOwner : string = "Placeholder";
    @Input() itemDescription : string = "Placeholder";
    @Input() imageSource : string = "../../assets/shiba2.jpg";

    like : string = "Like";

    constructor(private router: Router, private route: ActivatedRoute) {
        if(this.router.url !== "/swap-narrow")
        {
            this.swappingFor = false;
        }
    }

    onClickSwapFor(): void{
        this.router.navigate(['../user-profile'], {relativeTo: this.route});
    }

    onClickLike(): void{
        if(this.like = "Like")
        {
            this.like = "Liked!";
        }
    }
}
