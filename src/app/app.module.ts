import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import { MaterialModule } from './material/material.module';

import { AppComponent } from './app.component';
import { RouterModule, Routes } from '@angular/router';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { LoginPageComponent } from './login-page/login-page.component';
import { HomepageComponent } from './homepage/homepage.component';
import { TopBarComponent } from './top-bar/top-bar.component';
import { ItemDetailComponent } from './item-detail/item-detail.component';
import { UserProfileComponent } from './user-profile/user-profile.component';
import { SwapUiComponent } from './swap-ui/swap-ui.component';
import { SentOffersComponent } from './sent-offers/sent-offers.component';
import { SwapFinalComponent } from './swap-final/swap-final.component';
import { SwapNarrowDownComponent } from './swap-narrow-down/swap-narrow-down.component';
import { SmallSwapUiComponent } from './small-swap-ui/small-swap-ui.component';

const appRoutes: Routes = [
    { path: '**', component: HomepageComponent},
    { path: 'login', component: LoginPageComponent },
    { path: 'user-profile', component:UserProfileComponent},
    { path: 'swap-ui', component:SwapUiComponent},
    { path: 'sent-offers', component:SentOffersComponent},
    { path: 'swap-narrow', component:SwapNarrowDownComponent},
    { path: 'swap-final', component:SwapFinalComponent},
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
    SentOffersComponent,
    SwapFinalComponent,
    SwapNarrowDownComponent,
    SmallSwapUiComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    BrowserAnimationsModule,
    MaterialModule,
    FormsModule,
    RouterModule.forRoot(
        appRoutes,
        {enableTracing: true} // Debug only
    ),
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
