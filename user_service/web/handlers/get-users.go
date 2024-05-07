package handler

import (
	"net/http"
	"user-service/web/utils"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {

	utils.SendData(w, Users)

}
