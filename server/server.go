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
	http.HandleFunc("/get_users", handlers.Middleware(ctx, handlers.GetUsersHandler))
	http.HandleFunc("/get_user_notifications", handlers.Middleware(ctx, handlers.GetUserNotificationsHandler))
	http.HandleFunc("/get_events", handlers.Middleware(ctx, nil))
	http.HandleFunc("/get_decisions", handlers.Middleware(ctx, nil))

	// Start Server
	fmt.Printf("Listening on port %s\n", PORT)
	http.ListenAndServe(":"+PORT, nil)
}
