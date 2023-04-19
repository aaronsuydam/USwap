import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { StorageService } from '../services/storage.service';

@Component({
  selector: 'app-swap-narrow-down',
  templateUrl: './swap-narrow-down.component.html',
  styleUrls: ['./swap-narrow-down.component.scss']
})
export class SwapNarrowDownComponent {

    constructor(private http: HttpClient) {}

    swapFor: string = "";

    itemsToDisplay : string[] = [];

    filterNames : string[] = ["Filter 1", "Filter 2 - Longer", "Short", "Filter 4"];

    
}


