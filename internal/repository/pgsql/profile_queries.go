package pgsql

var (
	queryCreateProfileInDB = `
		INSERT INTO 
			profile(user_id,username)
		VALUES (
			$1,
			$2
		)
	`
)
