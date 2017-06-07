package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var answer = 25
var guesses = 0

func main() {
	http.HandleFunc("/hilo/", hilo)
	http.HandleFunc("/", index)
	fmt.Println("Server listening on port http://localhost:5000/")
	http.ListenAndServe("localhost:5000", nil)
}
func hilo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Guess a number between 1 and 100 using hilo/[number]</h1>")
	fmt.Fprintf(w, "<h1>Path: %s</h1>", r.URL.Path)
	guess := r.URL.Path[len("/hilo/"):]
	fmt.Fprintf(w, "<h1>Type: %T = Value: %v</h1>", guess, guess)

	// convert guess from string to int
	s, err := strconv.Atoi(guess)
	if err != nil {
		fmt.Fprint(w, "<h1>Error detected in String conversion of guess</h1>")
	}
	if s < answer {
		fmt.Fprintf(w, "<h1>Too Low</h1>")
	} else if s > answer {
		fmt.Fprintf(w, "<h1>Too high</h1>")
	} else {
		fmt.Fprintf(w, "<h1>Correct!</h1><h1>You win!!!</h1>")
	}

}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Guess a number between 1 and 100 using hilo/[number]</h1>")
}
