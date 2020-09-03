package users

import (
	"fmt"
	"github.com/Anatol-e/bookstore_users_api/datasources/mysql/users_db"
	"github.com/Anatol-e/bookstore_users_api/logger"
	"github.com/Anatol-e/bookstore_users_api/utils/errors"
	"github.com/Anatol-e/bookstore_users_api/utils/mysql_utils"
	"github.com/go-sql-driver/mysql"
)

const (
	queryInsertUser             = "INSERT INTO users (firstname, lastname, email, date_created, password) VALUES(?,?,?,?,?);"
	queryGetUser                = "SELECT id, firstname, lastname, email, date_created FROM users WHERE id = ?;"
	queryUpdateUser             = "UPDATE users SET firstname = ?, lastname = ?, email = ?, password = ? WHERE id = ?;"
	queryDeleteUser             = "DELETE FROM users WHERE id = ?;"
	queryFindByEmailAndPassword = "SELECT id, firstname, lastname, email, date_created FROM users WHERE email=? AND password = ?;"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.ClientDB.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
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

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Password)
	if err != nil {
		sqlErr, ok := err.(*mysql.MySQLError)
		if !ok {
			return errors.NewInternalServerError(fmt.Sprintf("error when trying to get user %d with an error %s",
				user.Id, err.Error()))
		}

		switch sqlErr.Number {
		case 1062:
			return errors.NewInternalServerError("email is already exists")
		default:
			return errors.NewInternalServerError(fmt.Sprintf("mysql cant save the user %d", user.Id))
		}
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}
	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.ClientDB.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Password, user.Id)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.ClientDB.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.Id); err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) FindByEmailAndPassword() *errors.RestErr {
	stmt, err := users_db.ClientDB.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Email, user.Password)
	user.Password = ""
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
		logger.Error("error when trying to get user by email and password", getErr)
		return errors.NewInternalServerError("database error")
	}

	return nil
}
