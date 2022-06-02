import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { UserDTO } from '../dtos/UserDTO';
import { ProfileEditService } from '../services/profile-edit.service';
import { Router } from '@angular/router';
import { ToastrService } from 'ngx-toastr';

@Component({
  selector: 'app-profile-edit-page',
  templateUrl: './profile-edit-page.component.html',
  styleUrls: ['./profile-edit-page.component.css']
})
export class ProfileEditPageComponent implements OnInit {

  public user$ = new UserDTO("", "", "", "", "", 0,0,0,"", "", "", [], [], [], [], true);
  public user: any  
      public exp1 = ''
      public exp2 = ''
      public exp3 = ''
      public expList = new Array<string>();
      public edu1 = ''
      public edu2 = ''
      public edu3 = ''
      public eduList = new Array<string>();
      public int1 = ''
      public int2 = ''
      public int3 = ''
      public interestList = new Array<string>();
      public skill1 = ''
      public skill2 = ''
      public skill3 = ''
      public skillList = new Array<string>();
      public x :any
      public privatetext = 'Make profile public'
 
  constructor(
    public service: ProfileEditService,
    private route: Router,
    private toastr:ToastrService,

  ) { }

  ngOnInit(): void {
    this.service.getUserByMail(localStorage.getItem('mail'))
      .subscribe(data => this.user$ = data);
      this.user$.isprivate=true
  }

  onSubmit(){
    this.addEducation();
    this.addExperience();
    this.addInterest();
    this.addSkill();
    this.service.editUser(this.user$).subscribe(
      res=>{
        
        this.toastr.success('Updated successfully', 'Success');
      }, err => {        this.toastr.error('Error!', 'Error');
    }
      
    );
    this.route.navigate(['/profile'])

  }

  addExperience(): void{
    if (this.user$.experience != null) {
    this.user$.experience.forEach(element => {
      if (element == '') {this.user$.experience.splice(this.user$.experience.indexOf(element),1)}
    });}
    if (this.user$.experience == null){
      this.user$.experience = []
    }
    if (this.exp1 != '')
    this.user$.experience.push(this.exp1)
    this.exp1 = ""
  }
  addEducation(): void{
    if (this.user$.education != null){
    this.user$.education.forEach(element => {
      if (element == '') {this.user$.education.splice(this.user$.education.indexOf(element),1)}
    });}
    if (this.user$.education == null){
      this.user$.education = []
    }
    if (this.edu1 != '')
    this.user$.education.push(this.edu1)
    this.edu1 = ""
  }
  addInterest(): void{
    if (this.user$.interests!= null){
    this.user$.interests.forEach(element => {
      if (element == '') {this.user$.interests.splice(this.user$.interests.indexOf(element),1)}
    });}
    if (this.user$.interests == null){
      this.user$.interests = []
    }
    if (this.int1 != '')
    this.user$.interests.push(this.int1)
    this.int1 = ""
  }
  addSkill(): void{
    if (this.user$.skills != null){
    this.user$.skills.forEach(element => {
      if (element == '') {this.user$.skills.splice(this.user$.skills.indexOf(element),1)}
    });}
    if (this.user$.skills == null){
      this.user$.skills = []
    }
    if (this.skill1 != '')
    this.user$.skills.push(this.skill1)
    this.skill1 = ""
  }
  profilePublic(){
    this.user$.isprivate = false
  }
  profilePrivate(){
    this.user$.isprivate = !this.user$.isprivate
    if (this.user$.isprivate == true){
      this.privatetext = 'Make profile public'
    }else if (this.user$.isprivate == false){
      this.privatetext = 'Make profile private'
    }
    this.service.editUser(this.user$).subscribe(
      res=>{
        
        this.toastr.success('Updated successfully', 'Success');
      }, err => {        this.toastr.error('Error!', 'Error');
    }
      
    );
  }
}
