/// <reference lib="webworker" />

import {DoWork, runWorker} from 'observable-webworker';
import {Observable} from 'rxjs';
import {mergeMap} from 'rxjs/operators';
import {SourceSyncMessage} from '../models/queue/source-sync-message';
import {NewRepositiory} from '../../lib/database/pouchdb_repository';
import {NewClient} from '../../lib/conduit/factory';

export class SourceSyncWorker implements DoWork<string, string> {
  public work(input$: Observable<string>): Observable<string> {
    return input$.pipe(
      //mergeMap allows us to convert a promise into an observable
      // https://stackoverflow.com/questions/53649294/how-to-handle-for-promise-inside-a-piped-map
      mergeMap(msg => {
        try {
          console.log(msg); // outputs 'Hello from main thread'
          const sourceSyncMessage = JSON.parse(msg) as SourceSyncMessage

          const db = NewRepositiory(sourceSyncMessage.userIdentifier)
          const client = NewClient(sourceSyncMessage.source.source_type, sourceSyncMessage.source)
          console.log("!!!!!!!!!!!!!!STARTING WORKER SYNC!!!!!!!!!", sourceSyncMessage)
          return client.SyncAll(db)
            .then((resp) => {
              console.log("!!!!!!!!!!!!!COMPLETE WORKER SYNC!!!!!!!!!!", resp)
              return JSON.stringify(resp)

            })
            .catch((err) => {
              console.error("!!!!!!!!!!!!!ERROR WORKER SYNC!!!!!!!!!!", err)
              throw err
            })
          // return from(resp)

        } catch (e) {
          console.log("CAUGHT ERROR", e)
          console.trace(e)
          throw e
        }
      }),
    );
  }
}

runWorker(SourceSyncWorker);
