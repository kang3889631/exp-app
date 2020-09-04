import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { MyruleDemoViewComponent } from './myrule-demo-view.component';

describe('MyruleDemoViewComponent', () => {
  let component: MyruleDemoViewComponent;
  let fixture: ComponentFixture<MyruleDemoViewComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ MyruleDemoViewComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MyruleDemoViewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
