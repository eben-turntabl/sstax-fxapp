import { ExchangeForm } from './../../types/ExchangeForm';
import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { BackendService } from 'src/app/services/backend.service';
import { OrderForm } from 'src/app/types/OrderForm';

@Component({
  selector: 'app-operation-card',
  templateUrl: './operation-card.component.html',
  styleUrls: ['./operation-card.component.css']
})
export class OperationCardComponent implements OnInit {
  orderForm!:FormGroup;
  constructor(private backend:BackendService){

}
  ngOnInit(): void {

    this.orderForm = new FormGroup({

      "currency": new FormControl("USD", Validators.required),
      "amount": new FormControl("0.00", Validators.required),
      "bank_account": new FormControl("bank_account", Validators.required),
      "client_id": new FormControl("client_id", Validators.required),
      "provider_id": new FormControl("provider_id", Validators.required)


    })


  }
purchase(){





  const {id,email} = JSON.parse(this.backend.getData("data")!)



  let data = new OrderForm(this.orderForm.value.currency,this.orderForm.value.amount,this.orderForm.value.bank_account,id,this.orderForm.value.provider_id)
  let update = new ExchangeForm(email,"GHS",this.orderForm.value.currency,this.orderForm.value.amount,)

  this.backend.orderList.push(data)
  
  this.backend.placeOrder(data).subscribe((result) => {console.log(result)})


  this.backend.updateOrder(update).subscribe((result) => {console.log(result)})
}

}










