import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { FrontPageComponent } from './front-page/front-page.component';
import { LoginPageComponent } from './login-page/login-page.component';
import { RegistrationPageComponent } from './registration-page/registration-page.component';

const routes: Routes = [
  {path:'', component: FrontPageComponent},
  {path:'register', component: RegistrationPageComponent},
  {path:'login', component: LoginPageComponent},
  //{path:'edit', component: EditProfilePageComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
