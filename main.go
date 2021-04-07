package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"personsal/go_health_check/health_check"
)

func main() {


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	health_check.Initialize()

	log.Fatal(http.ListenAndServe(":8081", nil))

}
