package main

import (
	"fmt"
	"net/http"

	"github.com/BeatAllTech/ChitSlip/auth/src/errors"
	"github.com/BeatAllTech/ChitSlip/auth/src/routes"
)

func main() {
	http.HandleFunc("/api/users/currentuser", routes.HandleCurrentUser)
	http.HandleFunc("/api/users/signup", routes.HandleSignUp)
	http.HandleFunc("/api/users/signout", routes.HandleSignOut)
	http.HandleFunc("/api/users/signin", routes.HandleSignIn)
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		var notFound errors.NotFoundError = "No link Found"
		errors.HTTPError(res, notFound, http.StatusBadRequest)
	})

	fmt.Println("Listening on port 3000!")
	http.ListenAndServe(":3000", nil)

}
