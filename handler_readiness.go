package main

import "net/http"

// Use this function to define HTTP Handler in GO standard.
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	// For learning purposes, an empty payload which is an empty struct is passed.
	respondWithJSON(w, 200, struct{}{})
}
