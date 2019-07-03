/**
* Handler
*
* Handles the HTTP Request
*/

package main

import (
    "fmt"
    "log"
    "net/http"
    "strings"

    "github.com/gobuffalo/packr"
    "github.com/gorilla/websocket"
)


var upgrader = websocket.Upgrader {
    ReadBufferSize: 1024,
    WriteBufferSize: 1024,

    // We don't care about origins
    CheckOrigin: func(_ *http.Request) bool { return true },
}


type Server struct {
    http.Handler

    sessions *SessionManager
    static packr.Box
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if r.Header.Get("Upgrade") == "websocket" {
        s.handleWS(w, r);
    } else {
        s.handleGET(w, r);
    }
}

func (s Server) handleGET(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/javascript")

    driver, err := s.static.FindString("grogger.js")
    if err != nil {
        w.WriteHeader(500) 
        w.Write([]byte("trouble loading 'grogger.js'"))
        return
    }

    driver = strings.Replace(driver, "{{SESSION}}", r.Host + r.URL.Path, 1)

    w.Write([]byte(driver))
}

func (s Server) handleWS(w http.ResponseWriter, r *http.Request) {
    session := s.sessions.getSession(r.URL.Path)
    client := session.getClient(r)

    if client == nil {
        w.WriteHeader(500)
        w.Write([]byte("couldn't connect client to session"))
        return
    }
        
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Panic(err) 
    }
    client.connection = conn

    for {
        msgType, msg, err := conn.ReadMessage()
        if err != nil {
            log.Println(err)
            return
        }

        sender := fmt.Sprintf("{\"sender\": \"%s\",", client.browser)
        msgWithSender := strings.Replace(string(msg), "{", sender, 1)
        session.log(msgWithSender)
        msg = []byte(msgWithSender)

        for _, peer := range session.peers(client) {
            if err := peer.connection.WriteMessage(msgType, msg); err != nil {
                log.Println(err)
            }
        }
    }
}

func NewServer(sessions *SessionManager, staticDir string) *Server {
    return &Server{
        sessions: sessions,
        static: packr.NewBox(staticDir),
    }
}
