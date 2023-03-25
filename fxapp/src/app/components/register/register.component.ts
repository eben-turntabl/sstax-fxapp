import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Component, OnInit } from '@angular/core';
import { RegisterForm } from 'src/app/types/RegisterForm';
import { BackendService } from 'src/app/services/backend.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit  {
  registerForm!: FormGroup;

  constructor(private backend:BackendService) {}
  ngOnInit(): void {

    this.registerForm = new FormGroup({

      "email": new FormControl("user@email.com", Validators.required),
      "password": new FormControl("0.password", Validators.required),
      "last_name": new FormControl("last_name", Validators.required),
      "first_name": new FormControl("first_name", Validators.required),



    })
  }









submit() {


  let data = new RegisterForm(this.registerForm.value.first_name, this.registerForm.value.last_name,this.registerForm.value.email, this.registerForm.value.password)
  this.backend.submitSignUp(data).subscribe((result) => {console.log(result)})
  console.log(this.registerForm)
  }

}
