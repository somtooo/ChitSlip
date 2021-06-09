package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Listening on port 3000!")
	http.ListenAndServe("loaclhost:3000", nil)

}
