import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, NgForm } from '@angular/forms';
import { Router } from '@angular/router';
import { UserDTO } from '../dtos/UserDTO';
import { RegistraionService } from '../services/registration-service';
import { ToastrService } from 'ngx-toastr';
import { RouterModule } from '@angular/router';
import { FormsModule } from '@angular/forms';


@Component({
  selector: 'app-registration-page',
  templateUrl: './registration-page.component.html',
  styleUrls: ['./registration-page.component.css']
})



export class RegistrationPageComponent implements OnInit {
  

  constructor(
   // private fb:FormBuilder,
    public service:RegistraionService,
    
    private toastr:ToastrService,
   // private http : HttpClient,
    private route: Router,

  ) { }
    public user = new UserDTO('','','','','',0,0,0,'','','',[],[],[],[],true)
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

  ngOnInit(): void {
    
  }


  addExperience(): void{
    if (this.exp1 != '')
      this.expList.push(this.exp1)
    if (this.exp2 != '')
      this.expList.push(this.exp2)
    if (this.exp3 != '')
      this.expList.push(this.exp3)
  }
  addEducation(): void{
    if (this.edu1 != '')
      this.eduList.push(this.edu1)
    if (this.edu2 != '')
      this.eduList.push(this.edu2)
    if (this.edu3 != '')
      this.eduList.push(this.edu3)
  }
  addInterest(): void{
    if (this.int1 != '')
      this.interestList.push(this.int1)
    if (this.int2 != '')
      this.interestList.push(this.int2)
    if (this.int3 != '')
      this.interestList.push(this.int3)
  }
  addSkill(): void{
    if (this.skill1 != '')
      this.skillList.push(this.skill1)
    if (this.skill2 != '')
      this.skillList.push(this.skill2)
    if (this.skill3 != '')
      this.skillList.push(this.skill3)
  }

  onSubmit(){
    this.addEducation();
    this.addExperience();
    this.addInterest();
    this.addSkill();
    this.user.skills = this.skillList
    this.user.education = this.eduList
    this.user.experience = this.expList
    this.user.interests = this.interestList
    this.service.RegisterUser(this.user).subscribe(
      res=>{
        
        this.toastr.success('Registered successfully', 'Success')
        this.route.navigate(['/login']);
        ;
      }, err => {        this.toastr.error('Error!', 'Error');
    }
      
    );;
    
  }

}
