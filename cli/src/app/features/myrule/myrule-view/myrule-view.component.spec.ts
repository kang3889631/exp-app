import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { MyruleViewComponent } from './myrule-view.component';

describe('MyruleViewComponent', () => {
  let component: MyruleViewComponent;
  let fixture: ComponentFixture<MyruleViewComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ MyruleViewComponent ],
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MyruleViewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
