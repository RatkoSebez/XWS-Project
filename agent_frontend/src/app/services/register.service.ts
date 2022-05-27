import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { RegisterUserDto } from '../model/dto/RegisterUserDto';

@Injectable({
  providedIn: 'root'
})
export class RegisterService {

  constructor(private http: HttpClient) { }

  register(userDto: RegisterUserDto){
    return this.http.post(
      'http://localhost:8090/api/register',
       {name: userDto.name, surname: userDto.surname, email: userDto.email, password: userDto.password, username: userDto.username}
    )
  }
}
