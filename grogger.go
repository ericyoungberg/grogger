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
    
    sessionManager := NewSessionManager()

    log.Println("Grogger has started on :" + port + " ...")
    err := http.ListenAndServe(":" + port, NewServer(sessionManager, "./assets"))
    if err != nil {
        log.Fatal(err)
    }
}
