package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/rohitbisht01/authentication-with-jwt-mongodb/database"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func Signup() {}

func Login() {}

func HashPassword() {}

func VerifyPassword() {}

func GetUsers() {}

func GetUserById() {}
