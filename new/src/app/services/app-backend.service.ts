import { Injectable } from '@angular/core';
import { catchError, Observable, tap } from 'rxjs';
import {
  HttpClient,
  HttpErrorResponse,
  HttpHeaders,
  HttpParams,
} from '@angular/common/http';
import { ServerDomainService } from './server-domain.service';

@Injectable({
  providedIn: 'root'
})
export class AppBackendService {
  serverDomain: string;


  constructor(
    private http: HttpClient,
    private httpDomain: ServerDomainService

  ) { this.serverDomain = this.httpDomain.getDomain(); }



  submitSignUp(form: any): Observable<any> {
    return this.http
      .post(`${this.serverDomain}/api/auth/register`, form)
      .pipe(
        tap((data) => {}),
        catchError(this.handleError)
      );
  }


  public submitSignIn(form: any): Observable<any> {
    return this.http
      .post(`${this.serverDomain}/api/auth/login`, form)
      .pipe(
        tap((data) => {}),
        catchError(this.handleError)
      );
  }
}
