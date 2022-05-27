import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { UserDTO } from '../dtos/UserDTO';
import { ProfileEditService } from '../services/profile-edit.service';

@Component({
  selector: 'app-profile-page',
  templateUrl: './profile-page.component.html',
  styleUrls: ['./profile-page.component.css']
})
export class ProfilePageComponent implements OnInit {
  public userDTO$ : any
  public mail : any
  public token :any
  public user$:any 
  public name : any
  public skills : Array<string> | undefined 
  public education : Array<string> | undefined
  public interests: Array<string> | undefined 
  public experience: any
  public bio: string | undefined 
  constructor(
    public service : ProfileEditService, private route :Router,

  ) { }

  ngOnInit(): void {
    this.mail = localStorage.getItem('mail')
    this.token = localStorage.getItem('token')
    this.service.getUserByMail('${mail}')
      .subscribe(data => this.user$ = data);
    console.log(this.user$)
    this.bio = this.user$.bio
    this.skills = this.user$.skills 
    this.education = this.user$.education
    this.interests = this.user$.interests
    this.experience = this.user$.experience
    console.log(this.experience)
    //this.userDTO$ = new UserDTO(this.user$.name,this.user$.surname , this.user$.email, this.user$.number, this.user$.sex,this.user$.bdateday,this.user$.bdatemonth,this.user$.bdateyear, this.user$.username, this.user$.password, this.user$.bio, this.user$.experience, this.user$.education, this.user$.skills, this.user$.isprivate);

}

edit(){
  this.route.navigate(['/edit']);
}


}
