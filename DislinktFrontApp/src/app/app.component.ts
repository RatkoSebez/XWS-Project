import { Component } from '@angular/core';


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  constructor(){}

  logged:boolean=false;
  loading:boolean=false;
  

  logOut(){

  }

  getUsername(){

  }
  
}


class ImageSnippet {
  constructor(public src: string, public file: File) {}
}