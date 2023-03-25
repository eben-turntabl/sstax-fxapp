import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { BackendService } from 'src/app/services/backend.service';
import { LoginForm } from 'src/app/types/LoginForm';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  constructor(private backend:BackendService,private router: Router){

  }


  loginForm!: FormGroup;

  ngOnInit(): void {

    this.loginForm = new FormGroup({

      "email": new FormControl("user@email.com", Validators.required),
      "password": new FormControl("password", Validators.required),



    })





}
submit() {
  let data = new LoginForm(this.loginForm.value.email, this.loginForm.value.password)
  this.backend.submitSignIn(data).subscribe((result) => {this.backend.saveData("data",JSON.stringify(result.data, null, "\t"),this.router)})



}

}
