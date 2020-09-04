import { TestBed } from '@angular/core/testing';

import { MyruleService } from './myrule.service';

describe('MyruleService', () => {
  let service: MyruleService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(MyruleService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
