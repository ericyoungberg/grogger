/**
* Handler
*
* Handles the HTTP Request
*/

package main

import (
    "fmt"
    "net/http"
)


type handler struct {
    http.Handler

    sessions *sessionManager
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    session := h.sessions.get(r.URL.Path)

    session.requests += 1

    fmt.Println(session.requests)

    /*
    if r.Method == 'POST' {
        h.handlePOST(w, r);
    } else {
        h.handleGET(w, r);
    }
    */
}

func newHandler(sessions *sessionManager) *handler {
    return &handler{
        sessions: sessions,
    }
}

/*
func (h handler) handleGET(w, http.ResponseWriter, r *http.Request) {

}

func (h handler) handlePOST(w http.ResponseWriter, r *http.Request) {

}
*/
