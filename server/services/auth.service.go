package services

import "example/server/models"


type AuthService interface{
	UserLogin(*models.ClientLogin) (*models.DBResponse,error)
	UserRegister(*models.ClientSignup) (*models.DBResponse,error)
}
