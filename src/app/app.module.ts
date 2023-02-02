import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatInputModule } from '@angular/material/input';
import { MatIconModule } from '@angular/material/icon';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatGridListModule } from '@angular/material/grid-list'

import { AppComponent } from './app.component';
import { RouterModule, Routes } from '@angular/router';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { LoginPageComponent } from './login-page/login-page.component';
import { HomepageComponent } from './homepage/homepage.component';
import { TopBarComponent } from './top-bar/top-bar.component';

const appRoutes: Routes = [
    { path: 'login', component: LoginPageComponent },
    { path: '**', component: HomepageComponent },
    //{ path: '**', component: HomepageComponent}, // Will Definitely need to update this in the future.
  ];

@NgModule({
  declarations: [
    AppComponent,
    LoginPageComponent,
    HomepageComponent,
    TopBarComponent
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
    MatToolbarModule,
    MatGridListModule,
    RouterModule.forRoot(
        appRoutes,
        {enableTracing: true} // Debug only
    ),
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
