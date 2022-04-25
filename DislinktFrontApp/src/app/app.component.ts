import { Component } from '@angular/core';
import { FileUploadService } from './services/file-upload-service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  constructor(private fileUploadService: FileUploadService){}

  logged:boolean=false;
  loading:boolean=false;
  file!: File;
  filename:string='';
  photoName:string='';
  fileinput!:ArrayBuffer;

  logOut(){

  }

  getUsername(){

  }
  
  onChange(imageInput:any) {
    this.file = imageInput.files[0];   
  }

onUpload(){
  const reader = new FileReader();
    reader.readAsDataURL(this.file);
    this.fileinput!=reader.result;

    this.photoName=this.filename.substring(this.filename.lastIndexOf("\\") + 1, this.filename.length); 
    
    reader.addEventListener('load', (event: any) => {

      this.fileUploadService.upload(this.fileinput, this.photoName).subscribe(
        (res) => {
          console.log(res);
          console.log(reader.result)
        },
        (err) => {
          console.log(err);
          console.log(reader.result)
        })
    });
}

}


class ImageSnippet {
  constructor(public src: string, public file: File) {}
}