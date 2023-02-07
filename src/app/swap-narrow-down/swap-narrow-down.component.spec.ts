import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SwapNarrowDownComponent } from './swap-narrow-down.component';

describe('SwapNarrowDownComponent', () => {
  let component: SwapNarrowDownComponent;
  let fixture: ComponentFixture<SwapNarrowDownComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SwapNarrowDownComponent ]
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
