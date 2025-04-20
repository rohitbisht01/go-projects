package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/rohitbisht01/resturant-management/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.SignupUser())
	incomingRoutes.POST("/users/login", controllers.LoginUser())
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
	// incomingRoutes.DELETE("/users/:user_id", controllers.DeleteUser())
}
