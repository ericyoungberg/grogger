/**
* Client
*/

package main

import (
    "io/ioutil"
    "fmt"
    "net/http"
)


type Client struct {
    address   string
    agent     string
    browser   string
}

func (c Client) Equal(_c *Client) bool {
    return _c.address == c.address && _c.browser == c.browser
}

func NewClient(r *http.Request) *Client {
    var browser string

    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
    }

    fmt.Printf("%v\n", body)

    return &Client{
        address: r.RemoteAddr,
        agent: "",
        browser: browser,
    }
}
