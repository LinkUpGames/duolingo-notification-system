package main

import (
	"fmt"
	"net/http"
	"os"
	"server/cmd"
	"server/handlers"
)

var PORT string = os.Getenv("SERVER_PORT")

func main() {
	// Setup Context
	ctx := cmd.SetupContext()

	// Handlers
	http.HandleFunc("/send_notification", handlers.Middleware(ctx, handlers.SendNotificationHandler))

	// Start Server
	http.ListenAndServe(":"+PORT, nil)

	fmt.Printf("Listening on port %s\n", PORT)
}
