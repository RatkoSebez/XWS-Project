import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Company } from '../model/Company';
import { CreateJobDTO } from '../model/dto/CreateJobDTO';
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

  comment(name: string, comment: string){
    return this.http.put(
      'http://localhost:8090/api/company/comment',
       {name: name, comment: comment}
    )
  }

  interview(name: string, review: string){
    return this.http.put(
      'http://localhost:8090/api/company/interview',
       {name: name, interviewReview: review}
    )
  }

  salary(name: string, jobPosition: string, salary: number){
    return this.http.put(
      'http://localhost:8090/api/company/salary',
       {name: name, jobPosition: jobPosition, salary: salary}
    )
  }

  job(jobDto: CreateJobDTO){
    console.log('poslao')
    console.log(jobDto.companyName)
    return this.http.put(
      'http://localhost:8090/api/company/job',
       {companyName: jobDto.companyName, name: jobDto.name, description: jobDto.description, requirements: jobDto.requirements, benefits: jobDto.benefits}
    )
  }
}
