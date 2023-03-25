import { DashboardComponent } from './dashboard.component';

import { NgModule } from "@angular/core";

@NgModule({
  imports:
  [WalletModule,

  ],
  declarations: [DashboardComponent],
  exports: [DashboardComponent],
})
export class WalletModule {}
