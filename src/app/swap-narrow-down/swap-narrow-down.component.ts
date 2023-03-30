import { Component } from '@angular/core';
import { StorageService } from '../services/storage.service';

@Component({
  selector: 'app-swap-narrow-down',
  templateUrl: './swap-narrow-down.component.html',
  styleUrls: ['./swap-narrow-down.component.scss']
})
export class SwapNarrowDownComponent {
    swapFor: string = "";
}
