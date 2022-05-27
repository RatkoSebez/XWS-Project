import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Router } from '@angular/router';
import { AuthRequestDTO } from '../dtos/AuthRequestDTO';
import { AuthResponseDTO } from '../dtos/AuthResponseDTO';



@Injectable({
  providedIn: 'root'
})

export class LoginService{

    constructor(private http: HttpClient, private route: Router) {}
    url = 'http://localhost:8080/' + 'login';
    user = new AuthResponseDTO('', '', '');

    login(authRequest:AuthRequestDTO){  
    return this.http.post<any>(this.url,authRequest);
    }

    setUserData(data:AuthResponseDTO){
        this.user.token=data.token;
        localStorage.setItem('activeUser', JSON.stringify(this.user))
        localStorage.setItem('token', data.token)
        localStorage.setItem('mail', data.email);
        localStorage.setItem('logedIn', 'true')
      }

    getHeaders() {
        const jwt = localStorage.getItem('token')
        const headers = new HttpHeaders({
          'Content-Type': 'application/json',
          Authorization: `Bearer ` + jwt,
        });
        return headers;
      }

}