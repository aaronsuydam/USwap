import { Component, Input } from '@angular/core';
import { Item } from '../item';
import { Router } from '@angular/router';

@Component({
  selector: 'app-user-profile',
  templateUrl: './user-profile.component.html',
  styleUrls: ['./user-profile.component.scss']
})
export class UserProfileComponent {
    

    @Input() username : string = "Placeholder";
    interests : string = "";
    profilePicPath : string = "../../assets/aaron-profile-pic.jpg";
    memberSinceDate : string = "Placeholder PHYear";
    userItems : Item[] = [];



}

    constructor(private router: Router) {}

    onClick() {
      this.router.navigate(['/add']);
    }
}
