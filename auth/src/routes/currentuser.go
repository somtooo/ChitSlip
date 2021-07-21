package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/somtooo/Chit-Slip-Lib/commons/middlewares/currentuser"
)

//HandleCurrentUser handles what happens when the current user visits page
func HandleCurrentUser(res http.ResponseWriter, req *http.Request) {

	if user := req.Context().Value(currentuser.Key); user != nil {
		fmt.Println("this is user ", user)
		data, _ := json.Marshal(user)
		fmt.Fprintf(res, string(data))
	} else {
		fmt.Println("Key not found: ", currentuser.Key)

	}

}
