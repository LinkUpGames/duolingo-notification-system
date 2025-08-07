package events

import (
	"fmt"
	"server/db"
)

// DecisionLog The struct that holds decision log for later computing
type DecisionLog struct {
	ID             string `json:"id"`
	UserID         string `json:"user_id"`
	NotificationID string `json:"notification_id"`
	Timestamp      int64  `json:"timestamp"`
}

// DecisionProbabilityLog The struct that holds the probability of each notification for a decision
type DecisionProbabilityLog struct {
	ID             string  `json:"id"`
	DecisionID     string  `json:"decision_id"`
	UserID         string  `json:"user_id"`
	NotificationID string  `json:"notification_id"`
	Probability    float64 `json:"probability"`
}

// EventLog The struct that holds the choice the user chose
type EventLog struct {
	DecisionID string `json:"decision_id"`
	Selected   bool   `json:"selected"`
	Timestamp  int64  `json:"timestamp"`
}

// getDecisions Get all decisions and return an array with Decision Log pointers
func getDecisions(user string, db *db.DB) []*DecisionLog {
	decisions := []*DecisionLog{}

	query := fmt.Sprintf("SELECT * FROM decisions WHERE user_id = '%s';", user)
	entries := db.GetEntries(query)

	for _, entry := range entries {
		id := entry["id"].(string)
		user := entry["user_id"].(string)
		notification := entry["notification_id"].(string)
		timestamp := entry["timestamp"].(int64)

		decision := createDecision(id, user, notification, timestamp)

		decisions = append(decisions, decision)
	}

	return decisions
}

// getDecisionEvent Given the id of a decision, return the event that is tied to it
func getDecisionEvent(id string, db *db.DB) *EventLog {
	var event *EventLog

	query := fmt.Sprintf("SELECT * FROM events WHERE decision_id = '%s'", id)
	entry := db.GetEntry(query)

	selected := entry["selected"].(bool)
	timestamp := entry["timestamp"].(int64)

	event = createEventLog(id, selected, timestamp)

	return event
}

// getDecisionProbabilities Get the probabilities for all notifications that were in the decision chosen
func getDecisionProbabilities(id string, db *db.DB) []*DecisionProbabilityLog {
	probabilities := []*DecisionProbabilityLog{}

	query := fmt.Sprintf("SELECT * FROM probabilities WHERE DECISION_ID = %s", id)
	entries := db.GetEntries(query)

	for _, entry := range entries {
		id := entry["id"].(string)
		decisionID := entry["decision_id"].(string)
		userID := entry["user_id"].(string)
		notificaitonID := entry["notification_id"].(string)
		probability := entry["probability"].(float64)

		log := createDecisionProbability(id, decisionID, userID, notificaitonID, probability)

		probabilities = append(probabilities, log)
	}

	return probabilities
}

// createDecision Create a decision struct and return the pointer
func createDecision(id string, user string, notification string, timestamp int64) *DecisionLog {
	return &DecisionLog{
		ID:             id,
		UserID:         user,
		NotificationID: notification,
		Timestamp:      timestamp,
	}
}

// createEventLog Create a new event log for the following information
func createEventLog(id string, selected bool, timestamp int64) *EventLog {
	return &EventLog{
		DecisionID: id,
		Selected:   selected,
		Timestamp:  timestamp,
	}
}

// createDecisionProbability Create a struct containing the information of a probability log for a decision made by the server
func createDecisionProbability(id string, decisionID string, userID string, notificationID string, probability float64) *DecisionProbabilityLog {
	return &DecisionProbabilityLog{
		ID:             id,
		DecisionID:     decisionID,
		UserID:         userID,
		NotificationID: notificationID,
		Probability:    probability,
	}
}
