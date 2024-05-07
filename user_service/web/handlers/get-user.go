package handler

import (
	"net/http"
	"strconv"
	"user-service/web/utils"
)

func GetUser(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(r.PathValue("id"))
	check := false
	for _, user := range Users {
		if user.Id == id {
			utils.SendData(w, user)
			check = true
			break
		}
	}
	if !check {
		utils.SendError(w, http.StatusBadRequest, "Bad Request")
	}

}
