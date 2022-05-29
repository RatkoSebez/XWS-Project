import { Component, OnInit } from '@angular/core';
import { Company } from 'src/app/model/Company';
import { CreateJobDTO } from 'src/app/model/dto/CreateJobDTO';
import { RegisterCompanyDto } from 'src/app/model/dto/RegisterCompanyDto';
import { Job } from 'src/app/model/Job';
import { CompanyService } from 'src/app/services/company.service';

@Component({
  selector: 'app-company',
  templateUrl: './company.component.html',
  styleUrls: ['./company.component.css']
})
export class CompanyComponent implements OnInit {
  companies!: Company[]
  companyName = ''
  // this is company that I am sending to child component (company edit modal)
  company = new Company('', '', '', '', '', false, '', [], [], [], [])

  constructor(private companyService: CompanyService) { }

  ngOnInit(): void {
    this.getCompanies()
  }

  setCompany(company: Company){
    this.company = company
  }

  setCompanyName(companyName: string){
    this.companyName = companyName
  }

  getCompanies(){
    this.companyService.getAll().subscribe(response => {
      this.companies = JSON.parse(JSON.stringify(response))
      this.companies = this.companies.filter(company => company.isApproved == true && company.ownerEmail == localStorage.getItem('email'))
    })
  }
}
