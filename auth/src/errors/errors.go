package errors

import (
	"encoding/json"
	"net/http"
)

//CustomErrors this interface is used to be able to make sure all errors are sent in the same format
type CustomErrors interface {
	SerializeErrors() []struct {
		Message string
		Field   string `json:"Field,omitempty" bson:"Field,omitempty"`
	}
}

//NotFoundError type is used to allow the use of string messages as customErrors when a resource is not found
type NotFoundError string

//HTTPError handles request error sending to the client
func HTTPError(w http.ResponseWriter, error CustomErrors, code int) {
	data, _ := json.Marshal(error.SerializeErrors())
	http.Error(w, string(data), code)
}

func (e NotFoundError) SerializeErrors() []struct {
	Message string
	Field   string `json:"Field,omitempty" bson:"Field,omitempty"`
} {

	serialized := make([]struct {
		Message string
		Field   string `json:"Field,omitempty" bson:"Field,omitempty"`
	}, 0)

	serialized = append(serialized, struct {
		Message string
		Field   string `json:"Field,omitempty" bson:"Field,omitempty"`
	}{string(e), ""})

	return serialized
}
