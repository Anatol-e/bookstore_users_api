package services

import (
	"github.com/Anatol-e/bookstore_users_api/domain/users"
	"github.com/Anatol-e/bookstore_users_api/utils/date"
	"github.com/Anatol-e/bookstore_users_api/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.DateCreated = date.GetNowString()
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}
