import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { UserDTO } from '../dtos/UserDTO';
import { ProfileEditService } from '../services/profile-edit.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-profile-edit-page',
  templateUrl: './profile-edit-page.component.html',
  styleUrls: ['./profile-edit-page.component.css']
})
export class ProfileEditPageComponent implements OnInit {

  public user$ = new UserDTO("", "", "", "", "", 0,0,0,"", "", "", [], [], [], [], true);
  public user: any  
  public x = 0
  constructor(
    public service: ProfileEditService,
    private route: Router,
  ) { }

  ngOnInit(): void {
    this.service.getUser()
      .subscribe(data => this.user$ = data);
  }

  get(){

  }
}
