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

	// Get the max score
	var maxScore float64 = 0
	for _, notification := range notifications {
		score := notification.Score

		if score > maxScore {
			maxScore = score
		}
	}

	// Compute Exponentials with Recovering differnce
	total := computeExpScores(notifications, float64(variables.TEMPERATURE), maxScore)

	// Normalize to get the probabilities
	computeProbabilities(notifications, total)
	printNotifications(notifications)

	// Sample an arm using a weight probability
	notification := selectRandom(notifications)

	// Log the decision
	err := addDecisionLog(db, notification.ID, userID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error with adding log: %s", err.Error())
	}

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
