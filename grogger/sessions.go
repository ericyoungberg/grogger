/**
* Sessions
*/

package main


type session struct {
    requests int
}

type sessionManager struct {
    sessions map[string]*session
}

func (s sessionManager) get(path string) *session {
    sesh, found := s.sessions[path]

    if !found {
        sesh = &session{0}
        s.sessions[path] = sesh
    }
     
    return sesh
}

func newSessionManager() *sessionManager {
    return &sessionManager{
        sessions: make(map[string]*session),
    }
}
