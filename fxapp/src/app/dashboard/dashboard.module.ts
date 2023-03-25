import { DashboardComponent } from './dashboard.component';

import { NgModule } from "@angular/core";
import { OperationCardComponent } from './operation-card/operation-card.component';
import { OrdersCardComponent } from './orders-card/orders-card.component';
import { WithdrawCardComponent } from './withdraw-card/withdraw-card.component';

@NgModule({
  imports:
  [WalletModule,
    

  ],
  declarations: [DashboardComponent, OperationCardComponent, OrdersCardComponent, WithdrawCardComponent],
  exports: [DashboardComponent],
})
export class WalletModule {}
