import { Component, Input, OnInit, Output, EventEmitter } from '@angular/core';
import { Company } from 'src/app/model/Company';
import { CompanyService } from 'src/app/services/company.service';

@Component({
  selector: 'app-company-card',
  templateUrl: './company-card.component.html',
  styleUrls: ['./company-card.component.css']
})
export class CompanyCardComponent implements OnInit {

  comment = ''
  salary = 0
  jobPosition = ''
  @Input() company!: Company;
  @Output() refresh = new EventEmitter<string>();

  callParent(): void {
    this.refresh.next('');
  }

  constructor(private companyService: CompanyService) { }

  ngOnInit(): void {
  }

  createComment(){
    // this.callParent()
    this.companyService.comment(this.company.name, this.comment).subscribe()
    this.comment = ''
  }

  createSalary(){
    // this.callParent()
    this.companyService.salary(this.company.name, this.jobPosition, this.salary).subscribe()
    this.salary = 0
    this.jobPosition = ''
  }

}
