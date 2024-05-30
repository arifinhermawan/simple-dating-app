package pgsql

import (
	"context"
	"database/sql"

	"github.com/arifinhermawan/simple-dating-app/internal/app/infrastructure/configuration"
)

type infraProvider interface {
	GetConfig() *configuration.AppConfig
}

type psqlProvider interface {
	// BeginTx starts a transaction.
	//
	// The provided context is used until the transaction is committed or rolled back.
	// If the context is canceled, the sql package will roll back
	// the transaction. Tx.Commit will return an error if the context provided to
	// BeginTx is canceled.
	//
	// The provided TxOptions is optional and may be nil if defaults should be used.
	// If a non-default isolation level is used that the driver doesn't support,
	// an error will be returned.
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)

	// ExecContext executes a query without returning any rows. The args are for any placeholder parameters in the query.
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)

	// QueryContext executes a query that returns rows, typically a SELECT.
	// The args are for any placeholder parameters in the query.
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)

	// QueryRowContext executes a query that is expected to return at most one row.
	// QueryRowContext always returns a non-nil value.
	// Errors are deferred until the *Row.Scan method is called.
	//  If the query selects no rows, the *Row.Scan will return ErrNoRows.
	// Otherwise, the *Row.Scan scans the first selected row and discards the rest.
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

type Repository struct {
	infra infraProvider
	db    psqlProvider
}

func NewRepository(infa infraProvider, db psqlProvider) *Repository {
	return &Repository{
		infra: infa,
		db:    db,
	}
}

func (repo *Repository) BeginTX(ctx context.Context, options *sql.TxOptions) (*sql.Tx, error) {
	return repo.db.BeginTx(ctx, options)
}

func (repo *Repository) Commit(tx *sql.Tx) error {
	return tx.Commit()
}

func (repo *Repository) Rollback(tx *sql.Tx) error {
	return tx.Rollback()
}
