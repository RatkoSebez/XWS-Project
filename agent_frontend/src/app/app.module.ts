import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HomeComponent } from './components/home/home.component';
import { MenuComponent } from './components/menu/menu.component';
import { PageNotFoundComponent } from './components/page-not-found/page-not-found.component';
import { LoginService } from './services/login.service';
import { RegisterComponent } from './components/register/register.component';
import { AuthGuard } from './auth.guard';
import { LoginComponent } from './components/login/login.component';
import { RegisterCompanyComponent } from './components/register-company/register-company.component';
import { CompaniesComponent } from './components/companies/companies.component';
import { CompanyComponent } from './components/company/company.component';
import { EditCompanyModalComponent } from './components/edit-company-modal/edit-company-modal.component';
import { CompanyCardComponent } from './components/company-card/company-card.component';
import { CommentComponent } from './components/comment/comment.component';
import { SalaryComponent } from './components/salary/salary.component';

@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    MenuComponent,
    PageNotFoundComponent,
    RegisterComponent,
    LoginComponent,
    RegisterCompanyComponent,
    CompaniesComponent,
    CompanyComponent,
    EditCompanyModalComponent,
    CompanyCardComponent,
    CommentComponent,
    SalaryComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    HttpClientModule
  ],
  providers: [LoginService, AuthGuard],
  bootstrap: [AppComponent]
})
export class AppModule { }
