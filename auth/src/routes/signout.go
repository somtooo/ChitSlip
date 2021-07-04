package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HandleSignOut handles signout
func HandleSignOut(res http.ResponseWriter, req *http.Request) {
	cookie, _ := req.Cookie("auth-session")
	cookie = &http.Cookie{
		Name:  "auth-session",
		Value: "",
	}
	http.SetCookie(res, cookie)

	type empty struct{}
	data, _ := json.Marshal(empty{})
	fmt.Fprintf(res, string(data))
}
