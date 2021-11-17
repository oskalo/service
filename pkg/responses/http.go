package responses

import (
	"encoding/json"
	"net/http"
)

const (
	jsonMimeType = "application/json"
	contentType  = "Content-Type"
)

type MessageResponse struct {
	Message string `json:"message"`
}

func JSONResponse(w http.ResponseWriter, code int, payload interface{}) error {
	w.Header().Set(contentType, jsonMimeType)
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(payload)
}

func JSONErrorResponse(w http.ResponseWriter, code int, err error) error {
	w.Header().Set(contentType, jsonMimeType)
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(MessageResponse{err.Error()})
}
