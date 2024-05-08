package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"user-service/database"
	"user-service/web/utils"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user NewUser
	id, _ := strconv.Atoi(r.PathValue("id"))
	err := json.NewDecoder(r.Body).Decode(&user)

	if err == nil {
		users := utils.GetUsersFromDB()
		if users == nil {
			utils.SendError(w, http.StatusInternalServerError, "Database Error")
		} else {
			for _, user := range users {
				if user.Id == id {
					params := map[string]interface{}{
						"id":       id,
						"name":     user.Name,
						"password": user.Password,
					}

					db := database.GetPostgreConnection()
					if db == nil {
						utils.SendError(w, http.StatusInternalServerError, "Database Error")
						return
					}

					_, err := db.NamedExec("UPDATE users SET name=:name ,password=:password WHERE id = :id", params)
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

			utils.SendError(w, http.StatusBadRequest, "Bad Request")

		}
	} else {
		utils.SendError(w, http.StatusBadRequest, "Error decoding body")
	}

}
