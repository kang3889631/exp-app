import { Component, OnInit } from '@angular/core';

@Component({
    selector: 'app-myrule-user-msg-area',
    templateUrl: './myrule-user-msg-area.component.html',
    styleUrls: ['./myrule-user-msg-area.component.scss']
})
export class MyruleUserMsgAreaComponent implements OnInit {
    // Events

    // Variables
    public userInputMsg: string;

    //
    constructor() { }

    ngOnInit(): void {
    }

    append(str: string): void {
        console.log('myrule-user-msg-area - append', str);

        this.userInputMsg = str;
    }
}
