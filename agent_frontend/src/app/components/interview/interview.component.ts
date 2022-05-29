import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-interview',
  templateUrl: './interview.component.html',
  styleUrls: ['./interview.component.css']
})
export class InterviewComponent implements OnInit {

  @Input() interviewReview = ''

  constructor() { }

  ngOnInit(): void {
  }

}
