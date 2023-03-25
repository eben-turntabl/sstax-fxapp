import { OperationCardComponent } from './../../dashboard/operation-card/operation-card.component';

import { CardModule,GridModule  } from '@coreui/angular';
import { NgModule } from "@angular/core";




@NgModule({
  imports:
  [CardModule,
    GridModule ,

  ],
  declarations: [OperationCardComponent],
  exports: [OperationCardComponent],
})
export class OperationCardModule {}
