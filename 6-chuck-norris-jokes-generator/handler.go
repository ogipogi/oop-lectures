package main

import (
	"encoding/json"
	"html/template"
	"net/http"
)

func jokeHandler(w http.ResponseWriter, r *http.Request) {
	// Juck Norris API call
	resp, err := http.Get("https://api.chucknorris.io/jokes/random")
	if err != nil {
		http.Error(w, "Failed to fetch https://api.chucknorris.io/jokes/random", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var randomJoke Joke
	if err := json.NewDecoder(resp.Body).Decode(&randomJoke); err != nil {
		http.Error(w, "Failed to decode joke", http.StatusInternalServerError)
		return
	}

	// render a html page
	tmpl := template.Must(template.ParseFiles("template.html"))
	tmpl.Execute(w, randomJoke)

}
