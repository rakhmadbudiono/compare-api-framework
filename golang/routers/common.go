package routers

import (
	"encoding/json"
	"net/http"
)

func handleJSONResponse(w http.ResponseWriter, v interface{}) {
	message, err := json.Marshal(v)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(message)
}