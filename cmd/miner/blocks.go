package main

import (
	"fmt"
	"html"
	"net/http"
)

func blocks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
