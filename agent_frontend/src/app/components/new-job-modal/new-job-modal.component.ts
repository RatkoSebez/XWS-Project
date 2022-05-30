import { Component, Input, OnInit } from '@angular/core';
import { Company } from 'src/app/model/Company';
import { CreateJobDTO } from 'src/app/model/dto/CreateJobDTO';
import { Job } from 'src/app/model/Job';
import { CompanyService } from 'src/app/services/company.service';

@Component({
  selector: 'app-new-job-modal',
  templateUrl: './new-job-modal.component.html',
  styleUrls: ['./new-job-modal.component.css']
})
export class NewJobModalComponent implements OnInit {

  @Input() companyName = ''
  job = new Job('', '', '', '')

  constructor(private companyService: CompanyService) { }

  ngOnInit(): void {
  }

  newJob(){
    this.companyService.job(new CreateJobDTO(this.companyName, this.job.name, this.job.description, this.job.requirements, this.job.benefits)).subscribe()
  }
}
