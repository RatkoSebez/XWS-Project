import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Router } from '@angular/router';
import { AuthRequestDTO } from '../dtos/AuthRequestDTO';
import { AuthResponseDTO } from '../dtos/AuthResponseDTO';
import { Observable } from 'rxjs';
import { UserDTO } from '../dtos/UserDTO';
import { LoginService } from './login-service';
import { PostDTO } from '../dtos/PostDTO';
import { Post } from '../model/post';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/catch';
import { JsonpInterceptor } from '@angular/common/http';
import { JsonPipe } from '@angular/common';
import { ReactionDTO } from '../dtos/ReactionDTO';


@Injectable({
  providedIn: 'root'
})

export class ProfileEditService{

    public mail:any
    public header : HttpHeaders | undefined
    constructor(private http: HttpClient, private route: Router, private loginservice: LoginService) {}
    url = 'http://localhost:8080/l';         //+ '${mail}';
    urlDefault = 'http://localhost:8080/';   
    //user = new AuthResponseDTO('', '');

    getUser():Observable<UserDTO>{
      this.header = this.loginservice.getHeaders()

      return this.http.get<UserDTO>(this.url, {'headers':this.header})
    }
    editUser(data:any){
      this.mail = localStorage.getItem('mail')
      this.header = this.loginservice.getHeaders()

      return this.http.put(this.urlDefault + this.mail , data, {'headers':this.header})
    }
    getFollowed(data:any){
      return this.http
    }
    getUserByMail(mail:any):Observable<UserDTO>{
      this.header = this.loginservice.getHeaders()
      this.mail = localStorage.getItem('mail')
      console.log(this.header)
      console.log(this.mail)
      return this.http.get<UserDTO>(this.urlDefault + mail, {'headers':this.header})
    }
    makeReaction(data:ReactionDTO){
      this.header = this.loginservice.getHeaders()
      var body = JSON.stringify(data)
      console.log('ovo se kreira ' + body)
      return this.http.post(this.urlDefault + 'make-reaction', body, {headers:this.header})
    }
    
}