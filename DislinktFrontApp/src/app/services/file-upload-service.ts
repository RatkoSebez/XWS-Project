import { Injectable } from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Observable} from 'rxjs';
import { LoginService } from './login-service';
import { HttpHeaders } from '@angular/common/http';
@Injectable({
  providedIn: 'root'
})
export class FileUploadService {
    
  public header : HttpHeaders | undefined
  // API url
  baseApiUrl = "http://localhost:8082/save-photo"
    
  constructor(private http:HttpClient, private ls : LoginService) { }
  
  // Returns an observable
  upload(file: string | undefined, filename:string):Observable<any> {

    const objec={
      fileName:filename,
      fileData:file
    }
      this.header = this.ls.getHeaders()
      return this.http.post(this.baseApiUrl, objec, {headers:this.header})
  }
}