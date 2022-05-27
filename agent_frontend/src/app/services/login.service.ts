import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { JwtData } from '../model/dto/JwtData';
import { LoginUserDto } from '../model/dto/LoginUserDto';

@Injectable({
  providedIn: 'root'
})
export class LoginService {

  constructor(private http: HttpClient) { }

  login(userDto: LoginUserDto){
    return this.http.post('http://localhost:8090/api/login', userDto)
  }

  logout(){
    localStorage.clear()
  }

  isLoggedIn(){
    return localStorage.getItem('token') !== null
  }

  saveToLocalStorage(jwtData: JwtData){
    localStorage.setItem('token', jwtData.token)
    localStorage.setItem('email', jwtData.email)
    localStorage.setItem('username', jwtData.username)
  }
}
