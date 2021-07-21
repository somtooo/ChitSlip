package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/BeatAllTech/ChitSlip/auth/src/db"
	"github.com/BeatAllTech/ChitSlip/auth/src/db/models"
	"github.com/golang-jwt/jwt"
	"github.com/somtooo/Chit-Slip-Lib/commons/errors"
	"github.com/somtooo/Chit-Slip-Lib/commons/validation"
	"go.mongodb.org/mongo-driver/bson"
)

//HandleSignIn handles Sing in
func HandleSignIn(res http.ResponseWriter, req *http.Request) {
	validate := new(validation.Validate)
	var requestError errors.BadRequestError = "Invalid credentials"
	schema := models.UserSchema{}
	user := db.Client.Database("user").Collection("User")
	validate.ValidateEmail(req.FormValue("email"), "Email must be valid")
	validate.IsPassword(req.FormValue("password"), "You must supply a password")

	if validate.ValidationResult != nil {
		errors.HTTPError(res, validate, http.StatusBadRequest)
		return
	}

	err := user.FindOne(context.Background(), bson.D{{Key: "email", Value: req.FormValue("email")}}).Decode(&schema)

	if err != nil {
		fmt.Println("SignError: ", err)
		errors.HTTPError(res, requestError, http.StatusBadRequest)
	} else {
		if answer := schema.CompareHashAndPassword(schema.Password, []byte(req.FormValue("password"))); answer == false {
			errors.HTTPError(res, requestError, http.StatusBadRequest)
			return
		}
		userJwt := jwt.MapClaims{
			"id":    schema.ID,
			"email": req.FormValue("email"),
			"iat":   time.Now().Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, userJwt)
		tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
		if err != nil {
			log.Println("Token Error: ", err)
		}
		http.SetCookie(res, &http.Cookie{
			Name:  "auth-session",
			Value: tokenString,
		})
		data, _ := json.Marshal(models.UserSchema{ID: schema.ID, Email: schema.Email})
		fmt.Fprint(res, string(data))
	}

}
