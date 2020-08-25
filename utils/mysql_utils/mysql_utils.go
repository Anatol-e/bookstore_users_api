package mysql_utils

import (
	"fmt"
	"github.com/Anatol-e/bookstore_users_api/utils/errors"
	"github.com/go-sql-driver/mysql"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		return errors.NewNotFoundError(fmt.Sprintf("data wasn't created : %s", err.Error()))
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewInternalServerError(fmt.Sprintf("data is already exists : %s", err.Error()))
	default:
		return errors.NewInternalServerError(fmt.Sprintf("internal error : %s", err.Error()))
	}
}
