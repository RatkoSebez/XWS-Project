import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Company } from '../model/Company';
import { RegisterCompanyDto } from '../model/dto/RegisterCompanyDto';

@Injectable({
  providedIn: 'root'
})
export class CompanyService {

  constructor(private http: HttpClient) { }

  register(dto: RegisterCompanyDto){
    return this.http.post(
      'http://localhost:8090/api/company',
       {name: dto.name, address: dto.address, email: dto.email, phoneNumber: dto.phoneNumber, description: dto.description, ownerEmail: dto.ownerEmail}
    )
  }

  getAll(){
    return this.http.get(
      'http://localhost:8090/api/company'
    )
  }

  approve(name: String){
    return this.http.post(
      'http://localhost:8090/api/company/approve',
       {name: name}
    )
  }

  edit(dto: Company){
    return this.http.put(
      'http://localhost:8090/api/company',
       {name: dto.name, address: dto.address, email: dto.email, phoneNumber: dto.phoneNumber, description: dto.description, ownerEmail: dto.ownerEmail}
    )
  }
}
