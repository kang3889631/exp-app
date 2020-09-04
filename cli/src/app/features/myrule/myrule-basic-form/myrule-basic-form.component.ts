import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl } from '@angular/forms';
import { MyruleService } from '../service/myrule.service';

@Component({
    selector: 'app-myrule-basic-form',
    templateUrl: './myrule-basic-form.component.html',
    styleUrls: ['./myrule-basic-form.component.scss']
})
export class MyruleBasicFormComponent implements OnInit {
    public profileForm: FormGroup = null;

    constructor(
        private myruleService: MyruleService,
    ) { }

    ngOnInit(): void {
        this.profileForm = new FormGroup({
            firstName: new FormControl(''),
            lastName: new FormControl(''),
        });
    }

    // onSaveBtnClicked
    onSaveBtnClicked($event: MouseEvent): void {
        console.log('step 1: onSaveBtnClicked');

        const data = this.profileForm.value;

        // NOTE: Observable is lazy. It will not does work until at least one
        // person (code) is subscribed
        const postBasic$ = this.myruleService.postBasic$(data);

        postBasic$.subscribe(
            next => {
                // Once the radio music is received, we do something, we can
                // either listen or sing
                console.log('step 3: postBasic$ - response ', next);
            },
        );
    }

    // onFirstNameChanged ...
    onFirstNameChanged($event: KeyboardEvent): void {
        console.log('myrule-basic-from - onFirstNameChanged', $event);

        const inputStr = ($event.target as HTMLInputElement).value;

        console.log('myrule-basic-from - inputStr', inputStr);

        // Raise the basicFormFirstNameChanged$ event
        // Emit a value to the observable variable.
        setTimeout(() => {
            this.myruleService.basicFormFirstNameChanged$.next(
                inputStr,
            );
        }, 0);
    }
}
