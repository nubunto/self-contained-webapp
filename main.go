package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/nubunto/self-contained-webapp/statik"
	"github.com/rakyll/statik/fs"
)

//go:generate statik -src=web/build
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
