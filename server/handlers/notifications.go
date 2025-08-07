package handlers

import (
	"encoding/json"
	"fmt"
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

	// Debug
	fmt.Printf("Selected: %t | Decision ID: %s | Timestamp: %d\n", selected, decisionID, timestamp)

	success := events.CreateDecisionEvent(decisionID, selected, timestamp, db)

	if success {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// UpdateNotificationScoresHandler Update the notification scores in the database
func UpdateNotificationScoresHandler(ctx *cmd.AppContext, w http.ResponseWriter, r *http.Request) {
	// Check for post request
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Context
	db := ctx.Ctx.Value(cmd.DATABASE).(*db.DB)

	// Types
	// NotificationJson The temporary json structure of the data being filled out
	type NotificationJSON struct {
		ID          string  `json:"id"`
		Score       float64 `json:"score"`
		Probability float64 `json:"probability"`
	}

	// Payload The payload
	type Payload struct {
		UserID        string             `json:"user_id"`
		Notifications []NotificationJSON `json:"notifications"`
	}

	var payload Payload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Create the notifications
	userID := payload.UserID
	ns := []*notifications.Notification{}

	fmt.Printf("\n--- Notification Scores Updated ---\n")
	for i := 0; i < len(payload.Notifications); i++ {
		temp := payload.Notifications[i]

		fmt.Printf("\nNotification: id[%s] | Score: [%f]\n", temp.ID, temp.Score)

		n := notifications.CreateNotification(temp.ID, userID, temp.Score, 0, 0, temp.Probability, "", "")

		ns = append(ns, n)
	}
	fmt.Printf("\n--- Notification Scores Updated ---\n")

	success := notifications.UpdateNotificationScores(ns, db)

	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("success"))
	}
}
