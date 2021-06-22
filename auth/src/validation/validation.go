package validation

import (
	"strings"
)

type format struct {
	Value    string
	Msg      string
	Param    string
	Location string
}

//ValidationResult

//Validate is a yeah
type Validate struct {
	ValidationResult []format
}

//ValidateEmail validates email
func (e *Validate) ValidateEmail(email string, message string) {
	f := format{Value: email, Msg: message, Param: "email", Location: "body"}
	trimmed := strings.Trim(email, " ")
	index := strings.Index(trimmed, ".")
	if len(trimmed)-4 != index {
		e.ValidationResult = append(e.ValidationResult, f)
		return
	}
	chars := trimmed[index+1:]
	shouldContainAt := trimmed[:index]
	if (strings.Compare(chars, "com") != 0) ||
		(strings.Index(shouldContainAt, "@") <= 0) ||
		(strings.Count(shouldContainAt, "@") != 1) {
		e.ValidationResult = append(e.ValidationResult, f)
		return
	}

}

//ValidatePassword validates password
func (e *Validate) ValidatePassword(password string, min int, max int, message string) {
	f := format{Value: password, Msg: message, Param: "password", Location: "body"}
	if !(min < len(password) && max > len(password)) {
		e.ValidationResult = append(e.ValidationResult, f)
		return
	}

}

func (e *Validate) SerializeErrors() []struct {
	Message string
	Field   string `json:"Field,omitempty" bson:"Field,omitempty"`
} {
	serialized := make([]struct {
		Message string
		Field   string `json:"Field,omitempty" bson:"Field,omitempty"`
	}, len(e.ValidationResult))
	for i, v := range e.ValidationResult {
		serialized[i] = struct {
			Message string
			Field   string `json:"Field,omitempty" bson:"Field,omitempty"`
		}{v.Msg, v.Param}
	}

	return serialized
}
