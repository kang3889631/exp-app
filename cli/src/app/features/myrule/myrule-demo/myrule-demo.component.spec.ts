import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { MyruleDemoComponent } from './myrule-demo.component';

describe('MyruleDemoComponent', () => {
  let component: MyruleDemoComponent;
  let fixture: ComponentFixture<MyruleDemoComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ MyruleDemoComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MyruleDemoComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
