import { WalletComponent } from './wallet.component';
import { CardModule,GridModule  } from '@coreui/angular';
import { NgModule } from "@angular/core";




@NgModule({
  imports:
  [CardModule,
    GridModule ,

  ],
  declarations: [WalletComponent],
  exports: [WalletComponent],
})
export class WalletModule {}
