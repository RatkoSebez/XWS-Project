import { Component } from '@angular/core';


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  constructor(){}

  public logged:boolean=true;
  public notlogged:boolean=true;
  loading:boolean=false;
  public  logVisible = true;
  public  outvisible = false;
  static lv = true;
  static ov = false;

  logOut(){

  }


  getUsername(){

  }
  
}


class ImageSnippet {
  constructor(public src: string, public file: File) {}
}