package main

import (
	"encoding/json"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]any{
		"status": status,
		"error":  err.Error(),
	})
}

func RespondAny(w http.ResponseWriter, httpStatus int, obj any) {
	w.WriteHeader(httpStatus)
	_ = json.NewEncoder(w).Encode(map[string]any{
		"status": httpStatus,
		"result": obj,
	})
}

func RespondPostRequest(w http.ResponseWriter, answer bool) {
	if answer {
		RespondAny(w, http.StatusOK, "your request has been fulfilled")
	} else {
		RespondWithError(w, http.StatusInternalServerError, ErrorServerSide)
	}
}
