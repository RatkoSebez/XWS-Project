import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { PostService } from '../services/post-service';

@Component({
  selector: 'app-front-page',
  templateUrl: './front-page.component.html',
  styleUrls: ['./front-page.component.css']
})
export class FrontPageComponent implements OnInit {

  public skillList = new Array<string>();
  public posts : any
  public mail : any 
  constructor(private postservice:PostService) { }

  ngOnInit(): void {
    this.mail = localStorage.getItem('mail')
  this.postservice.getFollowingPosts().subscribe(data=>{
    this.posts = data;
      })
  
  }

}
