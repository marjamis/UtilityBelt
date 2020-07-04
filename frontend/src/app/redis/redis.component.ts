import { Component, OnInit } from '@angular/core';

import { RedisService } from '../redis.service';
import { RedisItem } from '../redisItem';

@Component({
  selector: 'app-redis',
  templateUrl: './redis.component.html',
  styleUrls: ['./redis.component.css']
})
export class RedisComponent implements OnInit {
  redisItems: RedisItem[];
  newdatakey: string;
  newdatavalue: string;

  constructor(private redisService: RedisService) { }

  ngOnInit(): void {
    this.getAllItems();
  }

  getAllItems(): void {
    // TODO investigate the below to have less logic here
    // this.redisService.getAllItems().subscribe(x => this.redisItems = x);
    this.redisService.getAllItems().subscribe((x: any) => {
        this.redisItems = x.RedisItems;
      });
  }


  add(key: string, value: string): void {
    if (key != "" && value != "") {
      this.redisService.add(key, value);
    }
    this.getAllItems();
  }

  delete(key: string): void {
    this.redisService.delete(key);
    this.getAllItems();
  }
}
