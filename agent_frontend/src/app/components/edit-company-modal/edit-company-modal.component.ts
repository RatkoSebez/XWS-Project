import { Component, ElementRef, Input, OnInit, ViewChild } from '@angular/core';
import { Company } from 'src/app/model/Company';
import { RegisterCompanyDto } from 'src/app/model/dto/RegisterCompanyDto';
import { CompanyService } from 'src/app/services/company.service';

@Component({
  selector: 'app-edit-company-modal',
  templateUrl: './edit-company-modal.component.html',
  styleUrls: ['./edit-company-modal.component.css']
})
export class EditCompanyModalComponent implements OnInit {

  @Input() company!: Company;
  @ViewChild('address') address!: ElementRef;
  @ViewChild('phoneNumber') phoneNumber!: ElementRef;
  @ViewChild('email') email!: ElementRef;
  @ViewChild('description') description!: ElementRef;

  constructor(private companyService: CompanyService) { }

  ngOnInit(): void {
  }

  edit(company: Company){
    company.address = this.address.nativeElement.value
    company.phoneNumber = this.phoneNumber.nativeElement.value
    company.description = this.description.nativeElement.value
    company.email = this.email.nativeElement.value
    // var company = new Company(name, address, email, phoneNumber, description, ownerEmail)
    this.companyService.edit(this.company).subscribe(response => {
      // this.getCompanies()
    })
  }
}
