import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { catchError, map, tap } from 'rxjs/operators';

import { RedisItem } from './redisItem';

@Injectable({
  providedIn: 'root'
})
export class RedisService {
  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  constructor(private http: HttpClient) { }

  // getAllItemsLocal(): Observable<RedisItem[]> {
  //   return new Observable((observer) => {
  //     observer.next(this.data);
  //     return {
  //       unsubscribe() {
  //       }
  //     };
  //   });
  // }
  //
  // addLocal(key: string, value: string): void {
  //   this.data.push(<RedisItem> {
  //     key: key,
  //     value: value,
  //   })};
  //
  // deleteLocal(key: string): void {
  //   for( var i = 0; i < this.data.length; i++) {
  //     if (this.data[i].key === key) {
  //       this.data.splice(i, 1);
  //       i--;
  //     }}
  // }

  //TODO fix this so it return the RedisItem[] data rather than everything making the subscribers less specific and requiring the JSON knowledge.
  getAllItems(): Observable<RedisItem[]>{
    return this.http.get<RedisItem[]>("http://localhost:8081/redis?action=display", this.httpOptions).pipe(
      tap(_ => this.log("Message")),
      catchError(this.handleError<RedisItem[]>("something happened"))
    );
  }

  add(key: string, value: string): void {
    this.http.get(`http://localhost:8081/redis?action=add&key=${key}&value=${value}`).pipe(
      tap(_ => this.log("Message")),
      catchError(this.handleError("something happened"))
    ).subscribe();
  }

  delete(key: string): void {
    this.http.get(`http://localhost:8081/redis?action=del&key=${key}`).pipe(
      tap(_ => this.log("Message")),
      catchError(this.handleError("something happened"))
    ).subscribe();
  }

  handleError<T>(operation = 'operation', result?: T) {
    console.log("In the error")
    return (error: any): Observable<T> => {
      console.error(error);
      return of(result as T);
    }
  }

  private log(message: string) {
    console.log(`KubernetesService Issues`);
  }
}
