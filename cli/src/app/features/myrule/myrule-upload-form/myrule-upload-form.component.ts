import { Component, OnInit } from '@angular/core';
import { MyruleService } from '../service/myrule.service';

@Component({
    selector: 'app-myrule-upload-form',
    templateUrl: './myrule-upload-form.component.html',
    styleUrls: ['./myrule-upload-form.component.scss'],
})
export class MyruleUploadFormComponent implements OnInit {
    // Global Variable
    private fileToUpload: File = null;

    constructor(
        private myruleService: MyruleService,
    ) { }

    ngOnInit() {
    }

    // handleFileInput
    handleFileInput(files: FileList) {
        //debugger;

        const fileToUpload: File = files.item(0);
        //this.fileToUpload = files.item(0);

        // Call postFileUpload service
        this.myruleService.postFileUpload(fileToUpload).subscribe(
            next => {
                console.log('here', next);
            },
        );
    }

    // onUploadBtnClicked
    onUploadBtnClicked($event: any): void {
        console.log($event);
    }
}
