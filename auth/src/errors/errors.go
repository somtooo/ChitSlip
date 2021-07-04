package errors

import (
	"encoding/json"
	"net/http"
)

//CustomErrors this interface is used to be able to make sure all errors are sent in the same format
type CustomErrors interface {
	SerializeErrors() struct {
		Errors []struct {
			Message string `json:"message"`
			Field   string `json:"field,omitempty"`
		} `json:"errors"`
	}
}

//BadRequestError type is used to allow the use of string messages as customErrors when a resource is not found
type BadRequestError string

//HTTPError handles request error sending to the client
func HTTPError(w http.ResponseWriter, error CustomErrors, code int) {
	data, _ := json.Marshal(error.SerializeErrors())
	http.Error(w, string(data), code)
}

//SerializeErrors does..
func (e BadRequestError) SerializeErrors() struct {
	Errors []struct {
		Message string `json:"message"`
		Field   string `json:"field,omitempty"`
	} `json:"errors"`
} {

	d := struct {
		Errors []struct {
			Message string `json:"message"`
			Field   string `json:"field,omitempty"`
		} `json:"errors"`
	}{}

	serialized := make([]struct {
		Message string `json:"message"`
		Field   string `json:"Field,omitempty" bson:"Field,omitempty"`
	}, 0)

	serialized = append(serialized, struct {
		Message string `json:"message"`
		Field   string `json:"Field,omitempty" bson:"Field,omitempty"`
	}{string(e), ""})
	d.Errors = []struct {
		Message string "json:\"message\""
		Field   string "json:\"field,omitempty\""
	}(serialized)

	return d
}
