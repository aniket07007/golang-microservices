package controllers

import(
	"net/http"
	"log"
	"strconv"
	"encoding/json"

	"github.com/aniket07007/golang-microservices/mvc/services"
	"github.com/aniket07007/golang-microservices/mvc/utils"
)

func GetUsers(resp http.ResponseWriter, req *http.Request){
	uId := req.URL.Query().Get("user_id")
	log.Printf("user id is %v", uId)

	userId, err := strconv.ParseInt(uId, 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "user id must be a number",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}
		jsonVal, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonVal)
		return
	}
	user, apiErr := services.GetUsers(userId)
	if apiErr != nil {
		// handle apiError
		jsonVal, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write([]byte(jsonVal))
		return
	}
	jsonVal, _ := json.Marshal(user)
	resp.Write(jsonVal)
}