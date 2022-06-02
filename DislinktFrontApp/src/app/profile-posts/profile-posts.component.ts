import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ProfileEditService } from '../services/profile-edit.service';
import { PostService } from '../services/post-service';
import { FileUploadService } from '../services/file-upload-service';
import { PostDTO } from '../dtos/PostDTO';
import { Post } from '../model/post';
import { NewPostDTO } from '../dtos/NewPostDTO';
import { Observable } from 'rxjs';
import { HttpClient } from '@angular/common/http';
import { LoginService } from '../services/login-service';
import { LinkyModule } from 'angular-linky';
import { ReactionDTO } from '../dtos/ReactionDTO';
import { CommentDTO } from '../dtos/CommentDTO';
import { FollowService } from '../services/follow-service';
import { FollowRequestDTO } from '../dtos/FollowRequestDTO';
import { ThrowStmt } from '@angular/compiler';

@Component({
  selector: 'app-profile-posts',
  templateUrl: './profile-posts.component.html',
  styleUrls: ['./profile-posts.component.css']
})
export class ProfilePostsComponent implements OnInit {
  header : any
  constructor(private followService:FollowService,  private loginservice:LoginService, private http:HttpClient,public service:ProfileEditService,public route:Router,private fileUploadService: FileUploadService, private postService: PostService) { }
  mail : any
  public posts$ : PostDTO[] = []
  public  x : any
  public likenum : Array<number> = []
  public likenumber:number = 0
  public dislikenumber:number=0
  public proba :number = 0
  public dislikenum: Array<number> = []
  public commentsnum: Array<number> = []
  public commentText : any
  public cdto: any
  public comnum: number = 0
  public commentsSelected : boolean = false
  public xxx : number = 0
  public yyy : number = 0
  public zzz : number = 0
  public k : number = 0
  public comindex :number = 0
  public alreadyLiked : boolean = false
  public postsliked: Array<boolean> = []
  public newpostactive:boolean = false
  public myprofile:boolean = false
  static user$ : any
  public followreq:any
  public followbtn: string = ''
  public privateFollowbtn: string = ''
  public followers:boolean = false
  public following:boolean = false

  ngOnInit(): void {
  //this.posts$ = this.service.getUserPosts()
  this.followbtn = 'Follow'
  this.mail = localStorage.getItem('mail')
  this.isThisMyProfile()
  this.header = this.loginservice.getHeaders()

  this.http.get<PostDTO[]>('http://localhost:8080/get-user-posts', {'headers':this.header}).subscribe(data=>{
  this.posts$=data; console.log(data); this.x=data
    })
  console.log('postovi:')
    console.log(this.posts$)
  
    
    }
     reaction: ReactionDTO={ postID:"", userEmail:"",react: ""}


  like(id:string, index:number){
    this.mail = localStorage.getItem('mail')
    console.log('like')
    var id1 = id.split('"')[1]
    console.log('mail ' + this.mail)

    console.log(id1)
    //this.reaction = new ReactionDTO("",id1, this.mail, "", "LIKE")
    this.reaction.postID = id1; 
    this.reaction.userEmail = this.mail;
    this.reaction.react = "LIKE"
    this.service.makeReaction(this.reaction).subscribe(
      res=>{
        
        //this.toastr.success('Registered successfully', 'Success')
       console.log('success');
        ;
      }, err => {   console.log('error')     ;
    }
    );;
    this.likeplus()
    this.likenum[index] = this.likenum[index] + 1
    this.alreadyLiked = true
    this.postsliked[index] = true
    console.log(id)
    this.ngOnInit()
  
}
  dislike(id:string, index:number){
    console.log('disliek')
    this.mail = localStorage.getItem('mail')
    var id1 = id.split('"')[1]
    console.log('mail ' + this.mail)
    //this.reaction = new ReactionDTO("",id1, this.mail, "", "DISLIKE")
    this.reaction.postID = id1; 
    this.reaction.userEmail = this.mail;

    this.reaction.react = "DISLIKE"
    this.http.post<PostDTO>('http://localhost:8080/make-reaction', (this.reaction), {headers:this.header}).subscribe(res=>console.log(res))
    // this.service.makeReaction(this.reaction).subscribe(
    //   (res)=>{
        
    //     //this.toastr.success('Registered successfully', 'Success')
    //    console.log('success');
    //    console.log(res)
    //     ;
    //   }, err => {   console.log('error')     ;
    // }
      
    // );;
    this.reaction = new ReactionDTO("","", "" )
    this.dislikeplus()
    this.dislikenum[index] = this.dislikenum[index] + 1
    this.ngOnInit()

  }
  getPostID(id:string):string{
    var id1 = id.split("'")
    console.log(id[1])
    return id1[1]
    
  }
likeplus(){
  this.likenumber = this.likenumber+1;
  ++this.k 
}
complus(){
  this.comnum = this.comnum+1;
}
likereset(){
  this.likenum.push(this.likenumber)
  this.dislikenum.push(this.dislikenumber)
  this.likenumber=0
  this.dislikenumber=0
  
}
dislikeplus(){
  this.dislikenumber = this.dislikenumber+1;
}
comment(id:string, index:number){
  var id1 = id.split('"')[1]
  this.cdto = new CommentDTO(id1, this.commentText)
  this.postService.makeComment(this.cdto).subscribe(res=>console.log(res))
  this.commentText = ''
  this.commentsnum[index] = this.commentsnum[index] + 1
  this.complus()

  this.ngOnInit()

}
comments(index:number){
  this.commentsSelected=!(this.commentsSelected)
  this.comindex = index
}
resetcom(){
  this.commentsnum.push(this.comnum)

  
}
getlikes(){
  this.xxx = this.likenumber
}
getdislikes(){
 this.yyy = this.dislikenumber
}
getcomments(){
  this.zzz = this.comnum
}

newpost(){
  this.newpostactive=true
}
isThisMyProfile(){
  if (this.route.url.split('/')[1] == this.mail)
    this.myprofile = true;
}
public userEmail:string = ''
public user2$ : any
getUserEmail(){
  this.userEmail =  ProfilePostsComponent.user$.email
  this.user2$ = ProfilePostsComponent.user$
}
follow(){
  this.followreq = new FollowRequestDTO(this.mail, this.user2$.email )
  this.followService.follow(this.followreq).subscribe(
    res=>{
      
      //this.toastr.success('Registered successfully', 'Success')
     console.log('success');
     this.followbtn = 'Following'

      ;
    }, err => {   console.log('error')     ;
  }
  );;
}
followRequest(){
  
}
profileInfo(){
  this.route.navigate(['/profile'])
}
public follows:any

showFollowers(){
  this.following = false
  this.followers = true;
  this.followService.getFollowers(this.mail).subscribe(data=>{
    this.follows = data;
      })
}


showFollowing(){
  this.followers = false
  this.following = true;
  this.followService.getFollowers(this.mail).subscribe(data=>{
    this.follows = data;
      })
}
}
