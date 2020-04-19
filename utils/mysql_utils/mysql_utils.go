package mysql_utils

import(
	"strings"
	"github.com/go-sql-driver/mysql"
	"github.com/selvamshan/bookstore_user-api/utils/errors"
)

const(	
	errorNoRow = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if ! ok {
		if strings.Contains(err.Error(), errorNoRow) {
			return errors.NewNotFoundError("no records matching givne id")
		}
		return errors.NewInternalServerError("error parseing database response")
	}
	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("invalid data")
	}
	return errors.NewInternalServerError("error processing request")

}