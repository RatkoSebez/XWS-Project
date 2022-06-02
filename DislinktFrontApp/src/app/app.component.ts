import { Component } from '@angular/core';
import { ProfileEditService } from './services/profile-edit.service';
import { Router } from '@angular/router';
import { ProfilePostsComponent } from './profile-posts/profile-posts.component';
import { ThrowStmt } from '@angular/compiler';
@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  constructor(private service:ProfileEditService, private route:Router){}

  public logged:boolean = false

  public notlogged:boolean=true;
  loading:boolean=false;
  public  logVisible = true;
  public  outvisible = false;
  static lv = true;
  static ov = false;
  public mail : string = ''
  public user$:any
  static logged : boolean = false
  logOut(){
    localStorage.setItem('activeUser', "")
        localStorage.setItem('token', "")
        localStorage.setItem('mail', "");
        localStorage.setItem('logedIn', 'no')
        AppComponent.logged=false
        this.logged = false
        this.getlogged()
        this.ngOnInit()
        
  }
  ngOnInit(){
    this.getlogged()

}

  getUsername(){
    
  }
getlogged(){
  this.logged = AppComponent.logged
}
  search(){
    this.service.getUserByMail(this.mail)
    .subscribe(data => this.user$ = data);
    console.log(this.mail)
    this.route.navigate(['/'+this.mail]);
    this.mail=''
    ProfilePostsComponent.user$ = this.user$
  }
}


class ImageSnippet {
  constructor(public src: string, public file: File) {}
}