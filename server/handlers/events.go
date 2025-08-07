package handlers

import (
	"encoding/json"
	"net/http"
	"server/cmd"
	"server/cmd/events"
	"server/db"
)

func GetDecisionProbabilitiesHandler(ctx *cmd.AppContext, w http.ResponseWriter, r *http.Request) {
	// Context
	db := ctx.Ctx.Value(cmd.DATABASE).(*db.DB)

	// Parameters
	id := r.URL.Query().Get("id")

	// Fetch
	logs := events.GetDecisionEvent(id, db)

	// Marshall
	jsonByte, err := json.Marshal(logs)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err != nil {
		emptyArray := []any{}

		empty, _ := json.Marshal(emptyArray)

		w.Write(empty)
	} else {
		w.Write(jsonByte)
	}
}
