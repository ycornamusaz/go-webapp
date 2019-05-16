package main

import (
    "fmt"
    "net/http"
)

func welcome (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to my fist go server !")
}

func main() {
    fmt.Println("Strating web server...")
    http.HandleFunc("/", welcome)
    http.ListenAndServe(":8123", nil)
}
