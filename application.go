// Basic Integer Obfuscation service using hashids
// (http://hashids.org/go/)

// Usage:
// Start server: go run application.go
// Client: curl 127.0.0.1:5000?id=2
// {"id":2,"obfuscated_id":"Ppe3M9"}

package main

import (
	"encoding/json"
	"github.com/speps/go-hashids"
	"math/rand"
	"net/http"
	"strconv"
)

type Response struct {
	Id           int    `json:"id"`
	ObfuscatedId string `json:"obfuscated_id"`
}

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
			r := Response{Id: id, ObfuscatedId: generatedId}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(r)
		}
	}
}

func main() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":5000", nil)
}
