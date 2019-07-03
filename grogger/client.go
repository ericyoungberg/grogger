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
    browser   string

    connection *websocket.Conn
}

func (c Client) equal(_c *Client) bool {
    return _c.address == c.address && _c.browser == c.browser
}

func NewClient(r *http.Request) *Client {
    return &Client{
        address: r.RemoteAddr,
        agent: "",
        browser: r.Header.Get("Sec-WebSocket-Key"),
    }
}
