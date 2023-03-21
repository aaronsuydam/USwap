import { TestBed } from '@angular/core/testing';

import { APIInterceptor } from './interceptor.service';

describe('InterceptorService', () => {
  let service: APIInterceptor;

  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [
        APIInterceptor,
      ],
    });

    service = TestBed.inject(APIInterceptor);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
