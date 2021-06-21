package routes

import (
	"fmt"
	"net/http"
)

//HandleCurrentUser handles what happens when the current user visits page
func HandleCurrentUser(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Current User")
}
