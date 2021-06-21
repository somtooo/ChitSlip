package routes

import (
	"fmt"
	"net/http"
)

//HandleSignIn handles Sing in
func HandleSignIn(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Signin")
}
