package main

import (
    
    "net/http"
    "io"
)

func index(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "hello Docker file changed")
}

func main() {
    http.HandleFunc("/", index)
    http.ListenAndServe(":80", nil)
}
