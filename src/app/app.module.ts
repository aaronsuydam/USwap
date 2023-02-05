import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatInputModule } from '@angular/material/input';
import { MatIconModule } from '@angular/material/icon';
import { MatFormFieldModule } from '@angular/material/form-field';

import { AppComponent } from './app.component';
import { RouterModule, Routes } from '@angular/router';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { LoginPageComponent } from './login-page/login-page.component';
import { HomepageComponent } from './homepage/homepage.component';
import { ItemDetailComponent } from './item-detail/item-detail.component';
import { UserProfileComponent } from './user-profile/user-profile.component';
import { SwapUiComponent } from './swap-ui/swap-ui.component';
import { SentOffersComponent } from './sent-offers/sent-offers.component';
import { SwapFinalComponent } from './swap-final/swap-final.component';

const appRoutes: Routes = [
    { path: 'home', component: HomepageComponent },
    { path: 'login', component: LoginPageComponent },
    { path: 'user-profile', component:UserProfileComponent},
    { path: 'swap-ui', component:SwapUiComponent},
    { path: 'sent-offers', component:SentOffersComponent},
    { path: 'swapNarrow', component:SwapFinalComponent},
    { path: '**', component: HomepageComponent}, // Will Definitely need to update this in the future.
  ];

@NgModule({
  declarations: [
    AppComponent,
    LoginPageComponent,
    HomepageComponent,
    ItemDetailComponent,
    UserProfileComponent,
    SwapUiComponent,
    SentOffersComponent,
    SwapFinalComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    BrowserAnimationsModule,
    FormsModule,
    MatButtonModule,
    MatInputModule,
    MatIconModule,
    MatFormFieldModule,
    RouterModule.forRoot(
        appRoutes,
        {enableTracing: true} // Debug only
    ),
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
