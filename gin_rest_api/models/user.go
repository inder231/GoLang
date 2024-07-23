package models

import (
	"rest-api/db"
	"rest-api/utils"
)


type User struct {
	ID int64 ``
	Email string `binding:"required"`
	Password string `binding:"required"`
}


func (u User) Save() error {
	query := `INSERT INTO users(email, password) VALUES(?, ?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}
	
	result, err := stmt.Exec(u.Email, hashedPassword)

	stmt.Close()
	
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	u.ID = id
	
	return err
}

