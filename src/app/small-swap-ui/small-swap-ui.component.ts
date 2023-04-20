import { Component, Input } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
import { Item } from '../item';
import { ItemService } from '../services/item.service';

@Component({
  selector: 'app-small-swap-ui',
  templateUrl: './small-swap-ui.component.html',
  styleUrls: ['./small-swap-ui.component.scss']
})
export class SmallSwapUiComponent {
    swappingFor: boolean = true;
    item : Item = new Item(NaN, "", "", NaN, "");
    like : string = "Like";

    constructor(private router: Router, private route: ActivatedRoute, private itemService : ItemService) {
        if(this.router.url !== "/swap-narrow")
        {
            this.swappingFor = false;
        }
    }

    ngOnInit() {
        if(Number.isNaN(this.item.ID))
        {
            this.itemService.setDefaultValues(this.item);
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
