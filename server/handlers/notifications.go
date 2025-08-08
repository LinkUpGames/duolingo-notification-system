package handlers

import (
	"net/http"
	"server/cmd"
	"server/cmd/events"
	"server/cmd/notifications"
	"server/cmd/user"
	"server/db"
	"strconv"
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

// AcceptNotificationHandler Acceept the notification and update the database
func AcceptNotificationHandler(ctx *cmd.AppContext, w http.ResponseWriter, r *http.Request) {
	// Check and make sure that it's a POST request
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Context
	db := ctx.Ctx.Value(cmd.DATABASE).(*db.DB)

	// Parameters
	decisionID := r.URL.Query().Get("decision_id")
	selected, err := strconv.ParseBool(r.URL.Query().Get("selected"))
	if err != nil {
		selected = false
	}
	timestamp, err := strconv.ParseInt(r.URL.Query().Get("timestamp"), 10, 64)
	if err != nil {
		timestamp = -1
	}

	success := events.CreateDecisionEvent(decisionID, selected, timestamp, db)

	if success {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
