import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { MyruleBasicFormComponent } from './myrule-basic-form.component';

describe('MyruleBasicFormComponent', () => {
  let component: MyruleBasicFormComponent;
  let fixture: ComponentFixture<MyruleBasicFormComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ MyruleBasicFormComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MyruleBasicFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
