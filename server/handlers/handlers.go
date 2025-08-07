// Package handlers
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
		fmt.Printf("\nRequest Information\nURL: [%s] | HOST: [%s]\n", r.URL, r.RemoteAddr)

		handler(ctx, w, r)
	}
}
