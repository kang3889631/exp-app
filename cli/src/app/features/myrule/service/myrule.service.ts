import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import { tap } from 'rxjs/operators';

@Injectable({
    providedIn: 'root',
})
export class MyruleService {

    // For observable variable, the naming convetion is to end with a $ dollar sign.
    public basicFormFirstNameChanged$: Subject<string> = new Subject<string>();

    constructor(
        private http: HttpClient,
    ) {
    }

    // postFileUpload
    postFileUpload(fileToUpload: File): Observable<any> {
        const endpoint = '/api/Upload';
        const formData: FormData = new FormData();
        console.log('debug', fileToUpload);

        formData.append('fileKey', fileToUpload, fileToUpload.name);

        const httpOptions = {
            headers: new HttpHeaders({
                //'Content-Type': 'application/json',
            }),
            params: {
                clientFilename: fileToUpload.name,
            },
        };

        return this.http.post<any>(
            endpoint,
            formData,
            httpOptions,
        ).pipe(

        );
    }

    // postBasic
    // Some events will happen in the future that the code needs to take action
    // on
    postBasic$(data: any): Observable<any> {
        const endpoint = '/basic';

        return this.http.post<any>(
            endpoint,
            data,
        ).pipe(
            tap(
                next => {
                    console.log('step 2: tap - before emit value (return data back)');
                },
            ),
        );
    }
}
