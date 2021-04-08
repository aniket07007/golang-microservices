package domain

import (
	"fmt"
	"net/http"

	"github.com/aniket07007/golang-microservices/mvc/utils"
)

var (
	Users = map[int64]*User{
		123 : &User{
			Id: 123,
			FirstName: "Aniket",
			LastName: "Amin",
			Email: "aniketamin007@gmail.com",
		},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	
	if user := Users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message: fmt.Sprintf("user %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code: "not_found",
	}
}