import { ComponentFixture, TestBed } from '@angular/core/testing';
import { NO_ERRORS_SCHEMA } from '@angular/core';

import { SwapNarrowDownComponent } from './swap-narrow-down.component';
import { SwapUiComponent } from '../swap-ui/swap-ui.component';

describe('SwapNarrowDownComponent', () => {
  let component: SwapNarrowDownComponent;
  let fixture: ComponentFixture<SwapNarrowDownComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ 
        SwapNarrowDownComponent,
        SwapUiComponent
      ],
      schemas: [NO_ERRORS_SCHEMA]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SwapNarrowDownComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
