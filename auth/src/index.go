package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/users/currentuser", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hi there!")
	})

	fmt.Println("Listening on port 3000!")
	http.ListenAndServe("loaclhost:3000", nil)

}
