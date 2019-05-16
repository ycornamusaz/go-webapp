package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func getIP() string {
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			return ipv4.String()
		}
	}
	return ""
}

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to my fist go server ! :)")
	name, _ := os.Hostname()
	fmt.Fprintln(w, "You are on", getIP())
	fmt.Fprintln(w, "You are on host", name)
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
