import { ComponentFixture, TestBed } from '@angular/core/testing';
import { NO_ERRORS_SCHEMA } from '@angular/core';

import { SwapUiComponent } from './swap-ui.component';

describe('SwapUiComponent', () => {
  let component: SwapUiComponent;
  let fixture: ComponentFixture<SwapUiComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SwapUiComponent ],
      schemas: [NO_ERRORS_SCHEMA]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SwapUiComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
