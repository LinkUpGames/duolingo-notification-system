package notifications

import (
	"errors"
	"fmt"
	"math"
	"server/db"
	"time"

	"github.com/google/uuid"
)

// getNotificationIds Get the ids of all the notifications from the database
func getNotificationIds(db *db.DB) []string {
	ids := []string{}

	query := "SELECT ID FROM notifications"

	notifications := db.GetEntries(query)

	for _, notification := range notifications {
		id, ok := notification["id"]

		if ok {
			ids = append(ids, id.(string))
		}
	}

	return ids
}

// getNotification Get a notification given the id
func getNotification(id string, db *db.DB) map[string]any {
	query := fmt.Sprintf("SELECT ID FROM NOTIFICATIONS WHERE ID = %s", id)

	notification := db.GetEntry(query)

	_, ok := notification["id"].(string)

	if !ok {
		return nil
	}

	return notification
}

// getNotifcationScores Get the scores for the notifications stored from the database
func getNotifcationScores(db *db.DB, userID string) []map[string]any {
	query := fmt.Sprintf("SELECT * FROM scores WHERE user_id = %s", userID)

	results := db.GetEntries(query)

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

	for i := range scores {
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

	for i := range scores {
		val := scores[i] / total

		probabilities = append(probabilities, val)
	}

	return probabilities
}

// addDecisionLog Add the selected notification to the table that saves the decision logs
func addDecisionLog(db *db.DB, notification string, user string) error {
	id := uuid.New().String()
	now := time.Now().UnixMilli()

	query := fmt.Sprintf("INSERT INTO DECISIONS (id, user_id, notification_id, timestamp) VALUES(%s, %s, %s, %d);", id, user, notification, now)

	err := db.SetEntry(query)

	if err {
		return errors.New("error instering decision log")
	}

	return nil
}
