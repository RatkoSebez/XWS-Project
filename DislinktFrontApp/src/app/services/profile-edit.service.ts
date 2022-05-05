import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Router } from '@angular/router';
import { AuthRequestDTO } from '../dtos/AuthRequestDTO';
import { AuthResponseDTO } from '../dtos/AuthResponseDTO';
import { Observable } from 'rxjs';
import { UserDTO } from '../dtos/UserDTO';



@Injectable({
  providedIn: 'root'
})

export class ProfileEditService{

    constructor(private http: HttpClient, private route: Router) {}
    mail = localStorage.getItem('mail') 
    url = 'http://localhost:8081/saasd';         //+ '${mail}';
    //user = new AuthResponseDTO('', '');

    getUser():Observable<UserDTO>{
      return this.http.get<UserDTO>(this.url)
    }

}