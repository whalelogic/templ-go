package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/whalerapi/templ-go/templates"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", templ.Handler(templates.IndexPage()))
	mux.Handle("/about", templ.Handler(templates.AboutPage()))

	fmt.Println("Listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
