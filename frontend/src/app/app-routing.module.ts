import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { KubernetesComponent } from './kubernetes/kubernetes.component'
import { RedisComponent } from './redis/redis.component'


const routes: Routes = [
  { path: '', redirectTo: '/', pathMatch: 'full' },
  { path: "kubernetes", component: KubernetesComponent },
  { path: "redis", component: RedisComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})

export class AppRoutingModule { }
