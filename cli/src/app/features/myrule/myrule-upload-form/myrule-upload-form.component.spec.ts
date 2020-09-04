import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { MyruleUploadFormComponent } from './myrule-upload-form.component';

describe('MyruleUploadFormComponent', () => {
  let component: MyruleUploadFormComponent;
  let fixture: ComponentFixture<MyruleUploadFormComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ MyruleUploadFormComponent ],
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MyruleUploadFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
