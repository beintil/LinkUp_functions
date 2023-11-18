package responder

import (
	"encoding/json"
	"github.com/beintil/LinkUp_function/LUError"
	"net/http"
)

// APIResponse представляет собой общий формат API-ответа.
type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
}

func SendJSONResponse(w http.ResponseWriter, e LUError.Error, data interface{}) {
	switch {
	case e != nil:
		sendErrorResponse(w, e)
	default:
		sendOkResponse(w, data)
	}
}

func sendOkResponse(w http.ResponseWriter, data interface{}) {
	response := APIResponse{
		Success: true,
		Data:    data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(&response)
	if err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
	}
}

func sendErrorResponse(w http.ResponseWriter, e LUError.Error) {
	response := APIResponse{
		Success: false,
		Error:   e.Error(),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.HTTPCode())

	err := json.NewEncoder(w).Encode(&response)
	if err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
	}
}
