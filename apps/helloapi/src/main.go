// main.go
// This file is part of Epixel CRM Software

package main

import (
	"net/http"
	"log"

	"github.com/epixelnan/epicrm-module-demo/app"
)

func main() {
	hndlr := app.NewAppHandler()

	// TODO SECURITY use TLS or make access local-only
	// TODO use env for port no.
	err := http.ListenAndServe(":8080", hndlr)
	if err != nil {
		log.Fatal(err)
	}
}


