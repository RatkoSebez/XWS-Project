import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { FrontPageComponent } from './front-page/front-page.component';
import { LoginPageComponent } from './login-page/login-page.component';
import { ProfileEditPageComponent } from './profile-edit-page/profile-edit-page.component';
import { ProfilePageComponent } from './profile-page/profile-page.component';
import { ProfilePostsComponent } from './profile-posts/profile-posts.component';
import { RegistrationPageComponent } from './registration-page/registration-page.component';

var mail = ''
const routes: Routes = [
  {path:'', component: FrontPageComponent},
  {path:'register', component: RegistrationPageComponent},
  {path:'login', component: LoginPageComponent},
  {path:'profile', component: ProfilePageComponent},
  {path: 'edit', component: ProfileEditPageComponent},

  {path: ':mail', component: ProfilePostsComponent},

  //{path:'edit', component: EditProfilePageComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
