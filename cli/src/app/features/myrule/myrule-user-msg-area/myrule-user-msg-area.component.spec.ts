import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { MyruleUserMsgAreaComponent } from './myrule-user-msg-area.component';

describe('MyruleUserMsgAreaComponent', () => {
  let component: MyruleUserMsgAreaComponent;
  let fixture: ComponentFixture<MyruleUserMsgAreaComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ MyruleUserMsgAreaComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MyruleUserMsgAreaComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
