import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';

import { AppComponent } from './app.component';
import { SidebarComponent } from './sidebar/sidebar.component';
import { DevelopersComponent } from './developers/developers.component';
import { AppsComponent } from './apps/apps.component';
import { AppinstsComponent } from './appinsts/appinsts.component';
import { CloudletsComponent } from './cloudlets/cloudlets.component';
import { ClustersComponent } from './clusters/clusters.component';
import { FlavorsComponent } from './flavors/flavors.component';
import { HttpClientModule } from '@angular/common/http'; 
import {CdkTableModule} from '@angular/cdk/table';
import { OperatorsComponent } from './operators/operators.component';

@NgModule({
  declarations: [
    AppComponent,
    SidebarComponent,
    DevelopersComponent,
    AppsComponent,
    AppinstsComponent,
    CloudletsComponent,
    ClustersComponent,
    FlavorsComponent,
    OperatorsComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    CdkTableModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
