import { Currrency } from './../../types/Currency';
import { Component, Input } from '@angular/core';



@Component({
  selector: 'app-wallet',
  templateUrl: './wallet.component.html',
  styleUrls: ['./wallet.component.css']
})
export class WalletComponent {
  @Input()currencyArray!:Currrency

}
