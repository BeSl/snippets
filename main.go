package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Start server")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("hello snip"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Form show snippet"))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)

		http.Error(w, "метод запрещен", 405)
		return
	}

	w.Write([]byte("Form create snippet"))
}
