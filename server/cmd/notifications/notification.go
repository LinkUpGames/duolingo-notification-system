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
	// Fetch the Score of the notifications
	ids := getNotificationIds(db)
	notifications := getNotifcationScores(db, userID)

	// Get the normalized reward for all arms
	scores, maxScore := normalizeScores(notifications, variables.DEFAULT_REWARD)

	// Compute the time difference for each arm
	deltas := computeDeltas(notifications, int(variables.DEFAULT_DELTA))

	// Compute Exponentials with Recovering differnce
	expScores, total := computeExpScores(scores, deltas, maxScore, variables.TEMPERATURE)

	// Normalize to get the probabilities
	probabilities := computeProbabilities(expScores, total)

	// Sample of arm using a weight probability
	notificationID := ""
	r := rand.Float64()
	cumulative := 0.0
	for i, p := range probabilities {
		cumulative += p

		if r < cumulative {
			notificationID = ids[i]
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
