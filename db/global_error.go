package db

import (
	"errors"
	"fmt"

	"connectrpc.com/connect"
	"github.com/jackc/pgx/v5/pgconn"
)

const (
	ForeignKeyViolation = "23503"
	NoData              = "02000"
	UniqueViolation     = "23505"
	InvalidInputSyntax  = "22P02"
)

// type ErrorHandlerDB struct {

// ConstraintName string
// FieldName      string
// }

func (store *SQLStore) DbErrorParser(err error, errorHandler map[string]string) *connect.Error {
	if err == nil {
		return nil
	}
	var pgErr *pgconn.PgError
	if err.Error() == "no rows in result set" {
		return connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("no_data_found"))
	}

	if errors.As(err, &pgErr) {
		// Check for custom exception raised by RAISE EXCEPTION in PostgreSQL
		if pgErr.Code == "P0001" { // Custom SQLSTATE from RAISE EXCEPTION
			return connect.NewError(connect.CodeInvalidArgument, fmt.Errorf(pgErr.Message))
		}
		fieldName := errorHandler[pgErr.ConstraintName]
		var errResponse error
		if fieldName != "" {
			errResponse = fmt.Errorf(fieldName)
		} else {
			errResponse = fmt.Errorf(pgErr.ConstraintName)

		}

		switch pgErr.Code {
		case ForeignKeyViolation:
			return connect.NewError(connect.CodeInvalidArgument, errResponse)
		case NoData:
			return connect.NewError(connect.CodeInvalidArgument, errResponse)
		case UniqueViolation:
			return connect.NewError(connect.CodeAlreadyExists, errResponse)
		case InvalidInputSyntax:
			return connect.NewError(connect.CodeInvalidArgument, fmt.Errorf(pgErr.Message))
		default:
			return connect.NewError(connect.CodeInternal, fmt.Errorf(pgErr.Message))
		}
	}

	return connect.NewError(connect.CodeInternal, fmt.Errorf(err.Error()))

}
