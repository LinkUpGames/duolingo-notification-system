// Package events holds the logic for making decisions
package events

import "server/db"

// GetDecisions Get all the decision logs for a specific user
func GetDecisions(user string, db *db.DB) []*DecisionLog {
	return getDecisions(user, db)
}

func GetDecisionProbabilities(id string, db *db.DB) []*DecisionProbabilityLog {
	return getDecisionProbabilities(id, db)
}

// GetDecisionEvent Get the event logged for the decision sent
func GetDecisionEvent(id string, db *db.DB) *EventLog {
	return getDecisionEvent(id, db)
}

func CreateDecisionEvent(decisionID string, selected bool, timestamp int64, db *db.DB) bool {
	return createDecisionEvent(decisionID, selected, timestamp, db)
}
