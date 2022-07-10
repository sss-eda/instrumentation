package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// The path routed here will be stations. So take the paths from there.
func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("this is one"))
	}).Methods("GET")

	http.ListenAndServe(":8080", r)
}
