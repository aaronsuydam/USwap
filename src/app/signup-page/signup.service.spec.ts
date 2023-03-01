import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { TestBed } from '@angular/core/testing';
import { HttpClient, HttpResponse } from '@angular/common/http';

import { User } from '../interfaces/UserInterface';
import { SignupService } from './signup.service';

describe('#SignupService.addUser()', () => {
  let httpClient: HttpClient;
  let httpTestingController: HttpTestingController;
  let signupService: SignupService;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
      providers: [SignupService]
    });

    httpClient = TestBed.inject(HttpClient);
    httpTestingController = TestBed.inject(HttpTestingController);
    signupService = TestBed.inject(SignupService);
  });

  afterEach(() => {
    httpTestingController.verify();
  });

  // Test case 1
  it('should add a user and return it', async () => {
    const newUser: User = {
      username: "erob",
      email: "evan.robinson@ufl.edu",
      password: "test"
    };

    signupService.addUser(newUser).subscribe({
      next: (data: any) => {
        expect(data)
        .withContext('should return the user')
        .toEqual(newUser);
      },
      error: (err: any) => { console.log(err); }
    });

    const req = httpTestingController.expectOne(signupService.baseUrl);
    expect(req.request.method).toEqual('POST');
    expect(req.request.body).toEqual(newUser);

    const expectedResponse = new HttpResponse({ status: 201, body: newUser });
    req.event(expectedResponse);
  });
});

