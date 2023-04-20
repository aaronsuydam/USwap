import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Item } from '../item';
import { ItemService } from '../services/item.service';

@Component({
  selector: 'app-swap-narrow-down',
  templateUrl: './swap-narrow-down.component.html',
  styleUrls: ['./swap-narrow-down.component.scss']
})
export class SwapNarrowDownComponent implements OnInit {

    constructor(private http: HttpClient, private itemService : ItemService) {}

    onInit() {
        for (let index = 0; index < this.items.length; index++) {
            console.log("For Loop number" + index);
            this.items.push(this.itemService.getItem(index)); 
        }
    }
    
    swapFor: string = "";
    items : Item[] = [];
    numberOfItemsToDisplay = 6;
    filterNames : string[] = ["Filter 1", "Filter 2 - Longer", "Short", "Filter 4"];

    ngOnInit(): void {
        this.getItems();
    }

    getItems(): void {
        this.http.get<Item[]>('items').subscribe({
            next: (res) => {
                this.items = res;
            },
            error: (error) => {
                console.log(error)
            }
        });
    }

    // FIXME: Talk to andrew about implementation
    getFilterNames() : void {
        this.http.get('swap-filters').subscribe(data => this.filterNames);
    }
}


