package services

import (
	"my-portofolio/models"

	"go.mongodb.org/mongo-driver/mongo"
)

/*
	@Author: DevProblems(Sarang Kumar)
	@YTChannel: https://www.youtube.com/channel/UCVno4tMHEXietE3aUTodaZQ
*/
type UserService interface {
	CreateUser(*models.User) (*mongo.InsertOneResult, error)
	GetUser(*string) (*models.User, error)
	GetAll() ([]*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(*string) error
}