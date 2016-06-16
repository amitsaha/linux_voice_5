// Basic HTTP server listening on 5000
package main

import (
	"fmt"
	"net/http"
	"runtime"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World. Running: ")
	fmt.Fprintf(w, runtime.Version())
}

func main() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":5000", nil)
}
