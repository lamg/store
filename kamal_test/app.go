package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	fmt.Println("listening on port 80")
	http.ListenAndServe(":80", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><body><h1>Hello world!</h1></body></html>")
}
