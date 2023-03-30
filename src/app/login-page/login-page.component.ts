import { Component } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { AuthService } from '../services/auth.service';
import { StorageService } from '../services/storage.service';

@Component({
  selector: 'app-login-page',
  templateUrl: './login-page.component.html',
  styleUrls: ['./login-page.component.scss']
})
export class LoginPageComponent {
    form: FormGroup;
    isLoggedIn = false;

    constructor(private authService: AuthService,
                private storageService: StorageService,
                private router: Router,
                private fb: FormBuilder ) {
        
        this.form = this.fb.group({
            username: ['', Validators.required],
            password: ['', Validators.required]
        });
    }

    login() {
        const val = this.form.value;

        if (val.username && val.password) {
            this.authService.login(val.username, val.password)
                .subscribe(
                    data => {
                        this.storageService.saveUser(data)
                        this.isLoggedIn = true;
                        this.reloadPage();
                        this.router.navigate(['/swap-narrow']);
                    }
                )
        }
    }

    signUp() {
        this.router.navigate(['/signup']);
    }

    reloadPage(): void {
        setTimeout(()=>{
            window.location.reload();
        }, 100);
    }
}