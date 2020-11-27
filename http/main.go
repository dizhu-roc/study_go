package main

import (
	"net/http"
)

//Handler 这是一个什么东西
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func main() {
	http.HandleFunc("/hello", Handler)

	http.ListenAndServe(":8080", nil)
}
