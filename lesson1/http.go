package main

import (
	"net/http"
)

// Handlers 这是注释
func Handlers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!!!"))
}

func main() {
	http.HandleFunc("/hello", Handlers)

	http.ListenAndServe(":8080", nil)
}
