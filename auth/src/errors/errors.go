package errors

import (
	"encoding/json"
	"net/http"
)

//CustomErrors
type CustomErrors interface {
	SerializeErrors() []struct {
		Message string
		Field   string `json:"Field,omitempty" bson:"Field,omitempty"`
	}
}

//HttpError
func HttpError(w http.ResponseWriter, error CustomErrors, code int) {
	data, _ := json.Marshal(error.SerializeErrors())
	http.Error(w, string(data), code)
}

type DatabaseConnectionError struct {
	statusCode int
	reason     string
}
