import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-salary',
  templateUrl: './salary.component.html',
  styleUrls: ['./salary.component.css']
})
export class SalaryComponent implements OnInit {

  @Input() salary = 0
  @Input() jobPosition = ''

  constructor() { }

  ngOnInit(): void {
  }

}
