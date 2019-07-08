/**
* Utils
*/

package main

import (
    "net/http"
    "strings"
)


func ParseUserAgent(r *http.Request) string {
    agent := r.UserAgent()

    var browser string
    if strings.Contains(agent, "Chrome") {
        browser = "Chrome"
    } else if strings.Contains(agent, "Firefox") {
        browser = "Firefox"
    } else {
        browser = "Unknown"
    }

    var platform string
    if strings.Contains(agent, "X11") {
        platform = "Desktop"
    } else {
        platform = "Mobile"
    }

    return browser + " " + platform
}
