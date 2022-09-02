import { Injectable } from '@angular/core';
import { nameByRace } from 'fantasy-name-generator';
import { BehaviorSubject, Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class UserInformationService {
  private nameSubject = new BehaviorSubject<string>('');
  private isSpectatorSubject = new BehaviorSubject<boolean>(false);

  constructor() {
    // If new user or loading from localstorage failed, use a generated name
    let name = this.loadFromLocalStorage('name') || nameByRace('human');
    let isSpectator = this.loadFromLocalStorage('isSpectator') || false;
    this.nameSubject.next(name);
    this.isSpectatorSubject.next(isSpectator)
  }

  setName(name: string): void {
    this.nameSubject.next(name);
    this.saveToLocalStorage('name', name);
  }

  setSpectator(isSpectator: boolean): void {
    this.isSpectatorSubject.next(isSpectator);
    this.saveToLocalStorage('isSpectator', isSpectator);
  }

  nameObservable(): Observable<string> {
    return this.nameSubject.asObservable();
  }

  getName(): string {
    return this.nameSubject.getValue();
  }

  spectatorObservable(): Observable<boolean> {
    return this.isSpectatorSubject.asObservable();
  }

  isSpectator(): boolean {
    return this.isSpectatorSubject.getValue();
  }

  private loadFromLocalStorage(key: string): any {
    try {
      return localStorage.getItem(key);
    } catch (e) {
      console.debug('Localstorage is either unavailable or disabled');
      return null;
    }
  }

  private saveToLocalStorage(key: string, value: any): void {
    try {
      localStorage.setItem(key, value);
    } catch (e) {
      console.debug('Localstorage is either unavailable or disabled');
    }
  }
}
