package pgsql

var (
	queryGetTodaysSwipedList = `
		SELECT 
			swiped_id
		FROM 
			swipe_history
		WHERE 
			swiper_id=$1
		AND 
			swiped_id != $1
		AND
			created_at
		BETWEEN
			$2
		AND
			$3
		LIMIT 10
	`
)
