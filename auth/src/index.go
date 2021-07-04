package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/BeatAllTech/ChitSlip/auth/src/db"
	"github.com/BeatAllTech/ChitSlip/auth/src/errors"
	"github.com/BeatAllTech/ChitSlip/auth/src/middlewares/currentuser"
	"github.com/BeatAllTech/ChitSlip/auth/src/routes"
)

func main() {
	if os.Getenv("JWT_KEY") == "" {
		panic("JWT_KEY must be definied")
	}
	http.Handle("/api/users/currentuser", currentuser.CurrentUser(http.HandlerFunc(routes.HandleCurrentUser)))
	http.HandleFunc("/api/users/signup", routes.HandleSignUp)
	http.HandleFunc("/api/users/signout", routes.HandleSignOut)
	http.HandleFunc("/api/users/signin", routes.HandleSignIn)
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		var notFound errors.BadRequestError = "No link Found"
		errors.HTTPError(res, notFound, http.StatusBadRequest)
	})

	db.ConnectDb()
	fmt.Println("Listening on port 3000!")
	http.ListenAndServe(":3000", nil)

}
