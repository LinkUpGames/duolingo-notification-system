package handlers

import (
	"encoding/json"
	"net/http"
	"server/cmd"
	"server/cmd/events"
	"server/cmd/notifications"
	"server/cmd/user"
	"server/db"
)

// GetUsersHandler Returns all of the users currently saved in the database
func GetUsersHandler(ctx *cmd.AppContext, w http.ResponseWriter, r *http.Request) {
	// Context
	db := ctx.Ctx.Value(cmd.DATABASE).(*db.DB)

	users := user.GetUsers(db)

	jsonBytes, err := json.Marshal(users)

	// Always return true since we can send an empty array
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err != nil {
		emptyArray := []any{}
		value, _ := json.Marshal(emptyArray)
		w.Write(value)
	} else {
		w.Write(jsonBytes)
	}
}

// GetUserNotificationsHandler Get the notifications and their score given the user
func GetUserNotificationsHandler(ctx *cmd.AppContext, w http.ResponseWriter, r *http.Request) {
	// Context
	db := ctx.Ctx.Value(cmd.DATABASE).(*db.DB)
	variables := ctx.Ctx.Value(cmd.VARIABLES).(*cmd.Variables)

	// Parameters
	userID := r.URL.Query().Get("user_id")

	notifications := notifications.GetUserNotifications(userID, db, variables)

	jsonBytes, err := json.Marshal(notifications)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err != nil {
		emptyArray := []any{}
		value, _ := json.Marshal(emptyArray)
		w.Write(value)
	} else {
		w.Write(jsonBytes)
	}
}

// GetUserDecisionsHandler Get all of the decisions that the user has done so far to update the scores of notifications
func GetUserDecisionsHandler(ctx *cmd.AppContext, w http.ResponseWriter, r *http.Request) {
	// Context
	db := ctx.Ctx.Value(cmd.DATABASE).(*db.DB)

	// Parameters
	userID := r.URL.Query().Get("user_id")

	decisions := events.GetDecisions(userID, db)

	jsonBytes, err := json.Marshal(decisions)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err != nil {
		emptyArray := []any{}

		empty, _ := json.Marshal(emptyArray)
		w.Write(empty)
	} else {
		w.Write(jsonBytes)
	}
}
