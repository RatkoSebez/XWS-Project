import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { RegisterCompanyDto } from 'src/app/model/dto/RegisterCompanyDto';
import { CompanyService } from 'src/app/services/company.service';

@Component({
  selector: 'app-register-company',
  templateUrl: './register-company.component.html',
  styleUrls: ['./register-company.component.css']
})
export class RegisterCompanyComponent implements OnInit {
  name = ''
  address = ''
  email = ''
  description = ''
  phoneNumber = ''

  constructor(private router: Router, private companyService: CompanyService) { }

  ngOnInit(): void {}

  register(){
    var ownerEmail = JSON.parse(JSON.stringify(localStorage.getItem('email')) || '')
    console.log(ownerEmail)
    var dto = new RegisterCompanyDto(this.name, this.address, this.email, this.phoneNumber, this.description, ownerEmail)
    this.companyService.register(dto).subscribe(response => {
      // this.router.navigate(['/']);
    })
  }
}
