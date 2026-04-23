package main

import (
	"log"
	"net/http"
)

// define a home hanler
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// if r.URL.path != "/"{

	// }
	w.Write([]byte("Hello from Snippetbox"))
}

func ShowSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet"))
}

func CreateSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet"))
}

func main() {

	//initialize a new server
	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/snippet", ShowSnippet)
	mux.HandleFunc("/snippet/create", CreateSnippet)

	log.Println("Starting the server now...")
	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}
