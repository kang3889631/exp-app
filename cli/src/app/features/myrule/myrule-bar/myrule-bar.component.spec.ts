import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { MyruleBarComponent } from './myrule-bar.component';

describe('MyruleBarComponent', () => {
  let component: MyruleBarComponent;
  let fixture: ComponentFixture<MyruleBarComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ MyruleBarComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MyruleBarComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
