import { SignUpForm } from './../../types/SignUpForm';
import { Component } from '@angular/core';
import { RegisterForm } from 'src/app/types/RegisterForm';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent {




  form:  SignUpForm= {
    email: '',
    password: '',
    f_name: '',
    l_name: '',
    confirmPassword: ''
  };




submit() {
  const registerForm: RegisterForm = {
    first_name: this.form.f_name,
    last_name: this.form.l_name,
    email: this.form.email,
    password: this.form.password,
  };


  console.log(registerForm)
  }

}
