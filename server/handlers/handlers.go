// Package handlers
package handlers

import (
	"encoding/json"
	"net/http"
	"server/cmd"
	"server/cmd/notifications"
	"server/db"
)

type Handler func(ctx *cmd.AppContext, w http.ResponseWriter, r *http.Request)

// Middleware The middleware function that adds the context to request
func Middleware(ctx *cmd.AppContext, handler Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(ctx, w, r)
	}
}

// SendNotificationHandler Select an arm based on the user
func SendNotificationHandler(ctx *cmd.AppContext, w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")

	variables := ctx.Ctx.Value(cmd.VARIABLES).(*cmd.Variables)
	db := ctx.Ctx.Value(cmd.DATABASE).(*db.DB)

	// Select a notification
	id := notifications.SelectNotifcation(userID, variables, db)

	notification := notifications.SendNotifcation(userID, id, db)

	jsonBytes, err := json.Marshal(notification)

	if notification == nil || err != nil {

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusExpectationFailed)

		w.Write([]byte("Error"))
	} else {

		// Set Response Headers
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write(jsonBytes)
	}
}
