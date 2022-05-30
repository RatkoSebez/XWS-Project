import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { JwtData } from 'src/app/model/dto/JwtData';
import { LoginUserDto } from 'src/app/model/dto/LoginUserDto';
import { LoginService } from 'src/app/services/login.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  email = ''
  password = ''

  constructor(private router: Router, private loginService: LoginService) { }

  ngOnInit(): void {
  }

  login(){
    var userDto = new LoginUserDto(this.email, this.password)
    this.loginService.login(userDto).subscribe(response => {
      var jwtData: JwtData = JSON.parse(JSON.stringify(response))
      this.loginService.saveToLocalStorage(jwtData)
      this.router.navigate(['/']);
    })
  }
}
