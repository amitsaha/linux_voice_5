// Basic Integer Obfuscation service using hashids
// (http://hashids.org/go/)

// Usage:
// Start server: go run application.go
// Client: curl 127.0.0.1:5000?id=2
// Vj7ygQ

package main

import (
	"fmt"
	"github.com/speps/go-hashids"
	"math/rand"
	"net/http"
	"strconv"
)

func getHashId(id int) (string, error) {
	hd := hashids.NewData()
	hd.Salt = strconv.Itoa(rand.Int())
	hd.MinLength = 6
	h := hashids.NewWithData(hd)
	return h.Encode([]int{id})
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Bad id supplied", 400)
	} else {
		generatedId, err := getHashId(id)
		if err != nil {
			http.Error(w, "Error generating Id", 500)
		} else {
			fmt.Fprintf(w, generatedId)
		}
	}
}

func main() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":5000", nil)
}
