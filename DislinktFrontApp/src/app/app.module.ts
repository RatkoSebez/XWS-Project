import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { FormsModule } from '@angular/forms';
import { FrontPageComponent } from './front-page/front-page.component';
import { LoginPageComponent } from './login-page/login-page.component';
import { RegistrationPageComponent } from './registration-page/registration-page.component';
//<<<<<<< HEAD
import { RouterModule } from '@angular/router';
import {ToastrModule} from 'ngx-toastr';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { ProfileEditPageComponent } from './profile-edit-page/profile-edit-page.component';


import { ImageUploadComponent } from './image-upload/image-upload.component';
import { NewPostComponent } from './new-post/new-post.component';
import { PostBoxComponent } from './post-box/post-box.component';
//>>>>>>> d5b454f8c581b57932eed820dde1d1fa63716a55

@NgModule({
  declarations: [
    AppComponent,
    FrontPageComponent,
    LoginPageComponent,
    RegistrationPageComponent,
//<<<<<<< HEAD
    ProfileEditPageComponent,
          
//=======
    ImageUploadComponent,
    NewPostComponent,
    PostBoxComponent
//>>>>>>> d5b454f8c581b57932eed820dde1d1fa63716a55
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    HttpClientModule, 
    RouterModule,
    ToastrModule.forRoot(),
    BrowserAnimationsModule

  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
