package utils

import (
	"encoding/json"
	"net/http"
)

// RespondWithError will return error response.
func (utilsModule UtilsModule) RespondWithError(w http.ResponseWriter, code int, msg string) {
	utilsModule.RespondWithJSON(w, code, map[string]string{"error": msg})
}

// RespondWithJSON will return success response.
func (utilsModule UtilsModule) RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
