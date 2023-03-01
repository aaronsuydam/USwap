import { TestBed } from '@angular/core/testing';

import { APIInterceptor } from './interceptor.service';

describe('APIInterceptorService', () => {
  let service: APIInterceptor;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(APIInterceptor);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
