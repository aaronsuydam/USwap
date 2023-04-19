import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-swap-narrow-down',
  templateUrl: './swap-narrow-down.component.html',
  styleUrls: ['./swap-narrow-down.component.scss']
})
export class SwapNarrowDownComponent {

    constructor(private http: HttpClient) {}

    swapFor: string = "";

    itemsToDisplay : string[] = ["string1", "string2", "string3", "string4", "string5", "string6"];

    filterNames : string[] = ["Filter 1", "Filter 2 - Longer", "Short", "Filter 4"];

    // FIXME: Talk to andrew about implementation
    getFilterNames() : void {
        this.http.get('swap-filters').subscribe(data => this.filterNames);
    }

    search() : void {
        //this.http.get('items-matching=swapFor', this.swapFor).subscribe(data => this.itemsToDisplay);
    }
}


