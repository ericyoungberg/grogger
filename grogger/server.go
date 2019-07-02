/**
* Handler
*
* Handles the HTTP Request
*/

package main

import (
    "net/http"
)


type Server struct {
    http.Handler

    sessions *SessionManager
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    session := s.sessions.find(r.URL.Path)

    if r.Method == 'POST' {
        h.handlePOST(session, w, r);
    } else {
        h.handleGET(session, w, r);
    }
}

func (s Server) handleGET(s *Session, w http.ResponseWriter, r *http.Request) {
    
}

func (s Server) handlePOST(s *Session, w http.ResponseWriter, r *http.Request) {

}

func NewServer(sessions *SessionManager) *Server {
    return &Server{
        sessions: sessions,
    }
}
