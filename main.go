package main

import (
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/namsral/flag"
)

func main() {
	fs := flag.NewFlagSetWithEnvPrefix(os.Args[0], "REDIR", flag.ExitOnError)

	ur := fs.String("url", "", "url to redirect all requests to")
	scheme := fs.String("scheme", "http", "scheme to use if not specified in url")
	addr := fs.String("addr", ":5000", "address to run the server on")

	fs.Parse(os.Args[1:])

	if *ur == "" {
		log.Fatal("url must be provided")
	}

	u, err := url.Parse(*ur)

	if err != nil {
		log.Fatal(err)
	}

	if u.Scheme == "" {
		u.Scheme = *scheme
	}

	log.Printf("redirecting all traffic to %v", u)

	log.Fatal(http.ListenAndServe(*addr, http.RedirectHandler(u.String(), http.StatusSeeOther)))
}
