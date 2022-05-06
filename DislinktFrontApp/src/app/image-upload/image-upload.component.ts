import { Component, OnInit } from '@angular/core';
import { FileUploadService } from '../services/file-upload-service';

@Component({
  selector: 'app-image-upload',
  templateUrl: './image-upload.component.html',
  styleUrls: ['./image-upload.component.css']
})
export class ImageUploadComponent implements OnInit {

  constructor(private fileUploadService: FileUploadService){}

  ngOnInit(): void {
  }

  file!: File;
  filename:string='';
  photoName:string='';
  fileinput:string='';

  onChange(imageInput:any) {
    this.file = imageInput.files[0];   
  }

  onUpload(){
  const reader = new FileReader();
    reader.readAsDataURL(this.file);
    

    this.photoName=this.filename.substring(this.filename.lastIndexOf("\\") + 1, this.filename.length); 
    
    reader.addEventListener('load', (event: any) => {      
      this.fileUploadService.upload(reader.result?.toString(), this.photoName).subscribe(
        (res) => {
          console.log(res);
        },
        (err) => {
          console.log(err);
        })
    });
}

}
