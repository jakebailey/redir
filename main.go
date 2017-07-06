package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/alexflint/go-arg"
)

var args = struct {
	Port   uint   `arg:"env"`
	URL    string `arg:"required,env:REDIR_URL,help:url to redirect all requests to"`
	Scheme string `arg:"env:REDIR_SCHEME,help:scheme to use if not specified in url"`
}{
	Port:   5000,
	Scheme: "http",
}

func main() {
	p := arg.MustParse(&args)

	u, err := url.Parse(args.URL)
	if err != nil {
		p.Fail("error parsing URL: " + err.Error())
	}

	if u.Scheme == "" {
		u.Scheme = args.URL
	}

	log.Println("redirecting all traffic to", u)

	addr := fmt.Sprintf(":%v", args.Port)
	log.Println("starting server at", addr)

	handler := http.RedirectHandler(u.String(), http.StatusSeeOther)
	log.Fatal(http.ListenAndServe(addr, handler))
}
