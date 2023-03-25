import { BackendService } from 'src/app/services/backend.service';
import { OrderDto } from './../types/OrderDto';

import { Wallets } from './../types/Wallets';
import { Component, OnInit } from '@angular/core';
import { Currrency } from '../types/Currency';
import { Wallet } from '../types/Wallet';
import { OrderForm } from '../types/OrderForm';
import { Order } from '../types/Order';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit {
  constructor(private backend:BackendService){}

  ngOnInit(): void {




    var userWallet = <Wallets> this.helper()
    this.providers.GHS = userWallet.GHS
    this.providers.USD = userWallet.USD
    this.providers.NGN = userWallet.NGN
    this.providers.KES = userWallet.KES
    this.usersCurrency = [this.providers.USD,this.providers.KES,this.providers.GHS,this.providers.NGN]
    
    console.log(this.providers)
    // console.log(wallet)
  }

  helper(){
    const {wallet} = JSON.parse(localStorage.getItem("data")!)
    return wallet
  }








  providers:Wallet = new Wallet(new Currrency(true,0,"USD"),new Currrency(true,0,"KES"),new Currrency(true,0,"GHS"),new Currrency(false,0,"NGN"))

  usersCurrency!:Currrency[]
  orderList!:OrderForm[]




}
