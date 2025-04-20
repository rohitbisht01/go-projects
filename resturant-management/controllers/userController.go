package controllers

import (
	"github.com/gin-gonic/gin"
)

func LoginUser() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func SignupUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		// var ctx, err = context.WithTimeout(context.Background(), 10*time.Second)

	}
}

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func HashPassword(password string) string {}

func VerfifyPassword(userPassword string, providedPassword string) bool {

}
