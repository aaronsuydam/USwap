import { Component } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-signup-page',
  templateUrl: './signup-page.component.html',
  styleUrls: ['./signup-page.component.css']
})
export class SignupPageComponent {
  hide: boolean = true;
  confirmHide: boolean = true;

  userName: string = "";
  userPassword: string = "";
  validatePassword: string = "";
  userEmail: string = "";

  constructor(private router: Router, private route: ActivatedRoute ) {}
  
  // check if username is taken in database
  checkUsername() {
    return true;
  }

  // check if email is taken in database
  checkEmail() {
    // checking for a valid UF email
    const re = /^\w+([\.-]?\w+)*@ufl.edu/gm;
    if (re.test(this.userEmail))
      return true;
    return false;
  }

  checkPassword(): boolean {
    if ((this.userPassword && this.validatePassword) && this.userPassword !== this.validatePassword)
      return false;
    return true;
  }

  // create new user in database
  registerUser() {

  }

  onClick() {
    if (this.checkUsername()) {
      if (this.checkEmail()) {
        if (this.checkPassword()) {
          console.log("Registered");
          this.registerUser();
          this.router.navigate(['../login'], {relativeTo: this.route});
        }
      }
    }
  }
}
