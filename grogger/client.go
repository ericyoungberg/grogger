/**
* Client
*/

package main

import (
    "net/http"

    "github.com/gorilla/websocket"
)


type Client struct {
    address   string
    agent     string
    id        string

    connection *websocket.Conn
}

func (c Client) equal(_c *Client) bool {
    return _c.address == c.address && _c.id == c.id
}

func NewClient(r *http.Request) *Client {
    browserId := r.Header.Get("Sec-WebSocket-Key")[:6]

    return &Client{
        address: r.RemoteAddr,
        agent: ParseUserAgent(r),
        id: browserId,
    }
}
