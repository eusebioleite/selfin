package utils

import (
	"errors"

	"github.com/ncruces/go-sqlite3"
)

func IsConstraintViolation(err error) bool {
	var sqliteErr *sqlite3.Error
	if errors.As(err, &sqliteErr) && sqliteErr.ExtendedCode() == sqlite3.CONSTRAINT_FOREIGNKEY {
		return true
	} else {
		return false
	}
}
