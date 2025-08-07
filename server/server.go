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
	http.HandleFunc("/get_user_decisions", handlers.Middleware(ctx, handlers.GetUserDecisionsHandler))
	http.HandleFunc("/get_decision_probabilities", handlers.Middleware(ctx, handlers.GetDecisionProbabilitiesHandler))
	http.HandleFunc("/get_decision_event", handlers.Middleware(ctx, handlers.GetDecisionEventHandler))
	http.HandleFunc("/accept_notification", handlers.Middleware(ctx, handlers.AcceptNotificationHandler))
	http.HandleFunc("/update_notification_scores", handlers.Middleware(ctx, handlers.UpdateNotificationScoresHandler))

	// Start Server
	fmt.Printf("Listening on port %s\n", PORT)
	http.ListenAndServe(":"+PORT, nil)
}
