import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs';

@Component({
  selector: 'app-front-page',
  templateUrl: './front-page.component.html',
  styleUrls: ['./front-page.component.css']
})
export class FrontPageComponent implements OnInit {

  public skillList = new Array<string>();
  mail =  localStorage.getItem('mail')
  constructor() { }

  ngOnInit(): void {
  }

}
