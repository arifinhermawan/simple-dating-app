package pgsql

var (
	queryCreateUserAccount = `
	INSERT INTO
		"user"(username,"password")
	VALUES (
			$1,
			$2
		)
	RETURNING id
	`

	queryGetUserAccountByUsername = `
	SELECT 
		*
	FROM
		"user"
	WHERE
		username=$1
	`
)
