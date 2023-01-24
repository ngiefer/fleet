package main

import (
	"log"
	"net/http"
	"regexp"
)

// The directory to serve.
var (
	d          = http.Dir(".")
	fileserver = http.FileServer(d)
	tFile      = regexp.MustCompile("\\.crx$")
)

func myfileserver(w http.ResponseWriter, r *http.Request) {
	ruri := r.RequestURI
	if tFile.MatchString(ruri) {
		w.Header().Set("Content-Type", "application/x-chrome-extension")
	}
	fileserver.ServeHTTP(w, r)
}

func main() {
	http.HandleFunc("/", myfileserver)
	log.Fatal(http.ListenAndServe("localhost:1337", nil))
}
