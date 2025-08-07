package handlers

import (
	"net/http"
	"server/cmd"
	"server/cmd/notifications"
	"server/cmd/user"
	"server/db"
)

// SendNotificationHandler Select an arm based on the user
func SendNotificationHandler(ctx *cmd.AppContext, w http.ResponseWriter, r *http.Request) {
	// Get Parameters
	userID := r.URL.Query().Get("user_id")
	userName := r.URL.Query().Get("name")

	// Context
	variables := ctx.Ctx.Value(cmd.VARIABLES).(*cmd.Variables)
	db := ctx.Ctx.Value(cmd.DATABASE).(*db.DB)

	// Create user if it does not exist
	user.CheckUser(db, userID, userName)

	// Select a notification
	notification := notifications.SelectNotifcation(userID, variables, db)
	marhal := notifications.MarshalNotification(notification)

	if marhal == nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusExpectationFailed)

		w.Write([]byte("Error"))
	} else {

		// Set Response Headers
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write(marhal)
	}
}
