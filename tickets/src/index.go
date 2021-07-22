package main

import (
	"fmt"
	//"github.com/somtooo/Chit-Slip-Lib/commons/middlewares/currentuser"
	"net/http"
)

func main() {
	http.HandleFunc("/api/tickets/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprintf(writer, "Hello")
	})

	//http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
	//	var notFound errors.BadRequestError = "No link Found"
	//	errors.HTTPError(res, notFound, http.StatusBadRequest)
	//})

	fmt.Println("Listening on port 3000!")
	_ = http.ListenAndServe(":3000", nil)

}
