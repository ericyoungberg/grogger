/**
* Grogger
* 
* A `console.log` multiplexer
*/

package main

import (
    "flag"
    "log"
    "net/http"
    "strconv"
)


func main() {
    intport := flag.Int("port", 2358, "port to listen for requests")
    flag.Parse()

    port := strconv.Itoa(*intport)
    
    sessionManager := newSessionManager()

    err := http.ListenAndServe(":" + port, newHandler(sessionManager))
    if err != nil {
        log.Fatal(err)
    }
}
