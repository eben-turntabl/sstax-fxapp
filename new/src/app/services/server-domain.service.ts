import { Injectable } from '@angular/core';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root',
})
export class ServerDomainService {
  public serverDomain: string;

  constructor() {
    this.serverDomain = environment.serverDomain;
  }

  public getDomain(): string {
    return this.serverDomain;
  }
}
