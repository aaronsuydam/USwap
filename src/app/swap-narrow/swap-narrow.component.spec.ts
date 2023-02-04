import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SwapNarrowComponent } from './swap-narrow.component';

describe('SwapNarrowComponent', () => {
  let component: SwapNarrowComponent;
  let fixture: ComponentFixture<SwapNarrowComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SwapNarrowComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SwapNarrowComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
