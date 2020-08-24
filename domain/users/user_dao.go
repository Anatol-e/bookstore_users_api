package users

import (
	"fmt"
	"github.com/Anatol-e/bookstore_users_api/datasources/mysql/users_db"
	"github.com/Anatol-e/bookstore_users_api/utils/errors"
)

const (
	queryInsertUser = "INSERT INTO users (firstname, lastname, email, date_created) VALUES(?,?,?,?);"
	queryGetUser    = "SELECT id, firstname, lastname, date_created FROM users WHERE id = ?;"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.ClientDB.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to get user %d with an error %s",
			user.Id, err.Error()))
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.ClientDB.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}
	user.Id = userId
	return nil
}
