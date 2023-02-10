package main

import (
    "fmt"
    "net/http"
)

func Hello(name string) string {
    return "hello, " + name
}

func main() {
    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "hello hello hello")
    })

    fmt.Println("start")
    http.ListenAndServe(":8080", nil)
}