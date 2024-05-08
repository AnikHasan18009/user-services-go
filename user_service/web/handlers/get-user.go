package handler

import (
	"net/http"
	"strconv"
	"user-service/web/utils"
)

func GetUser(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(r.PathValue("id"))
	check := false
	users := utils.GetUsersFromDB()

	if users == nil {
		utils.SendError(w, http.StatusInternalServerError, "Database Error")
	} else {
		for _, user := range users {
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

}
