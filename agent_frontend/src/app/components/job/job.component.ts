import { Component, Input, OnInit } from '@angular/core';
import { Job } from 'src/app/model/Job';

@Component({
  selector: 'app-job',
  templateUrl: './job.component.html',
  styleUrls: ['./job.component.css']
})
export class JobComponent implements OnInit {

  @Input() job = new Job('', '', '', '')

  constructor() { }

  ngOnInit(): void {
  }

}
