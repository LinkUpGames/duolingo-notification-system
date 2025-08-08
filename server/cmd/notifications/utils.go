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
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Days        int     `json:"-"`
}

// getNotifications Get the notifications from the database
func getNotifications(db *db.DB) []map[string]any {
	query := "SELECT * FROM notifications"
	notifications := db.GetEntries(query)

	return notifications
}

// getNotifcationScores Get the scores for the notifications stored from the database
func getNotifcationScores(db *db.DB, userID string) []map[string]any {
	query := fmt.Sprintf("SELECT * FROM scores WHERE user_id = '%s'", userID)

	results := db.GetEntries(query)

	return results
}

// computeNotificationDecay Compute the decay based on the hyperparamaters
// The penalty and factor control the penalty size for the notification score
// The cutoff controls the decay speed (larger size slower decay)
func computeNotificationDecay(notifications []*Notification, penalty float32, factor float32, cutoff int) {
	for _, notification := range notifications {
		base := float64(penalty) * float64(factor)

		power := float64(notification.Days) / float64(cutoff)

		decay := math.Pow(base, power)

		notification.Score -= decay
	}
}

// computeSoftmaxProb Compute the probability for the notifications using softmax
func computeSoftmaxProb(notifications []*Notification, explore float64) {
	for _, notification := range notifications {
		value := notification.Score / explore

		probability := math.Exp(value)

		notification.Probability = probability
	}
}

// addDecisionLog Add the selected notification to the table that saves the decision logs
func addDecisionLog(db *db.DB, selected string, notifications []*Notification) error {
	// Create the decision log
	id := uuid.New().String()
	now := time.Now().UnixMilli()

	// Fetch the selected notification
	var selectedNotification *Notification

	for _, notification := range notifications {
		if notification.ID == selected {
			selectedNotification = notification
			break
		}
	}

	query := fmt.Sprintf("INSERT INTO DECISIONS (id, user_id, notification_id, timestamp) VALUES('%s', '%s', '%s', %d);", id, selectedNotification.UserID, selectedNotification.ID, now)
	success := db.SetEntry(query)

	// Create the probability map
	if success {
		for _, notification := range notifications {
			// New id for probability
			probID := uuid.New().String()

			query := fmt.Sprintf("INSERT INTO PROBABILITIES (id, decision_id, user_id, notification_id, probability) VALUES ('%s', '%s', '%s', '%s', %f);", probID, id, notification.UserID, notification.ID, notification.Probability)

			success := db.SetEntry(query)

			if !success {
				return errors.New("error inserting probability log for decision id: " + id)
			}
		}
	} else {
		return errors.New("error inserting decision log")
	}

	return nil
}

// createNotification Creates a notification object
func createNotification(id string, userID string, score float64, timestamp int, days int, probability float64, title string, description string) *Notification {
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

// getUserNotifications Get all the notifications and their scores for the user inquiered
func getUserNotifications(userID string, db *db.DB, variables *cmd.Variables) []*Notification {
	notifications := []*Notification{}

	dbNotifications := getNotifications(db)

	// Query the scores table
	scores := getNotifcationScores(db, userID)

	// Go through all the ids of notification and query stored data associated with the user
	now := time.Now()
	for _, dbNotification := range dbNotifications {
		id, ok := dbNotification["id"].(string)
		if !ok {
			id = ""
		}

		title, ok := dbNotification["title"].(string)
		if !ok {
			title = ""
		}

		description, ok := dbNotification["description"].(string)
		if !ok {
			description = ""
		}

		var notification *Notification
		var score map[string]any = nil

		for _, _score := range scores {
			notificationID, ok := _score["id"]
			if !ok {
				notificationID = ""
			}

			if notificationID == id {
				score = _score
				break
			}
		}

		// If the score is not nil then populate with database values
		if score != nil {
			notificationScore, ok := score["score"].(float64)
			if !ok {
				notificationScore = float64(variables.Score)
			}

			timestamp, ok := score["timestamp"].(int)
			if !ok {
				timestamp = -1
			}

			milliseconds := time.UnixMilli(int64(timestamp))

			days := int(now.Sub(milliseconds).Hours() / 24)

			notification = createNotification(id, userID, notificationScore, timestamp, days, 0, title, description)
		} else {
			notification = createNotification(id, userID, float64(variables.Score), -1, variables.CutOff, 0, title, description)
		}

		notifications = append(notifications, notification)
	}

	return notifications
}

// printNotifications Prints the notification as a json string, each on a line
func printNotifications(notifications []*Notification) {
	fmt.Print("\n\n---Notifications---")
	for _, notification := range notifications {

		marshal := MarshalNotification(notification)
		content := string(marshal)

		fmt.Printf("\nNotification: %s\n", content)
	}
	fmt.Print("---Notifications---\n\n")
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
