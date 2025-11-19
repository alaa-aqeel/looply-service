package database

import (
	"errors"

	domain_errors "github.com/alaa-aqeel/looply-app/src/core/Domain/errors"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

func MapPgError(err error) error {
	if err == nil {
		return nil
	}

	var pgErr *pgconn.PgError

	if !errors.As(err, &pgErr) {
		return err // not a postgres error
	}

	switch pgErr.Code {

	case pgerrcode.UniqueViolation:
		return &domain_errors.DatabaseError{
			Code:    domain_errors.ErrDBUnique,
			Field:   pgErr.ConstraintName,
			Message: "duplicate value violates unique constraint",
			Details: pgErr.Detail,
		}

	case pgerrcode.ForeignKeyViolation:
		return &domain_errors.DatabaseError{
			Code:    domain_errors.ErrDBForeignKey,
			Field:   pgErr.ConstraintName,
			Message: "foreign key constraint failed",
			Details: pgErr.Detail,
		}

	case pgerrcode.NotNullViolation:
		return &domain_errors.DatabaseError{
			Code:    domain_errors.ErrDBNotNull,
			Field:   pgErr.ColumnName,
			Message: "required field is missing",
			Details: pgErr.Detail,
		}

	case pgerrcode.CheckViolation:
		return &domain_errors.DatabaseError{
			Code:    domain_errors.ErrDBCheck,
			Field:   pgErr.ConstraintName,
			Message: "check constraint validation failed",
			Details: pgErr.Detail,
		}

	case pgerrcode.InvalidTextRepresentation:
		return &domain_errors.DatabaseError{
			Code:    domain_errors.ErrDBInvalidInput,
			Field:   pgErr.ColumnName,
			Message: "invalid input format",
			Details: pgErr.Detail,
		}

	default:
		return &domain_errors.DatabaseError{
			Code:    domain_errors.ErrUnknown,
			Message: pgErr.Message,
			Details: pgErr.Detail,
		}
	}
}
