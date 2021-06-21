package validation

import (
	"encoding/json"
	"strings"
)

type validationResult struct {
	Value    string
	Msg      string
	Param    string
	Location string
}

//ValidateEmail validates email
func ValidateEmail(email string, message string) []byte {
	data, _ := json.Marshal(validationResult{Value: email, Msg: message, Param: "email", Location: "body"})
	trimmed := strings.Trim(email, " ")
	index := strings.Index(trimmed, ".")
	if len(trimmed)-4 != index {
		return data
	}
	chars := trimmed[index+1:]
	shouldContainAt := trimmed[:index]
	if (strings.Compare(chars, "com") != 0) ||
		(strings.Index(shouldContainAt, "@") <= 0) ||
		(strings.Count(shouldContainAt, "@") != 1) {
		return data
	}
	datae, _ := json.Marshal("{}")
	return datae
}

//ValidatePassword validates password
func ValidatePassword(password string, min int, max int, message string) string {
	data, _ := json.Marshal(validationResult{Value: password, Msg: message, Param: "password", Location: "body"})
	if !(min < len(password) && max > len(password)) {
		return string(data)
	}
	return "{}"
}
