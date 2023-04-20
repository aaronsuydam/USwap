import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Item } from '../item';
import { ItemService } from '../services/item.service';

@Component({
  selector: 'app-swap-narrow-down',
  templateUrl: './swap-narrow-down.component.html',
  styleUrls: ['./swap-narrow-down.component.scss']
})
export class SwapNarrowDownComponent {

    constructor(private http: HttpClient, private itemService : ItemService) {}

    onInit() {
        this.search();
        for (let index = 0; index < this.items.length; index++) {
            this.items.push(this.itemService.getItem(index)); 
        }
    }

    swapFor: string = "";
    items : Item[] = [];
    numberOfItemsToDisplay = 6;
    filterNames : string[] = ["Filter 1", "Filter 2 - Longer", "Short", "Filter 4"];

    // FIXME: Talk to andrew about implementation
    getFilterNames() : void {
        this.http.get('swap-filters').subscribe(data => this.filterNames);
    }

    search() : void {
        if(this.swapFor == "") 
        {
            //Issue http request for homepage without a search
            //this.http.get()
        }
        //this.http.get('items-matching=swapFor', this.swapFor).subscribe(data => this.itemsToDisplay);
    }
}


