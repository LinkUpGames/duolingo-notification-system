// Package notifications is still in progress
package notifications

import (
	"database/sql"
	"fmt"
	"math/rand"
	"os"
	"server/cmd"
)

// SelectNotifcation Returns the id of the notifcation to send based on the current values
func SelectNotifcation(user_id string, database *sql.DB, variables *cmd.Variables) string {
	// Fetch the Score of the notifications
	ids := getNotificationIds(database)
	notifications := getNotifcationScores(database, user_id, ids)

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
	err := addDecisionLog(database, notificationID, user_id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error with adding log: %s", err.Error())
	}

	return notificationID
}

// SendNotifcation Send a notification to the user given their id
func SendNotifcation(user_id string, notification_id string) {
	// Check if the user exists
}
