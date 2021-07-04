package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/BeatAllTech/ChitSlip/auth/src/middlewares/currentuser"
)

//HandleCurrentUser handles what happens when the current user visits page
func HandleCurrentUser(res http.ResponseWriter, req *http.Request) {

	if user := req.Context().Value(currentuser.Key); user != nil {
		data, _ := json.Marshal(user)
		fmt.Fprintf(res, string(data))
	} else {
		fmt.Println("Key not found: ", currentuser.Key)

	}

}
