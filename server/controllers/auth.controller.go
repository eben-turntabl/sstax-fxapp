package controllers

import (
	"example/server/models"
	"example/server/services"
	"fmt"
	"net/http"


	"github.com/gin-gonic/gin"
)

/// removed client serivce
type AuthController struct{
	authService services.AuthService
	clientService services.ClientService
	
}

func NewAuthController(authService services.AuthService,clientService services.ClientService)AuthController{
	return AuthController{authService,clientService }
}
func (ac AuthController) UserLogin(ctx *gin.Context){
	var user *models.ClientLogin
	if err := ctx.ShouldBindJSON(&user) ;err != nil{
		return
			}
	loginUser, err := ac.authService.UserLogin(user)
	fmt.Println(loginUser,"hellooo")

	if err != nil{
		
			ctx.JSON(http.StatusNotFound,gin.H{"status":"error","message":err.Error()})
		
	}else {
		ctx.IndentedJSON(http.StatusAccepted,gin.H{"status":"success","data":loginUser})

	}


}
func (ac *AuthController) UserRegister(ctx *gin.Context){
	var user *models.ClientSignup

	if err := ctx.ShouldBindJSON(&user) ;err != nil{
return
	}

	
	
	newUser, err := ac.authService.UserRegister(user)

	fmt.Println(newUser, err)
	if err !=nil{
		
			ctx.JSON(http.StatusConflict,gin.H{"status":"error","message":err.Error()})
		
	}else {
		
		ctx.IndentedJSON(http.StatusCreated,newUser)

	}


}