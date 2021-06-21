package routes

import (
	"fmt"
	"net/http"
)

// HandleSignOut handles signout
func HandleSignOut(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Signout")
}
