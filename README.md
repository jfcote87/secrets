# README #

secrets is a collection of simple routines initializing authorized services using
Username/Password, keys, json objects and other credentials from the system keyring.  

### Example ###

To store a Google Service Account in the keyring:

```
package main

import (
    "github.com/jfcote87/secrets"
    "github.com/jfcote87/secrets/googlejwt"
)

func main() {
    serviceAccountID := "cloud-function@project.iam.gserviceaccount.com" 
    // create key
    if err := secrets.SetFile(secrets.ServiceGoogle, 
        serviceAccountID, "keyfiles/project-740afd094af1.json"); err != nil {
            log.Fatalf("set failed: %v", err)
        }
    cl, err := googlejwt.Client(serviceAccountID, "jcote@example.com", "scope1","scope2)
    if err != nil {
        log.Fatalf("%v", err)
    }
    _ = cl
}
