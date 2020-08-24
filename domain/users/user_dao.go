package users

import (
	"fmt"
	"github.com/Anatol-e/bookstore_users_api/utils/errors"
)

var usersDB = make(map[int64]*User)

func (user *User) Get() *errors.RestErr {
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
	usersDB[user.Id] = user
	return nil
}