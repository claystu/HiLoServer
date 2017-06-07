package main

import "fmt"
import "net/http"

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/route/", route)
	fmt.Println("Server listening on port http://localhost:5000/")
	http.ListenAndServe("localhost:5000", nil)
}
func route(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from route</h1>")
	fmt.Fprintf(w, "<h1>Path: %s</h1>", r.URL.Path)
	fmt.Fprintf(w, "<h1>Type: %T</h1>", r.URL.Path)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from index</h1>")
}
