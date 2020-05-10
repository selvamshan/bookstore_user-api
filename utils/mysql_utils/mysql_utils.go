package mysql_utils

import(
	"strings"
	"errors"
	"github.com/go-sql-driver/mysql"
	"github.com/selvamshan/bookstore_utils-go/rest_errors"
)

const(	
	ErrorNoRow = "no rows in result set"
)

func ParseError(err error) rest_errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if ! ok {
		if strings.Contains(err.Error(), ErrorNoRow) {
			return rest_errors.NewNotFoundError("no records matching givne id")
		}
		return rest_errors.NewInternalServerError("error parseing database response", errors.New("database error"))
	}
	switch sqlErr.Number {
	case 1062:
		return rest_errors.NewBadRequestError("invalid data")
	}
	return rest_errors.NewInternalServerError("error parseing database response", errors.New("database error"))

}