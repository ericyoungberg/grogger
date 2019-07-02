/**
* Sessions
*/

package main

import "bytes"


type Session struct {
    clients  []*Client
    requests int
    owner *Client
}

func NewSession() *Session {
    return &Session{
        clients: []*Client{},
        requests: 0,
    }
}


type SessionManager struct {
    sessions map[string]*Session
}

func (sm SessionManager) get(path string) *Session {
    session, found := sm.sessions[path]

    if !found {
        session = NewSession()
        sm.sessions[path] = session
    }
     
    return session
}

func NewSessionManager() *SessionManager {
    return &SessionManager{
        sessions: make(map[string]*Session),
    }
}
