package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var answer = 25
var guesses = 0

//MAXGUESSES sets maximum number of guesses allowed before you lose
const MAXGUESSES int = 5

func main() {
	http.HandleFunc("/hilo/", hilo)
	http.HandleFunc("/", index)
	fmt.Println("Server listening on port http://localhost:5000/")
	http.ListenAndServe("localhost:5000", nil)
}
func hilo(w http.ResponseWriter, r *http.Request) {
	guesses++
	pathGuess := r.URL.Path[len("/hilo/"):]

	// convert guess from string to int
	guess, _ := strconv.Atoi(pathGuess)

	if guess < answer && guesses < MAXGUESSES {
		fmt.Fprintf(w, "<h1>Too Low</h1>")
		fmt.Fprintf(w, "<h1>Guesses Taken: %d</h1>", guesses)
		fmt.Fprintf(w, "<h1>Guesses Left: %d</h1>", MAXGUESSES-guesses)
	} else if guess > answer && guesses < MAXGUESSES {
		fmt.Fprintf(w, "<h1>Too high</h1>")
		fmt.Fprintf(w, "<h1>Guesses Taken: %d</h1>", guesses)
		fmt.Fprintf(w, "<h1>Guesses Left: %d</h1>", MAXGUESSES-guesses)
	} else if guess == answer {
		fmt.Fprintf(w, "<h1>Correct!</h1>")
		fmt.Fprintf(w, "<h1>You win!!!!</h1>")
		guesses = 0
		http.Redirect(w, r, "/", http.StatusResetContent)
	} else {
		fmt.Fprintf(w, "<h1>You lose!!!</h1>")
		guesses = 0
		http.Redirect(w, r, "/", http.StatusResetContent)

	}

}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Guess a number between 1 and 100 using hilo/[number]</h1>")
}
