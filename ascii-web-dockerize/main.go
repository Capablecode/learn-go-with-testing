package main

import (
	asciiartweb "ascii-web-dockerize/handler"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", asciiartweb.ServeHome)
	mux.HandleFunc("/ascii-art", asciiartweb.PrintAscii)
	fmt.Println("Server starting at :8080...")
	http.ListenAndServe(":8080", mux)
}
