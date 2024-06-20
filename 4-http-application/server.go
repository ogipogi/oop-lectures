// server.go
package main

import (
	"fmt"
	"net/http"
	"strings"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	switch r.Method {
	case http.MethodPost:
		w.WriteHeader(http.StatusAccepted)
	case http.MethodGet:
		fmt.Fprint(w, GetPlayerScore(player))
	}
}

func GetPlayerScore(name string) string {
	if name == "Pepper" {
		return "20"
	}
	if name == "Floyd" {
		return "10"
	}
	return ""
}

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(PlayerServer))
}
