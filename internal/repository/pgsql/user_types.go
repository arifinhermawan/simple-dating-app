package pgsql

import "database/sql"

type UserAccount struct {
	ID       int64
	Username string
	Password string
}

type CreateUserReq struct {
	Tx       *sql.Tx
	Username string
	Password string
}
