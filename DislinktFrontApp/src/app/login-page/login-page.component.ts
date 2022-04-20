import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Router } from '@angular/router';
import { LoginService } from '../services/login-service';
import { AuthRequestDTO } from '../dtos/AuthRequestDTO';
import { AuthResponseDTO } from '../dtos/AuthResponseDTO';

@Component({
  selector: 'app-login-page',
  templateUrl: './login-page.component.html',
  styleUrls: ['./login-page.component.css']
})
export class LoginPageComponent implements OnInit {

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
    console.log(data);
    this.loginService.setUserData(data);
  }

  ngOnInit(): void {
  }

}
