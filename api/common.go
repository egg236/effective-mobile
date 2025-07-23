package api

import (
	"effective-mobile/entities"
	"encoding/json"
	"fmt"
	"net/http"
)

func getRecordFromRequest(r *http.Request) (*entities.Record, error) {
	var record entities.Record
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&record); err != nil {
		return nil, fmt.Errorf("failed to decode request body: %w", err)
	}

	if record.Validate() != nil {
		return nil, record.Validate()
	}

	return &record, nil
}

func writeResponse(w http.ResponseWriter, data interface{}, status int, isSuccess bool) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	resp := entities.Response{
		IsSuccess: isSuccess,
		Result:    data,
	}
	if data != nil {
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, fmt.Sprintf("failed to encode response: %s", err.Error()), http.StatusInternalServerError)
		}
	}
}
