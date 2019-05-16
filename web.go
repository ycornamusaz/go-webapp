package main

import (
	"fmt"
	"log"
	"net/http"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to my fist go server !")
}

func whoami(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Your IP adress is, %q", r.Header.Get("X-Forwarded-For"))
}

func main() {
	log.Println("Starting web server...")

	http.HandleFunc("/", welcome)
	http.HandleFunc("/whoami", whoami)
	err := http.ListenAndServe(":8123", nil)

	if err != nil {
		log.Fatalf("Failed to start server:%v", err)
	}
}
