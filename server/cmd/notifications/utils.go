package notifications

import (
	"errors"
	"fmt"
	"math"
	"server/cmd"
	"server/db"
	"time"

	"github.com/google/uuid"
	"github.com/mroth/weightedrand/v2"
)

// Notification The notification with the scores for a specific user
type Notification struct {
	ID          string  `json:"id"`
	UserID      string  `json:"user_id"`
	Score       float64 `json:"score"`
	Timestamp   int     `json:"timestamp"`
	Selected    int     `json:"selected"`
	Probability float64 `json:"probability"`

	Delta int
}

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

// getNotifcationScores Get the scores for the notifications stored from the database
func getNotifcationScores(db *db.DB, userID string) []map[string]any {
	query := fmt.Sprintf("SELECT * FROM scores WHERE user_id = '%s'", userID)

	results := db.GetEntries(query)

	return results
}

// Compute the Exponential recovering difference for the scores
// This will modify the scores of the actual notification and return the total value sum value of the Scores
func computeExpScores(notifications []*Notification, temperature float64, maxScore float64) float64 {
	total := 0.0

	for _, notification := range notifications {
		norm := (notification.Score - maxScore) / temperature
		val := norm * float64(notification.Delta)
		score := math.Exp(val)
		notification.Score = score

		total += score
	}

	return total
}

// computeProbabilities Compute the probabilities for all notifications to be sent to the user
func computeProbabilities(notifications []*Notification, total float64) {
	for _, notification := range notifications {
		val := notification.Score / total

		notification.Probability = val
	}
}

// addDecisionLog Add the selected notification to the table that saves the decision logs
func addDecisionLog(db *db.DB, notification string, user string) error {
	id := uuid.New().String()
	now := time.Now().UnixMilli()

	query := fmt.Sprintf("INSERT INTO DECISIONS (id, user_id, notification_id, timestamp) VALUES('%s', '%s', '%s', %d);", id, user, notification, now)

	err := db.SetEntry(query)

	if err {
		return errors.New("error instering decision log")
	}

	return nil
}

// createNotification Creates a notification object
func createNotification(id string, userID string, score float64, timestamp int, selected int, delta int, probability float64) *Notification {
	notification := &Notification{
		ID:          id,
		UserID:      userID,
		Score:       score,
		Timestamp:   timestamp,
		Selected:    selected,
		Delta:       delta,
		Probability: probability,
	}

	return notification
}

// getUserNotifications Get all the notifications and their scores for the user inquiered
func getUserNotifications(userID string, db *db.DB, variables *cmd.Variables) []*Notification {
	notifications := []*Notification{}

	ids := getNotificationIds(db)

	// Query the scores table
	scores := getNotifcationScores(db, userID)

	// Go through all the ids of notification and query stored data associated with the user
	now := time.Now().UnixMilli()
	for _, id := range ids {
		var notification *Notification
		var score map[string]any = nil

		for _, _score := range scores {
			notificationID := _score["id"]

			if notificationID == id {
				score = _score
				break
			}
		}

		// If the score is not nil then populate with database values
		if score != nil {
			reward := score["reward"].(float64)
			timestamp := score["timestamp"].(int)
			selected := score["selected"].(int)

			diff := timestamp - int(now)
			diffAbs := int(math.Abs(float64(diff)))

			notification = createNotification(id, userID, reward, timestamp, selected, diffAbs, 0)
		} else {
			notification = createNotification(id, userID, float64(variables.DEFAULT_REWARD), -1, 0, int(variables.DEFAULT_DELTA), 0)
		}

		notifications = append(notifications, notification)
	}

	return notifications
}

// printNotifications Prints the notification as a json string, each on a line
func printNotifications(notifications []*Notification) {
	fmt.Print("\n---Notifications---")
	for _, notification := range notifications {

		marshal := MarshalNotification(notification)
		content := string(marshal)

		fmt.Printf("\nNotification: %s\n", content)
	}
	fmt.Print("---Notifications---\n")
}

// selectRandom Select a notification at random given the weighted probability of each notification
func selectRandom(notifications []*Notification) *Notification {
	choices := []weightedrand.Choice[*Notification, int]{}

	for _, notification := range notifications {
		val := int(notification.Probability * 100)

		choice := weightedrand.NewChoice(notification, val)

		choices = append(choices, choice)
	}

	// Select one at random
	chooser, err := weightedrand.NewChooser(choices...)

	var result *Notification
	if err != nil {
		result = nil
	} else {
		result = chooser.Pick()
	}

	return result
}
