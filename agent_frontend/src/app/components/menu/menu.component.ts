import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { LoginService } from 'src/app/services/login.service';

@Component({
  selector: 'app-menu',
  templateUrl: './menu.component.html',
  styleUrls: ['./menu.component.css']
})
export class MenuComponent implements OnInit {
  public isLoggedIn = false;
  public isAdmin = false;

  constructor(private loginService: LoginService, private router: Router) {
    this.router.routeReuseStrategy.shouldReuseRoute = () => false;
  }
  
  ngOnInit(): void {
    this.isLoggedIn = this.loginService.isLoggedIn()
    this.isAdmin = this.loginService.isAdmin()
  }

  public logout(){
    this.loginService.logout()
  }
}
