import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './components/login/login.component';
import { RegisterComponent } from './components/register/register.component';
import { DashboardNavlinkComponent } from './dashboard/dashboard-navlink/dashboard-navlink.component';

const routes: Routes = [


{path: 'register', component:RegisterComponent},
{path: 'login', component:LoginComponent},
{path: 'dashboard', component:DashboardNavlinkComponent}

];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
