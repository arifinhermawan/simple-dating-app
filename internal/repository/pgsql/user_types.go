package pgsql

import "database/sql"

// ------------------
// | request struct |
// ------------------
type CreateUserReq struct {
	Tx       *sql.Tx
	Username string
	Password string
}

// -------------------
// | response struct |
// -------------------
type UserAccount struct {
	ID       int64
	Username string
	Password string
}
