package pgsql

import "database/sql"

type UserAccount struct {
	ID       int64  `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
}

type CreateUserReq struct {
	Tx       *sql.Tx
	Username string
	Password string
}
