package main

import (
	"fmt"
	"net/http"
)

// https://codeql.github.com/codeql-query-help/go/go-reflected-xss/
func serve() {
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {

	})
	http.ListenAndServe(":80", nil)
}
