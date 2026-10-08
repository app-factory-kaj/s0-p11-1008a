package handlers

import (
	"encoding/json"
	"net/http"

	"greeter/internal/gen"
)

// WriteJSONError writes an OpenAPI Error-shaped body for a request that
// failed parameter binding before reaching the strict handler (e.g. a
// missing required query parameter).
func WriteJSONError(w http.ResponseWriter, _ *http.Request, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	_ = json.NewEncoder(w).Encode(gen.Error{
		Code:        http.StatusBadRequest,
		Message:     "name is required",
		Description: err.Error(),
	})
}
