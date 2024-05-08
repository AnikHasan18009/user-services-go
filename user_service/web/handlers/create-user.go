package handler

import (
	"encoding/json"
	"net/http"
	"user-service/database"
	"user-service/web/utils"
)

type NewUser struct {
	Name     string `json:"name" db:"name"`
	Password string `json:"password" db:"password"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	var user NewUser

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Error decoding body")
	} else {
		db := database.GetPostgreConnection()

		if db == nil {
			utils.SendError(w, http.StatusInternalServerError, "Database Error")
		} else {
			_, err = db.NamedExec("INSERT INTO users(NAME,PASSWORD) VALUES(:name,:password)", user)
			db.Close()
			if err != nil {
				utils.SendError(w, http.StatusInternalServerError, "Database Error")
			} else {
				utils.SendJson(w, http.StatusOK, map[string]interface{}{
					"status":  true,
					"message": "Success"})
			}
		}

	}

}
