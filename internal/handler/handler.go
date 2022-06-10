package handler

import (
	"fmt"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("endpoint hit: homepage")
	fmt.Fprintf(w, "Welcome to Homepage")
}
