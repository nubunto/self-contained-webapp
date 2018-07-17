package main

import (
	"fmt"
	"log"
	"net/http"

	// TODO: replace with absolute import path!
	_ "./statik"
	"github.com/rakyll/statik/fs"
)

func main() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", http.FileServer(statikFS))
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "oh hai")
	})
	http.ListenAndServe(":8080", nil)
}
