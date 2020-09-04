import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MyruleViewComponent } from './myrule-view/myrule-view.component';
import { MyruleUploadFormComponent } from './myrule-upload-form/myrule-upload-form.component';
import { Routes, RouterModule } from '@angular/router';
import { MyruleBasicFormComponent } from './myrule-basic-form/myrule-basic-form.component';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MyruleUserMsgAreaComponent } from './myrule-user-msg-area/myrule-user-msg-area.component';
import { MyruleDemoComponent } from './myrule-demo/myrule-demo.component';
import { MyruleDemoViewComponent } from './myrule-demo-view/myrule-demo-view.component';
import { MyruleBarComponent } from './myrule-bar/myrule-bar.component';

const routes: Routes = [
    {
        path: 'view',
        component: MyruleViewComponent,
    },
    {
        path: 'demo',
        component: MyruleDemoViewComponent,
    }
];


@NgModule({
    declarations: [
        MyruleViewComponent,
        MyruleUploadFormComponent,
        MyruleBasicFormComponent,
        MyruleUserMsgAreaComponent,
        MyruleDemoComponent,
        MyruleDemoViewComponent,
        MyruleBarComponent,
    ],
    imports: [
        CommonModule,
        FormsModule,
        ReactiveFormsModule,
        RouterModule.forChild(routes),
    ],
})
export class MyruleModule { }
