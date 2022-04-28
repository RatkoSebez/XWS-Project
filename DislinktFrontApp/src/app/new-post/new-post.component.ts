import { Component, OnInit } from '@angular/core';
import { NewPostDTO } from '../dtos/NewPostDTO';
import { Post } from '../model/post';
import { FileUploadService } from '../services/file-upload-service';
import { PostService } from '../services/post-service';

@Component({
  selector: 'app-new-post',
  templateUrl: './new-post.component.html',
  styleUrls: ['./new-post.component.css']
})
export class NewPostComponent implements OnInit {

  constructor(private fileUploadService: FileUploadService, private postService: PostService){}

  ngOnInit(): void {
  }

  file!: File;
  filename:string='';
  photoName:string='';
  fileinput:string='';
  newPost:NewPostDTO={
    links:[],
    photos:[],
    postText:''
  }
  responsePost:Post=new Post();
  showPost:boolean=false;

  onChange(imageInput:any) {
    this.file = imageInput.files[0];   
  }

  onUpload(){
  const reader = new FileReader();
    reader.readAsDataURL(this.file);
    

    this.photoName=this.filename.substring(this.filename.lastIndexOf("\\") + 1, this.filename.length); 
    this.newPost.photos.push(this.photoName)
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

makeNewPost(){

  this.postService.makeNewPost(this.newPost).subscribe( (res) => {
    console.log(res);
    this.responsePost=res;
    this.showPost=true;
  },
  (err) => {
    console.log(err);
  })

}

}
