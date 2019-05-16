package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net"
	"net/http"
	"os"
	"time"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}

func RequestLogger(targetMux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		targetMux.ServeHTTP(w, r)

		// log request by who(IP address)
		requesterIP := r.Header.Get("X-Forwarded-For")

		log.WithFields(log.Fields{
			"Method":      r.Method,
			"RequestURI":  r.RequestURI,
			"RequesterIP": requesterIP,
			"Time":        time.Since(start),
		}).Info("")

	})
}

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
	log.Info("Starting web server...")

	mux := http.NewServeMux()
	mux.HandleFunc("/", welcome)
	mux.HandleFunc("/whoami", whoami)
	err := http.ListenAndServe(":8123", RequestLogger(mux))

	if err != nil {
		log.Fatalf("Failed to start server:%v", err)
	}
}
