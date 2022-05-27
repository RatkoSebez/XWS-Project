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

  url = 'http://localhost:8080/' + 'register';
  RegisterUser(data:any){
    return this.http.post(this.url , data )
  }


}