import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import * as d3 from 'd3'
import { Company } from 'src/app/model/Company';
import { User } from 'src/app/model/User';
import { CompanyService } from 'src/app/services/company.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {
  user = new User('', 1)
  show = localStorage.getItem('email') != null

  companies!: Company[]

  constructor(private companyService: CompanyService) { }

  ngOnInit(): void {
    this.getCompanies()
  }

  getCompanies(){
    this.companyService.getAll().subscribe(response => {
      this.companies = JSON.parse(JSON.stringify(response))
      this.companies = this.companies.filter(company => company.isApproved)
    })
  }

  refresh(event){
    this.getCompanies()
  }
}