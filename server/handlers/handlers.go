package handlers

import (
	"fmt"
	"net/http"
	"server/cmd"
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
	name := r.URL.Query().Get("user_id")

	fmt.Fprintf(w, "%s", name)
}
