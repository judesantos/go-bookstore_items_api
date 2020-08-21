package http_utils

import (
	"encoding/json"
	"net/http"

	"github.com/judesantos/go-bookstore_utils/rest_errors"
)

func JsonSuccessResponse(w http.ResponseWriter, status int, body interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}

func JsonErrorResponse(w http.ResponseWriter, err rest_errors.IRestError) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(err.Status())
	json.NewEncoder(w).Encode(err)
}
