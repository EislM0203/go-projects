package models

import (
	"errors"

	"traunseenet.com/rest-api/db"
	"traunseenet.com/rest-api/utils"
)

type User struct {
	ID int64 		`json:"id"`
	Email string 	`json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) Save() error{
	query := "INSERT INTO users (email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	
	pwd, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = pwd

	res, err := stmt.Exec(u.Email, u.Password)
	if err != nil {
		return err
	}

	userId, err := res.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = userId
	return nil
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)
	if err != nil {
		return errors.New("User does not exist")
	}
	return utils.VerifyPassword(retrievedPassword, u.Password)
}