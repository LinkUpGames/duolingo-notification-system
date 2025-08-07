// Package notifications is still in progress
package notifications

import (
	"encoding/json"
	"fmt"
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

	// Update Scores
	success := updateNotificationScores(notifications, db)

	if !success {
		fmt.Printf("Error with updating notification scores!\n")
	}

	// Log the decision
	decisionID, err := addDecisionLog(db, notification.ID, notifications)
	if err != nil {
		fmt.Printf("Error with adding log: %s\n", err.Error())
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

// CreateNotification Creates a notification object
func CreateNotification(id string, userID string, score float64, timestamp int, days int, probability float64, title string, description string) *Notification {
	notification := &Notification{
		ID:          id,
		UserID:      userID,
		Score:       score,
		Timestamp:   timestamp,
		Days:        days,
		Probability: probability,
		Title:       title,
		Description: description,
	}

	return notification
}

// UpdateNotificationScores Update the notifications scores for the users
func UpdateNotificationScores(notifications []*Notification, db *db.DB) bool {
	return updateNotificationScores(notifications, db)
}
