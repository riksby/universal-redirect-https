package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func redirectHTTPS(w http.ResponseWriter, r *http.Request) {
	host, _, err := net.SplitHostPort(r.Host)
	if err != nil {
		host = r.Host
	}

	url := fmt.Sprintf("https://%s%s", host, r.RequestURI)
	http.Redirect(w, r, url, http.StatusMovedPermanently)
}

func main() {
	log.Fatal(http.ListenAndServe(":80", http.HandlerFunc(redirectHTTPS)))
}
