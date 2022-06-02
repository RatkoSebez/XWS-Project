import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders, HttpResponse } from '@angular/common/http';
import { Router } from '@angular/router';
import { LoginService } from '../services/login-service';
import { AuthRequestDTO } from '../dtos/AuthRequestDTO';
import { AuthResponseDTO } from '../dtos/AuthResponseDTO';
import {HttpClientModule} from '@angular/common/http';
import { AppComponent } from '../app.component';
@Component({
  selector: 'app-login-page',
  templateUrl: './login-page.component.html',
  styleUrls: ['./login-page.component.css']
})
export class LoginPageComponent implements OnInit {

  public mail : any;
  public tokenString : any;
  public header : HttpHeaders | undefined
  constructor(private loginService: LoginService, private route: Router) { 
    this.authenticationRequest = new AuthRequestDTO("","");
  }
  authenticationRequest:AuthRequestDTO;
  errorMessage='';

  login(){
    if (this.authenticationRequest.email == '' || this.authenticationRequest.password == '') {
      this.errorMessage = 'Email or password missing.';
    } else {
      this.loginService.login(this.authenticationRequest).subscribe(
        (data) => this.successfulLogin(data), 
        (res) => (this.errorMessage = 'Invalid email or password.')
      ); 
    }
  }

  successfulLogin(data:AuthResponseDTO) {
    this.errorMessage = '';
    console.log(data)
    //console.log(data.token)
    //this.loginService.setUserData(data);
    this.tokenString = JSON.stringify(data).split('"token":"')[1].split('"')[0]
    localStorage.setItem('token', this.tokenString)
    localStorage.setItem('mail', this.authenticationRequest.email)
    localStorage.setItem('isLogged','yes')
    
    AppComponent.logged = true
    this.mail = localStorage.getItem('mail')
    console.log(localStorage.getItem('mail'))
    this.route.navigate(['/'+this.mail]);
  }

  ngOnInit(): void {
  }

}
