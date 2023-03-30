import { Component } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
import { SignupService } from './signup.service';
import { User } from '../interfaces/UserInterface';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-signup-page',
  templateUrl: './signup-page.component.html',
  styleUrls: ['./signup-page.component.scss']
})
export class SignupPageComponent {
  hide: boolean = true;
  confirmHide: boolean = true;
  validatePassword: string = "";

  user: User = {
    username: "",
    email: "",
    password: ""
  };

  constructor(private router: Router, private route: ActivatedRoute, private signupService: SignupService ) {}

  // check if username is taken in database
  checkUsername() {
    return true;
  }

  // check if email is taken in database
  checkEmail() {
    // checking for a valid UF email
    const re = /^\w+([\.-]?\w+)*@ufl.edu/gm;
    if (re.test(this.user.email))
      return true;
    return false;
  }

  checkPassword(): boolean {
    if ((this.user.password && this.validatePassword) && this.user.password !== this.validatePassword)
      return false;
    return true;
  }

  // create new user in database
  async addUser() {
    this.signupService.addUser(this.user).subscribe(
      data => {
        console.log(data);
      }
    );
  }

  onClick() {
    this.addUser();
    //this.router.navigate(['../login'], {relativeTo: this.route});
    if (this.checkUsername()) {
      if (this.checkEmail()) {
        if (this.checkPassword()) {
          console.log("Registered");
          this.addUser();
          this.router.navigate(['../login'], {relativeTo: this.route});
        }
      }
    }
  }
}
