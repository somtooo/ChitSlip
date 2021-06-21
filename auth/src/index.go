package main

import (
	"fmt"
	"net/http"

	"github.com/BeatAllTech/ChitSlip/auth/src/routes"
)

func main() {
	http.HandleFunc("/api/users/currentuser", routes.HandleCurrentUser)
	http.HandleFunc("/api/users/signup", routes.HandleSignUp)
	http.HandleFunc("/api/users/signout", routes.HandleSignOut)
	http.HandleFunc("/api/users/signin", routes.HandleSignIn)

	fmt.Println("Listening on port 3000!")
	http.ListenAndServe(":3000", nil)

}
