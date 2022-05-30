import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { RegisterUserDto } from 'src/app/model/dto/RegisterUserDto';
import { RegisterService } from 'src/app/services/register.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {
  name = ''
  surname = ''
  email = ''
  password = ''
  username = ''

  constructor(private router: Router, private registerService: RegisterService) { }

  ngOnInit(): void {
    // this.http.get('http://localhost:8090/api/test').subscribe(
    //   response => {
    //     console.log(response)
    //   }
    // );

    // this.http.post('http://localhost:8090/api/login', {email: "dusan@gmail.com", password: "dulecar"}).subscribe(
    //   response => {
    //     console.log(response)
    //   }
    // );
  }

  register(){
    var userDto = new RegisterUserDto(this.name, this.surname, this.email, this.username, this.password)
    this.registerService.register(userDto).subscribe(response => {
      this.router.navigate(['/']);
    })
  }
}
