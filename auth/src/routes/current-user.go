package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//HandleCurrentUser handles what happens when the current user visits page
func HandleCurrentUser(res http.ResponseWriter, req *http.Request) {

	key := "currentUser"
	if user := req.Context().Value(key); user != nil {
		data, _ := json.Marshal(user)
		fmt.Fprintf(res, string(data))
	}
	fmt.Println("key not found: ", key)

}
