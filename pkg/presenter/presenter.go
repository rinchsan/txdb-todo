package presenter

import (
	"encoding/json"
	"net/http"
)

func Success(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	body := Response{Message: "success"}
	encode(w, body)
}

func Encode(w http.ResponseWriter, body interface{}) {
	w.WriteHeader(http.StatusOK)
	encode(w, body)
}

func BadRequest(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusBadRequest)
	body := Response{Message: message}
	encode(w, body)
}

func Forbidden(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusForbidden)
	body := Response{Message: message}
	encode(w, body)
}

func Error(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	body := Response{Message: err.Error()}
	encode(w, body)
}

func encode(w http.ResponseWriter, body interface{}) {
	json.NewEncoder(w).Encode(body)
}

type Response struct {
	Message string `json:"message"`
}
