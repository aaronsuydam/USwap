import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class MessageService {

  constructor(private http: HttpClient) { }

    // FIXME: Talk to andrew about implementation
    getFilterNames() : void {
        this.http.get('swap-filters').subscribe(data => this.filterNames);
    }

    search() : void {
        this.http.get('items-matching=swapFor', this.swapFor).subscribe(data => this.itemsToDisplay);
    }

}
