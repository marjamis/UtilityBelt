import { Component, OnInit } from '@angular/core';

import { KubernetesService } from '../kubernetes.service';

@Component({
  selector: 'app-kubernetes',
  templateUrl: './kubernetes.component.html',
  styleUrls: ['./kubernetes.component.css']
})
export class KubernetesComponent implements OnInit {
  kubernetesData: string;

  constructor(private kubernetesService: KubernetesService) { }

  ngOnInit(): void {
    this.getData();
  }

  getData(): void {
    this.kubernetesService.getData()
        .subscribe(
          response => this.kubernetesData = JSON.stringify(response),
          err => this.kubernetesData = "Error",
          () => console.log("Kubernetes request sent and actioned.")
    );
  }
}
