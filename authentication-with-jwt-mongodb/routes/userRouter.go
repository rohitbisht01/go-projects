package routes

import (
	middleware "github.com/rohitbisht01/authentication-with-jwt-mongodb//middlewares"
	controller "github.com/rohitbisht01/authentication-with-jwt-mongodb/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/user/:user_id", controller.GetUserById())
}
