package mapper

import (
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
)

func MapMySQLError(err error) error {
	mapping := NewPersistanceErrMapper()
	if err == nil {
		return nil
	}

	if errors.Is(err, sql.ErrNoRows) {
		return mapping.NotFound
	}

	var myErr *mysql.MySQLError
	if errors.As(err, &myErr) {
		switch myErr.Number {
		case 1062:
			return mapping.Duplicate
		case 1452:
			return mapping.RelatedNotFound
		case 1451:
			return mapping.StillReferenced
		case 1048:
			return mapping.RequiredField
		case 1406:
			return mapping.ValueTooLong
		case 1213, 1205:
			return mapping.Retryable
		case 2006, 2013:
			return mapping.Connection
		}
	}

	return err
}
