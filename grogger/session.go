/**
* Sessions
*/

package main

import (
    "log"
    "net/http"
)


type Session struct {
    clients  []*Client
    path string
}

func (s *Session) getClient(r *http.Request) *Client {
    client := NewClient(r)

    // First, try to return an existing client
    for _, _client := range s.clients {
        if _client.equal(client) {
            return _client
        }
    }

    s.log(client.browser + " connected")

    // Create a new client since the client is new to this session
    s.clients = append(s.clients, client)

    return client
}

func (s *Session) log(message string) {
    log.Printf("#%s: %s", s.path, message)
}

func (s *Session) peers(client *Client) []*Client {
    peers := []*Client{}

    for _, peer := range s.clients {
        if !client.equal(peer) {
            peers = append(peers, peer) 
        }
    }

    return peers
}

func NewSession(path string) *Session {
    return &Session{
        clients: []*Client{},
        path: path,
    }
}


type SessionManager struct {
    sessions map[string]*Session
}

func (sm SessionManager) getSession(path string) *Session {
    session, found := sm.sessions[path]

    if !found {
        session = NewSession(path)
        sm.sessions[path] = session
    }
     
    return session
}

func NewSessionManager() *SessionManager {
    return &SessionManager{
        sessions: make(map[string]*Session),
    }
}
