package main

import (
	"fmt"
	"net/http"
)

var port = ":80"
var content = "Que bolón!💪"

func main() {
	http.HandleFunc("/", index)
	fmt.Println("listening on port", port)
	http.ListenAndServe(port, nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><body><h1>%s</h1></body></html>", content)
}
