import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { User } from 'src/app/model/User';
import { LoginService } from 'src/app/services/login.service';

@Component({
  selector: 'app-menu',
  templateUrl: './menu.component.html',
  styleUrls: ['./menu.component.css']
})
export class MenuComponent implements OnInit {
  public isLoggedIn = false;
  public isAdmin = false;
  public isClient = false;
  user = new User('', 1)

  constructor(private loginService: LoginService, private http: HttpClient) { }
  
  ngOnInit(): void {
    localStorage.setItem('ok', 'false');
    // console.log(localStorage.getItem('ok') + "??");
    // this.http.get<any>('api/user/loggedInUser/').subscribe(
    //   response => {
    //     if(response == null){
    //       this.isLoggedIn = false, this.isAdmin = false;
    //       return
    //     }
    //     localStorage.setItem('ok', 'true');
    //     this.isLoggedIn = true;
    //     this.user = response;
    //     if(this.user.role == "ROLE_ADMIN") this.isAdmin = true;
    //     if(this.user.role == "ROLE_CLIENT") this.isClient = true;
    //   }
    // );
  }
}
