import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Router } from '@angular/router';
import { NewPostDTO } from '../dtos/NewPostDTO';
import { LoginService } from './login-service';



@Injectable({
  providedIn: 'root'
})

export class PostService{
    constructor(private http: HttpClient, private route: Router, private loginService:LoginService) {}
    url = 'http://localhost:8082/'

    makeNewPost(post:NewPostDTO){
        const headers = this.loginService.getHeaders();
        const body=JSON.stringify(post) 
        return this.http.post<any>(this.url+'make-new-post',body,{headers:headers})
    }
}