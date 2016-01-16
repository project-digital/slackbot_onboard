// Package slackbot_onboard (Placeholder for future documentation)/
// Via this format godoc will auto-generate and display the documenation online.
package main 

import (
    "fmt"
    "net/http"
)

func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, world! We making moves :O")
}
