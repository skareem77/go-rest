package handler

import (
	"encoding/json"
	"net/http"
)

//ResponseJSON send response
func ResponseJSON(w http.ResponseWriter, status int, payload interface{}) {

	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "applicaiton/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

//RespondError send error response
func RespondError(w http.ResponseWriter, code int, message string) {
	ResponseJSON(w, code, map[string]string{"error": message})
}
