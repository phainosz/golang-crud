package utils

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func WriteErrorResponse(w http.ResponseWriter, httpStatus int, errorResponse ErrorResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)

	errorResponseJson, err := json.Marshal(errorResponse)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "error handling the response."}`))
		return
	}

	w.Write(errorResponseJson)
}

func WriteSuccessResponse(w http.ResponseWriter, httpStatus int, response any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)

	if response != nil {
		if erro := json.NewEncoder(w).Encode(response); erro != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "error handling response."}`))
		}
	}
}
