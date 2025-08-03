package main

import (
	"fmt"
	"net/http"
	"server/cmd"
)

const PORT = 8080

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

func main() {
	cmd.SelectArm()

	http.HandleFunc("/", helloHandler)

	// Start Server
	http.ListenAndServe(":8080", nil)

	fmt.Printf("Listening on port %d\n", PORT)
}
