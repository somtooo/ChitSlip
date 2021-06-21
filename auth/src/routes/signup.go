package routes

import (
	"fmt"
	"net/http"

	"github.com/BeatAllTech/ChitSlip/auth/src/validation"
)

// HandleSignUp Handles SignUp
func HandleSignUp(res http.ResponseWriter, req *http.Request) {
	emailErr := validation.ValidateEmail(req.FormValue("email"), "Email must be valid")
	passErr := validation.ValidatePassword(req.FormValue("password"), 4, 20, "Password must be between 4 and 20 char")
	s := [1]string{string(emailErr) + "," + "\n" + string(passErr)}

	if (len(emailErr) > 1) || (len(passErr) > 1) {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(res, s)
		return
	}

	fmt.Fprint(res, s)
}
