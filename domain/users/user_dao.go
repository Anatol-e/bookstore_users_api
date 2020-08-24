package users

import (
	"fmt"
	"github.com/Anatol-e/bookstore_users_api/datasources/mysql/users_db"
	"github.com/Anatol-e/bookstore_users_api/utils/date"
	"github.com/Anatol-e/bookstore_users_api/utils/errors"
)

var usersDB = make(map[int64]*User)

func (user *User) Get() *errors.RestErr {
	if err := users_db.ClientDB.Ping(); err != nil {
		panic(err)
	}

	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

func (user *User) Save() *errors.RestErr {
	if usersDB[user.Id] != nil {
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user))
	}
	user.DateCreated = date.GetNowString()
	usersDB[user.Id] = user
	return nil
}
