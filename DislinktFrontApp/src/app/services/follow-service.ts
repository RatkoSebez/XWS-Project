import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Router } from '@angular/router';
import { NewPostDTO } from '../dtos/NewPostDTO';
import { LoginService } from './login-service';
import { CommentDTO } from '../dtos/CommentDTO';



@Injectable({
  providedIn: 'root'
})

export class FollowService{
    constructor(private http: HttpClient, private route: Router, private loginService:LoginService) {}
    url = 'http://localhost:8080/'
    url1 = 'http://localhost:8080//follow'

  
    follow( data:any){
        const headers = this.loginService.getHeaders();
        return this.http.post<any>(this.url+'follow/'+data.email, data, {headers:headers})
    }
    followRequest(data:any){
        const headers = this.loginService.getHeaders();
        return this.http.post<any>(this.url+'followRequest/'+data.email, data, {headers:headers})
    }
    getFollowers(mail:string){
        const headers = this.loginService.getHeaders();
        return this.http.get<any>(this.url+'follow/'+mail,  {headers:headers})

    }
    getRequests(mail:string){
        const headers = this.loginService.getHeaders();
        return this.http.get<any>(this.url+'followRequest/'+mail,  {headers:headers})
    }
}