import { Component, OnInit, ViewChild } from '@angular/core';
import { MyruleService } from '../service/myrule.service';
import { MyruleUserMsgAreaComponent } from '../myrule-user-msg-area/myrule-user-msg-area.component';

@Component({
    selector: 'app-myrule-view',
    templateUrl: './myrule-view.component.html',
    styleUrls: ['./myrule-view.component.scss'],
})
export class MyruleViewComponent implements OnInit {
    // Set up a direct access to the child by using template reference
    @ViewChild('myruleUserMsgArea')
    myruleUserMsgAreaComponent: MyruleUserMsgAreaComponent;

    //
    constructor(
        private myruleService: MyruleService,
    ) {
    }

    ngOnInit(): void {
        this.initListener();
    }

    initListener(): void {
        // Subscribe
        this.myruleService.basicFormFirstNameChanged$.subscribe(
            (next: string) => {
                console.log('myrule-view - next', next);

                this.myruleUserMsgAreaComponent.append(next);
            },
            //error => {
            //},
        );
    }

}
