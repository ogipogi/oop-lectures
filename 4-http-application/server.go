// server.go
package main

import (
	"fmt"
	"net/http"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20")
}

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(PlayerServer))
}
