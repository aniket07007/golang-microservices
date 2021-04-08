package services

import (
	"github.com/aniket07007/golang-microservices/mvc/domain"
	"github.com/aniket07007/golang-microservices/mvc/utils"
)

func GetUsers(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}