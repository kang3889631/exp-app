import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// Define routes
const routes: Routes = [
    {
        path: 'myrule',
        loadChildren: () => import('./features/myrule/myrule.module').then(m => m.MyruleModule),
    },
    {
        path: '',
        redirectTo: '/myrule/view',
        pathMatch: 'full',
    },
];

@NgModule({
    imports: [
        RouterModule.forRoot(routes, { useHash: true }),
    ],
    exports: [
        RouterModule,
    ]
})
export class AppRoutingModule { }
