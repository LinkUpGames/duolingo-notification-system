// Package notifications is still in progress
package notifications

import (
	"fmt"
	"math/rand"
	"os"
	"server/cmd"
	"server/cmd/user"
	"server/db"
)

// SelectNotifcation Returns the id of the notifcation to send based on the current values
func SelectNotifcation(userID string, variables *cmd.Variables, db *db.DB) string {
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

	// Sample an arm using a weight probability
	notificationID := ""
	r := rand.Float64()
	cumulative := 0.0
	for _, notification := range notifications {
		cumulative += notification.Probability

		if r < cumulative {
			notificationID = notification.ID
		}
	}

	// Log the decision
	err := addDecisionLog(db, notificationID, userID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error with adding log: %s", err.Error())
	}

	return notificationID
}

// SendNotifcation Send a notification to the user given their id
func SendNotifcation(userID string, notificationID string, db *db.DB) map[string]any {
	var id string

	// Check if the user exists
	u := user.GetUser(db, userID)

	// Create the user if they don't exist
	if u == nil {
		id, _ = user.SetUser(db, userID+"cool-beans") // This should change later with an actual user name
	} else {
		id, _ = u["id"].(string)
	}

	// Fetch the nofication
	notification := getNotification(id, db)

	return notification
}
