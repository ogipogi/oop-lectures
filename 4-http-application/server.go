// server.go
package main

import (
	"fmt"
	"net/http"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	if player == "Pepper" {
		fmt.Fprint(w, "20")
	} else if player == "Floyd" {
		fmt.Fprint(w, "10")
	}
}

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(PlayerServer))
}
