import { WithdrawCardComponent } from './withdraw-card.component';

import { CardModule,GridModule  } from '@coreui/angular';
import { NgModule } from "@angular/core";




@NgModule({
  imports:
  [CardModule,
    GridModule ,

  ],
  declarations: [WithdrawCardComponent],
  exports: [WithdrawCardComponent],
})
export class WithdrawCardModule {}
