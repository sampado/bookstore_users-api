package mysqlutils

import (
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/sampado/bookstore_utils-go/rest_errors"
)

const (
	ErrorNoRows = "no rows in result set"
)

func ParseError(err error) *rest_errors.RestError {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return rest_errors.NewNotFoundError("row not found ")
		}
		return rest_errors.NewInternalServerError("error when trying to cast error", err)
	}

	switch sqlErr.Number {
	case 1062:
		return rest_errors.NewBadRequestError("invalid data")
	}

	return rest_errors.NewInternalServerError(fmt.Sprintf("error processing request: %s", sqlErr.Error()), err)
}
