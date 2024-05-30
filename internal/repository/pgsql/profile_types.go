package pgsql

import "database/sql"

type CreateProfileReq struct {
	Tx       *sql.Tx
	UserID   int64
	Username string
}
