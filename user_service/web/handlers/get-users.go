package handler

import (
	"net/http"
	"user-service/web/utils"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {

	users := utils.GetUsersFromDB()

	if users == nil {
		utils.SendError(w, http.StatusInternalServerError, "Database Error")
	} else {

		utils.SendData(w, users)
	}

}
