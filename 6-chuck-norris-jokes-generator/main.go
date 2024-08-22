package main

import "net/http"

func main() {
	http.HandleFunc("/joke", jokeHandler)
	http.ListenAndServe(":8081", nil)
}
