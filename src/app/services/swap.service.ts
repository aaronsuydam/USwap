import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class SwapService {

    swapFor: string = "";
    swapWith: string = "";
    constructor() { }

    swap(swapFor?: string, swapWith?: string): void{

    }
}
