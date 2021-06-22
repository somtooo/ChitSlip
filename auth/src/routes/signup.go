package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/BeatAllTech/ChitSlip/auth/src/errors"
	"github.com/BeatAllTech/ChitSlip/auth/src/validation"
)

// HandleSignUp Handles SignUp
func HandleSignUp(res http.ResponseWriter, req *http.Request) {
	validate := new(validation.Validate)
	validate.ValidateEmail(req.FormValue("email"), "Email must be valid")
	validate.ValidatePassword(req.FormValue("password"), 4, 20, "Password must be between 4 and 20 char")

	if validate.ValidationResult != nil {
		errors.HTTPError(res, validate, http.StatusBadRequest)
		return
	}

	res.WriteHeader(http.StatusOK)
	data, _ := json.Marshal("{}")
	fmt.Fprint(res, string(data))
}
