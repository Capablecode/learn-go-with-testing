package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not Found", 404)
		return
	}

	w.Write([]byte("Hello From Snippet..."))
}
func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		// http.Error(w, "404 Not Found", 404)
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/createsnippet", createSnippet)
	mux.HandleFunc("/snippet", showSnippet)
	log.Println("Starting Server on 8000...")
	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}
