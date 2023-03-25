import { BackendService } from 'src/app/services/backend.service';
import { OrderForm } from 'src/app/types/OrderForm';
import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-orders-card',
  templateUrl: './orders-card.component.html',
  styleUrls: ['./orders-card.component.css']
})
export class OrdersCardComponent implements OnInit  {
  constructor(private backend:BackendService){

  }
  ngOnInit(): void {
    this.orderList =this.backend.orderList
  }
  orderList!:OrderForm[]


 

}
