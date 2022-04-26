import { Injectable } from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Observable} from 'rxjs';
@Injectable({
  providedIn: 'root'
})
export class FileUploadService {
    
  // API url
  baseApiUrl = "http://localhost:8082/save-photo"
    
  constructor(private http:HttpClient) { }
  
  // Returns an observable
  upload(file: string | undefined, filename:string):Observable<any> {

    const objec={
      fileName:filename,
      fileData:file
    }

      return this.http.post(this.baseApiUrl, objec)
  }
}