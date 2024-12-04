package main

import (
	"github.com/spacefall/stalewall-proxy/src"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		src.Handler(w, r)
	})
	log.Print("Serving on http://localhost:3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
