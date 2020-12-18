package responses

import (
	"encoding/json"
	"net/http"
)

func SuccessResponse(w http.ResponseWriter, status int, data interface{}) error {


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
	return nil
}

func ServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]string{"err":"Oops Its not you, its me, please wait"})
	return
}

func BadRequest(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]string{"err":message})
	return
}