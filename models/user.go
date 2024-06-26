package models

import (
	"github.com/avijeetpandey/event-booking/db"
	"github.com/avijeetpandey/event-booking/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

// save user to the database
func (u *User) Save() error {
	query := `INSERT INTO users(email, password) VALUES (?,?)`

	preparedStmtnt, err := db.GlobalDB.Prepare(query)

	if err != nil {
		return err
	}

	defer preparedStmtnt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := preparedStmtnt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return err
	}

	u.ID = id
	u.Password = hashedPassword

	return err
}
