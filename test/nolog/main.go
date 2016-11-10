package main

import (
    "net/http"
    "text/template"
)

var (//content
    head = []byte(`<html><head><title>"oi"<title></head>`)
    body = []byte("<html>Hello man!</html>")
    foot = []byte(`</html>`)

)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write(head)
        template.HTMLEscape(w, body)
        w.Write(foot)
    })
    http.ListenAndServe(":8080", nil)
}