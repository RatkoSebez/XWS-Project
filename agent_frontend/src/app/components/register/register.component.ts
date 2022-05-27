import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {
  firstName = ''
  lastName = ''
  email = ''
  password = ''

  constructor(private http: HttpClient, private router: Router) { }

  ngOnInit(): void {
    // this.http.get('localhost:8081/test').subscribe(
    //   response => {
    //     console.log(response)
    //   }
    // );

    this.http.post('localhost:8080/login', {email: "dusan@gmail.com", password: "dulecar"}).subscribe(
      response => {
        console.log(response)
      }
    );
  }

  register(){
    this.http.post('localhost:8090/api/register', {firstName: this.firstName, lastName: this.lastName, username: this.email, password: this.password}).subscribe(
      response => {
      }
    );
    // do not wait for response because server is sending email (which is slow)
    // this.router.navigate(['/']);
  }
}
