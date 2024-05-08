package utils

import (
	"user-service/database"
)

type User struct {
	Id       int    `json:"id"  db:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func GetUsersFromDB() []User {
	var users []User
	db := database.GetPostgreConnection()
	if db == nil {
		return nil
	}
	err := db.Select(&users, "SELECT  * FROM users")
	db.Close()

	if err != nil {

		return nil
	}

	return users

}
