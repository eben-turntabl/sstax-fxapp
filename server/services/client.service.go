package services


import "example/server/models"

type ClientService interface{
	FindUserById(string)	(*models.DBResponse, error)
	FindUserByEmail(string) (*models.DBResponse,error)
	UpdateUserInfo(string ,*models.UserUpdate) (*models.DBResponse, error)
	UpdateUserInfoEmail(string ,*models.UserUpdate) (*models.DBResponse, error)
	
}