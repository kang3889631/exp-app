import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { SharedModule } from './shared/shared.module';
import { MyruleModule } from './features/myrule/myrule.module';
import { HttpClientModule } from '@angular/common/http';

import { FormsModule } from '@angular/forms';
import { AccountingModule } from './features/accounting/accounting.module';

@NgModule({
    declarations: [
        AppComponent,
    ],
    imports: [
        BrowserModule,
        FormsModule,
        AppRoutingModule,
        HttpClientModule,
        SharedModule,
        MyruleModule,
        AccountingModule,
    ],
    providers: [],
    bootstrap: [AppComponent],
})
export class AppModule { }
