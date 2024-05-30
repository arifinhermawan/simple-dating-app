package pgsql

var (
	queryCreateProfileInDB = `
		INSERT INTO 
			profile(user_id,username,photo_url)
		VALUES (
			$1,
			$2,
			$3
		)
	`

	queryGetProfileByUserIDFromDB = `
		SELECT 
			*
		FROM 
			profile
		WHERE
			user_id = $1

	`
)
