import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SwapFinalComponent } from './swap-final.component';

describe('SwapNarrowComponent', () => {
  let component: SwapFinalComponent;
  let fixture: ComponentFixture<SwapFinalComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SwapFinalComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SwapFinalComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
