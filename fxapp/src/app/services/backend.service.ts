import { Router } from '@angular/router';
import { OrderDto } from './../types/OrderDto';
import { ExchangeForm } from './../types/ExchangeForm';
import { OrderForm } from './../types/OrderForm';
import { Injectable } from '@angular/core';
import { catchError, Observable, tap } from 'rxjs';
import {
  HttpClient,
  HttpErrorResponse,
  HttpHeaders,
  HttpParams,
} from '@angular/common/http';
import { ServerDomainService } from './server-domain.service';
import { Order } from '../types/Order';

@Injectable({
  providedIn: 'root'
})
export class BackendService {
  serverDomain: string;


  constructor(
    private http: HttpClient,
    private httpDomain: ServerDomainService

  ) { this.serverDomain = this.httpDomain.getDomain(); }



  submitSignUp(form: any): Observable<any> {
    return this.http
      .post(`${this.serverDomain}/api/auth/register`, form)
      ;
  }

  orderHelper(){
    const {orders} = JSON.parse(localStorage.getItem("data")!)
    return orders ? orders:[]
  }

  public placeOrder(form:OrderForm) :Observable<OrderForm>{

    return this.http
    .post<OrderForm>(`${this.serverDomain}/api/client/purchase`, form);

  }

  public updateOrder(form:ExchangeForm) :Observable<OrderForm>{

    return this.http
    .post<OrderForm>(`${this.serverDomain}/api/client/exchange`, form);

  }



   orderList:Array<OrderForm>=[]




  public submitSignIn(form: any): Observable<any> {
    return this.http
      .post(`${this.serverDomain}/api/auth/login`, form)

  }






// Local Storage Service


  public saveData(key: string, value: string,router :Router) {

    localStorage.setItem(key, value);
    router.navigate(['/dashboard'])
  }

  public getData(key: string) {
    return localStorage.getItem(key)
  }
  public removeData(key: string) {
    localStorage.removeItem(key);
  }

  public clearData() {
    localStorage.clear();
  }





}
