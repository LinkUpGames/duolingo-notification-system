// Package cmd is still in progress
package cmd

import (
	"database/sql"
	"fmt"
	"math"
	"math/rand"
	"server/db"
	"time"
)

// SelectNotifcation Returns the id of the notifcation to send based on the current values
func SelectNotifcation(user_id string, database *sql.DB, variables *Variables) string {
	// Fetch the Score of the notifications
	ids := getNotificationIds(database)
	notifications := getNotifcationScores(database, user_id, ids)

	// Get the normalized reward for all arms
	scores := []float32{}
	var maxScore float32 = -1
	for _, notification := range notifications {
		var score float32

		reward, ok := notification["reward"]

		// Parse Value correctly
		if ok {
			score = reward.(float32)
		} else {
			score = variables.DEFAULT_REWARD
		}

		if score > maxScore {
			maxScore = score
		}

		// Add the scores to the array
		scores = append(scores, score)
	}

	// Compute the time difference for each arm
	deltas := []float64{}
	for _, notification := range notifications {
		now := time.Now().UnixMilli()
		last_seen, ok := notification["timestamp"].(int64)
		var delta float64

		if ok {
			delta = math.Abs(float64(last_seen - now))
		} else {
			delta = float64(variables.DEFAULT_DELTA)
		}

		deltas = append(deltas, delta)
	}

	// Compute Exponentials with Recovering differnce
	expScores := []float64{}
	total := 0.0
	temperature := variables.TEMPERATURE

	for i := 0; i < len(notifications); i++ {
		normScore := (scores[i] - maxScore) / temperature

		val := float64(normScore) * deltas[i]
		expScore := math.Exp(val)

		expScores = append(expScores, expScore)
		total = expScore // Save the total for the average
	}

	// Normalize to get the probabilities
	probabilities := []float64{}
	for i := 0; i < len(notifications); i++ {
		val := expScores[i] / total
		probabilities = append(probabilities, val)
	}

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

	return notificationID
}

// getNotificationIds Get the ids of all the notifications from the database
func getNotificationIds(database *sql.DB) []string {
	ids := []string{}

	query := "SELECT ID FROM NOTIFICATIONS"

	notifications := db.GetEntries(database, query)

	for _, notification := range notifications {
		id, ok := notification["id"]

		if ok {
			ids = append(ids, id.(string))
		}
	}

	return ids
}

// getNotifcationScores Get the scores for the notifications stored from the database
func getNotifcationScores(database *sql.DB, userID string, notifications []string) []map[string]any {
	query := fmt.Sprintf("SELECT * FROM scores WHERE user_id = %d", userID)

	results := db.GetEntries(database, query)

	return results
}
