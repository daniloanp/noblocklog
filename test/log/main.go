package main

import (
    "net/http"
    "text/template"
    "log"
)

var (//content
    head = []byte(`<html><head><title>"oi"</title></head>`)
    body = []byte("<html>Hello man!</html>")
    foot = []byte(`</html>`)
)

func main() {
    log.SetFlags(0)
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        log.Println("Começa...")
        w.Write(head)
        log.Println("vai corpo...")
        template.HTMLEscape(w, body)
        log.Println("Foi corpo...")
        w.Write(foot)
    })
    http.ListenAndServe(":8080", nil)
}