package handler

import (
	"net/http"
	"strconv"
	"user-service/database"
	"user-service/web/utils"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(r.PathValue("id"))
	check := false
	users := utils.GetUsersFromDB()

	if users == nil {
		utils.SendError(w, http.StatusInternalServerError, "Database Error")
	} else {
		for _, user := range users {
			if user.Id == id {
				params := map[string]interface{}{
					"id": user.Id,
				}
				db := database.GetPostgreConnection()
				if db == nil {
					utils.SendError(w, http.StatusInternalServerError, "Database Error")
					return
				}

				_, err := db.NamedExec("DELETE FROM users WHERE id = :id", params)
				db.Close()

				if err != nil {
					utils.SendError(w, http.StatusInternalServerError, "Database Error")
					return
				} else {
					utils.SendJson(w, http.StatusOK, map[string]interface{}{
						"status":  true,
						"message": "Success",
					})
					return
				}
			}
		}

		if !check {
			utils.SendError(w, http.StatusBadRequest, "Bad Request")

		}
	}

}
