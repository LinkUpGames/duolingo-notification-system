package notifications

import (
	"database/sql"
	"fmt"
	"math"
	"server/db"
	"time"

	"github.com/google/uuid"
)

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

// normalizeScores Normalize the scores for notifications and return a map with the scores per notification
// and also the maxScore seen from all scores (as a touple)
func normalizeScores(notifications []map[string]any, def float32) ([]float32, float32) {
	var maxScore float32 = -1.0
	scores := []float32{}

	for _, notification := range notifications {
		var score float32

		reward, ok := notification["reward"]

		if ok {
			score = reward.(float32)
		} else {
			score = def // Set default
		}

		if score > maxScore {
			maxScore = score
		}

		scores = append(scores, score)
	}

	return scores, maxScore
}

// computeDeltas Compute the difference between the now and the timestamp of the notification
func computeDeltas(notifications []map[string]any, def int) []int {
	now := time.Now().UnixMilli()
	deltas := []int{}

	for _, notification := range notifications {
		lastSeen, ok := notification["timestamp"].(int)
		var delta int

		if ok {
			diff := lastSeen - int(now)
			diffAbs := int(math.Abs(float64(diff)))

			delta = diffAbs
		} else {
			delta = def
		}

		deltas = append(deltas, delta)
	}

	return deltas
}

// Compute the Exponential recovering difference for the scores
func computeExpScores(scores []float32, deltas []int, maxScore float32, temperature float32) ([]float64, float64) {
	expScores := []float64{}
	total := 0.0

	for i := 0; i < len(scores); i++ {
		norm := (scores[i] - maxScore) / temperature

		val := float64(norm) * float64(deltas[i])

		score := math.Exp(val)
		expScores = append(expScores, score)

		total += score
	}

	return expScores, total
}

// computeProbabilities Compute the probabilities for all notifications to be sent to the user
func computeProbabilities(scores []float64, total float64) []float64 {
	probabilities := []float64{}

	for i := 0; i < len(scores); i++ {
		val := scores[i] / total

		probabilities = append(probabilities, val)
	}

	return probabilities
}

// addDecisionLog Add the selected notification to the table that saves the decision logs
func addDecisionLog(db *sql.DB, notification string, user string) error {
	id := uuid.New().String()
	now := time.Now().UnixMilli()

	query := fmt.Sprintf("INSERT INTO DECISIONS (id, user_id, notification_id, timestamp) VALUES(%s, %s, %s, %d);", id, user, notification, now)

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
