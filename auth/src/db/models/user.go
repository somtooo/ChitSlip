package models

import (
	"log"

	"github.com/BeatAllTech/ChitSlip/auth/src/db"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

//UserSchema Document
type UserSchema struct {
	ID       interface{} `bson:"_id,omitempty" json:"_id,omitempty"`
	Email    string      ` bson:"email" json:"email"`
	Password []byte      `json:"password,omitempty" json:"password,omitempty"`
}

//NewUser makes a user collection
func NewUser() *mongo.Collection {
	return db.Client.Database("user").Collection("User")

}

//HashPassword hashes a passwrd from string and updates it
func (schema UserSchema) HashPassword(password string) []byte {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	if err != nil {
		log.Println("Error: ", err)
	}
	return hashedPassword
}

//CompareHashAndPassword compares a bcrypt hashed password with its possible plaintext equivalent.
//Returns true on success, or false on failure.
func (schema UserSchema) CompareHashAndPassword(hashedPassword, password []byte) bool {
	if err := bcrypt.CompareHashAndPassword(hashedPassword, password); err != nil {
		return false
	}
	return true
}
