import { Time } from "@angular/common";


export interface OrderDto {

  id: string;
  currency:string;
  placed_at:Time;
  amount:number;
  status:string
  bank_account:string
  client_id:string;
  provider_id:string;
}

