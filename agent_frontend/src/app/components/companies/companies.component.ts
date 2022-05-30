import { Component, OnInit } from '@angular/core';
import { Company } from 'src/app/model/Company';
import { CompanyService } from 'src/app/services/company.service';

@Component({
  selector: 'app-companies',
  templateUrl: './companies.component.html',
  styleUrls: ['./companies.component.css']
})
export class CompaniesComponent implements OnInit {
  companies!: Company[]

  constructor(private companyService: CompanyService) { }

  ngOnInit(): void {
    this.getCompanies()
  }

  getCompanies(){
    this.companyService.getAll().subscribe(response => {
      this.companies = JSON.parse(JSON.stringify(response))
    })
  }

  approve(name: String){
    this.companyService.approve(name).subscribe(response => {
      this.getCompanies()
    })
  }
}
