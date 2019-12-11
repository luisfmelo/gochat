package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func JSON(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	w.WriteHeader(statusCode)
	res := Response{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		_, _ = fmt.Fprintf(w, "%s", err.Error())
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, err.Error(), nil)
		return
	}
	JSON(w, http.StatusBadRequest, "Bad Request", nil)
}

func LogChatMessage(from string, to string, msg string) {
	log.Println(fmt.Sprintf("%s to %s: %s", from, to, msg))
}
