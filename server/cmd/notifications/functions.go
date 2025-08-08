// Package notifications is still in progress
package notifications

import (
	"encoding/json"
	"fmt"
	"os"
	"server/cmd"
	"server/db"
)

// SelectNotifcation Returns the id of the notifcation to send based on the current values
func SelectNotifcation(userID string, variables *cmd.Variables, db *db.DB) *Notification {
	// Fetch the notifications and the scores for this user
	notifications := getUserNotifications(userID, db, variables)

	// Calculte the recency delay
	computeNotificationDecay(notifications, variables.Penalty, variables.Factor, variables.CutOff)

	// Probabilities
	computeSoftmaxProb(notifications, float64(variables.Explore))

	// DEBUG: This is for debugging purposes only
	printNotifications(notifications)

	// Sample an arm using a weight probability
	notification := selectRandom(notifications)

	// Log the decision
	decisionID, err := addDecisionLog(db, notification.ID, notifications)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error with adding log: %s", err.Error())
	}
	notification.DecisionID = decisionID

	return notification
}

// MarshalNotification Send a notification to the user given their id by marshalling the struct and turning it into a json string
func MarshalNotification(notification *Notification) []byte {
	jsonBytes, err := json.Marshal(notification)
	if err != nil {
		return nil
	}

	return jsonBytes
}

func GetUserNotifications(userID string, db *db.DB, variables *cmd.Variables) []*Notification {
	notifications := getUserNotifications(userID, db, variables)

	return notifications
}
