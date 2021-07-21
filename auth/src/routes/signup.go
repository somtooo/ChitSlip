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
	"github.com/somtooo/Chit-Slip-Lib/commons/errors"
	"github.com/somtooo/Chit-Slip-Lib/commons/validation"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
)

// HandleSignUp Handles SignUp
func HandleSignUp(res http.ResponseWriter, req *http.Request) {
	validate := new(validation.Validate)
	user := db.Client.Database("user").Collection("User")
	var schema models.UserSchema
	req.ParseForm()

	for key, value := range req.Form {
		fmt.Printf("%s = %s\n", key, value)
	}

	validate.ValidateEmail(req.FormValue("email"), "Email must be valid")
	fmt.Println("this is email:", req.Form["email"])
	validate.ValidatePasswordLength(
		req.FormValue("password"), 4, 20, "Password must be between 4 and 20 char",
	)

	if validate.ValidationResult != nil {
		errors.HTTPError(res, validate, http.StatusBadRequest)
		return
	}

	err := user.FindOne(context.Background(), bson.D{{Key: "email", Value: req.FormValue("email")}}).Decode(&schema)
	log.Println(schema)

	if err != nil {
		log.Println("Error", err)
		result, error := user.InsertOne(context.Background(), models.UserSchema{Email: req.FormValue("email"), Password: schema.HashPassword(req.FormValue("password"))})
		if error != nil {
			log.Println("Insert Err: ", error)
		}

		userJwt := jwt.MapClaims{
			"id":    result.InsertedID,
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
		data, _ := json.Marshal(models.UserSchema{ID: result.InsertedID, Email: req.FormValue("email")})
		fmt.Fprint(res, string(data))
	} else {
		log.Println("Email in use")
		var badRequest errors.BadRequestError = "Email in use"
		errors.HTTPError(res, badRequest, http.StatusBadRequest)

	}

}
