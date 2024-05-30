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

	queryGetSwappableProfileFromDB = `
		SELECT 
			user_id,
			username,
			photo_url,
			is_verified
		FROM
			profile
		WHERE
			user_id NOT IN(%s)
		LIMIT 10
	`

	queryUpdateProfilePremiumPackageInDB = `
		UPDATE
			profile
		SET 
			is_verified=$1,
			is_infinite_scroll=$2
		WHERE
			user_id=$3
	`

	queryUpdateSwipeCountInDB = `
		UPDATE
			profile
		SET
			swipe_count = $1,
			last_swipe = $2
		WHERE
			user_id = $3
	`
)
