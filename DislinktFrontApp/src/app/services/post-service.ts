import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Router } from '@angular/router';
import { NewPostDTO } from '../dtos/NewPostDTO';
import { LoginService } from './login-service';
import { CommentDTO } from '../dtos/CommentDTO';



@Injectable({
  providedIn: 'root'
})

export class PostService{
    constructor(private http: HttpClient, private route: Router, private loginService:LoginService) {}
    url = 'http://localhost:8080/'
    url1 = 'http://localhost:8080/following-posts'

    makeNewPost(post:NewPostDTO){
        const headers = this.loginService.getHeaders();
        const body=JSON.stringify(post) 
        return this.http.post<any>(this.url+'make-new-post',body,{headers:headers})
    }
    getFollowingPosts(){
        const headers = this.loginService.getHeaders();

        return this.http.get<any>(this.url1, {headers:headers})
    }
    makeComment(comm:CommentDTO){
        const headers = this.loginService.getHeaders();

        return this.http.post<any>(this.url+'comment',comm, {headers:headers} )
    }
}