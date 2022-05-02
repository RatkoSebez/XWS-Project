import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Router } from '@angular/router';
import { AuthRequestDTO } from '../dtos/AuthRequestDTO';
import { AuthResponseDTO } from '../dtos/AuthResponseDTO';



@Injectable({
  providedIn: 'root'
})

export class RegistraionService{

  constructor(private http:HttpClient) { }

  url = 'http://localhost:8081/' + 'register';
  RegisterUser(data:any){
    return this.http.post(this.url , data )
  }
 

    // constructor(private http: HttpClient, private route: Router) {}
    // url = 'http://localhost:8081/' + 'register';
    // user = new AuthResponseDTO('', '','','','',0,0,0,'','','',{},{},{},{}, true);

    // login(authRequest:AuthRequestDTO){  
    // return this.http.post<any>(this.url,authRequest);
    // }

    // setUserData(data:AuthResponseDTO){
    //     this.user=data;
    //     localStorage.setItem('newUser', JSON.stringify(this.user))
    //     localStorage.setItem('token', this.user.token)
    //     localStorage.setItem('mail', this.user.email);
    //     localStorage.setItem('logedIn', 'true')
    //   }


}