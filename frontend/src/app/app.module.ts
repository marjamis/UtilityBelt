import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpClientModule }    from '@angular/common/http';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { KubernetesComponent } from './kubernetes/kubernetes.component';
import { RedisComponent } from './redis/redis.component';

@NgModule({
  declarations: [
    AppComponent,
    KubernetesComponent,
    RedisComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
