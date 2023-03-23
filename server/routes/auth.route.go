package routes

import (
	"example/server/controllers"

	"github.com/gin-gonic/gin"
)

type AuthRouteController struct{
	authController controllers.AuthController
}

func NewAuthRouteController(authController controllers.AuthController) AuthRouteController{
	return AuthRouteController{authController}
}

func (rc *AuthRouteController) AuthRoute(rg *gin.RouterGroup){
	router := rg.Group("/auth")

	router.POST("/register", rc.authController.UserRegister)
	router.POST("/login",rc.authController.UserLogin)
}





































