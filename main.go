package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Caddy Backend Server")
	})
	fmt.Println("Serving Go app on port 8081")
	http.ListenAndServe(":8081", nil)
}
