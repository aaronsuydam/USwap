
import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import { MaterialModule } from '../assets/material.module';
import { RouterTestingModule } from '@angular/router/testing';
import { HttpClientTestingModule } from '@angular/common/http/testing';

import { AppComponent } from './app.component';
import { RouterModule, Routes } from '@angular/router';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { LoginPageComponent } from './login-page/login-page.component';
import { HomepageComponent } from './homepage/homepage.component';
import { TopBarComponent } from './top-bar/top-bar.component';
import { ItemDetailComponent } from './item-detail/item-detail.component';
import { UserProfileComponent } from './user-profile/user-profile.component';
import { SwapUiComponent } from './swap-ui/swap-ui.component';
import { SwapFinalComponent } from './swap-final/swap-final.component';
import { SwapNarrowDownComponent } from './swap-narrow-down/swap-narrow-down.component';
import { SmallSwapUiComponent } from './small-swap-ui/small-swap-ui.component';
import { APIInterceptor } from './services/interceptor.service';
import { authGuard } from './services/auth.service';
import { SignupPageComponent } from './signup-page/signup-page.component';

const appRoutes: Routes = [
    { path: '', title: "USwap Home", component: HomepageComponent},
    { path: 'login', title: "Login - USwap", component: LoginPageComponent },
    { path: 'signup', title: "Signup - USwap", component: SignupPageComponent },
    { path: 'user-profile', title: "Profile and Items - USwap", canActivate: [authGuard], component:UserProfileComponent},
    { path: 'swap-narrow', title: "Swap For Anything! - USwap", canActivate: [authGuard], component:SwapNarrowDownComponent},
    { path: 'swap-final', title: "Confirm Swap - USwap", canActivate: [authGuard], component:SwapFinalComponent},
    { path: '**', redirectTo: '', pathMatch: 'full'} // Can direct to an about page or error page
  ];

@NgModule({
  declarations: [
    AppComponent,
    LoginPageComponent,
    HomepageComponent,
    TopBarComponent,
    ItemDetailComponent,
    UserProfileComponent,
    SwapUiComponent,
    SwapFinalComponent,
    SwapNarrowDownComponent,
    SmallSwapUiComponent,
    SignupPageComponent,
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    BrowserAnimationsModule,
    MaterialModule,
    FormsModule,
    RouterModule.forRoot(appRoutes),
  ],
  providers: [
    {
      provide: HTTP_INTERCEPTORS,
      useClass: APIInterceptor,
      multi: true
    },
    HttpClientModule,
    APIInterceptor,
    RouterTestingModule,
    HttpClientTestingModule
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }