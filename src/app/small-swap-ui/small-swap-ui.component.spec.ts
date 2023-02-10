import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SmallSwapUiComponent } from './small-swap-ui.component';

describe('SmallSwapUiComponent', () => {
  let component: SmallSwapUiComponent;
  let fixture: ComponentFixture<SmallSwapUiComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SmallSwapUiComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SmallSwapUiComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
