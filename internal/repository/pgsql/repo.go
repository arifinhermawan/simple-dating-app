package pgsql

import (
	"github.com/jmoiron/sqlx"
)

var (
	funcSQLXNamed = sqlx.Named
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}
