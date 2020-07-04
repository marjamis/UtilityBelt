import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { catchError, map, tap } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class KubernetesService {
  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  constructor(private http: HttpClient) { }

  getData(): Observable<string> {
    return this.http.get<string>("http://localhost:8081/kubernetes", this.httpOptions).pipe(
      tap(_ => this.log("Message")),
      catchError(this.handleError<string>(`something happened`))
    );
  }

  handleError<T>(operation = 'operation', result?: T) {
    return (error: any): Observable<T> => {
      console.error(error);
      return of(result as T);
    }
  }

  private log(message: string) {
    console.log(`KubernetesService Issues`);
  }
}
