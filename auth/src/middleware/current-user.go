package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
)

type user struct {
	CurrentUser jwt.MapClaims `json:"currentUser"`
}

//IsLoggedIn is a Middleware that checks if user is logged in
func IsLoggedIn(handler http.Handler) func(http.ResponseWriter, *http.Request) {
	type currentUser string
	ctx := context.WithValue(context.Background(), currentUser("currentUser"), user{CurrentUser: nil})
	return func(res http.ResponseWriter, req *http.Request) {
		cookie := req.Cookies()
		if len(cookie) == 0 {
			req = req.Clone(ctx)
			handler.ServeHTTP(res, req)
			return
		}

		token, err := jwt.Parse(cookie[0].Value, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_KEY")), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			ctx = context.WithValue(context.Background(), currentUser("currentUser"), user{CurrentUser: claims})
			req = req.Clone(ctx)
			handler.ServeHTTP(res, req)
		} else {
			fmt.Println("Token verify Error: ", err)
			req = req.Clone(ctx)
			handler.ServeHTTP(res, req)
		}
	}

}
