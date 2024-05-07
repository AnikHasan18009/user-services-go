package handler

import (
	"net/http"
	"strconv"
	"user-service/web/utils"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(r.PathValue("id"))
	check := false
	for index, user := range Users {
		if user.Id == id {
			Users = append(Users[:index], Users[index+1:]...)
			check = true
			break
		}
	}
	if check {
		utils.SendJson(w, http.StatusOK, map[string]interface{}{
			"status":  true,
			"message": "Success",
		})
	} else {
		utils.SendError(w, http.StatusBadRequest, "Bad Request")
	}

}
