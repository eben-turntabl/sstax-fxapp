import { OrdersCardComponent } from './dashboard/orders-card/orders-card.component';
import { WithdrawCardComponent } from './dashboard/withdraw-card/withdraw-card.component';
import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { LoginComponent } from './components/login/login.component';
import { RegisterComponent } from './components/register/register.component';
import { HomeComponent } from './components/home/home.component';

import { WalletComponent } from './dashboard/wallet/wallet.component';
import { CardModule, GridModule } from '@coreui/angular';
import { DashboardComponent } from './dashboard/dashboard.component';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { RouterModule } from '@angular/router';
import { CommonModule } from '@angular/common';

import { FaIconLibrary } from '@fortawesome/angular-fontawesome';
import { faStar as farStar } from '@fortawesome/free-regular-svg-icons';
import { faStar as fasStar } from '@fortawesome/free-solid-svg-icons';
import { OperationCardComponent } from './dashboard/operation-card/operation-card.component';
import { HttpClientModule } from '@angular/common/http';






@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    RegisterComponent,
    HomeComponent,
    WalletComponent,
    DashboardComponent,
    OperationCardComponent,
    WithdrawCardComponent,
    OrdersCardComponent,


  ],
  imports: [
    RouterModule,
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    CardModule,
    GridModule,
    CommonModule,
    FontAwesomeModule,
    ReactiveFormsModule,
    HttpClientModule,



  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule {
  constructor(library: FaIconLibrary) {
    library.addIcons(fasStar, farStar);
  }



}


