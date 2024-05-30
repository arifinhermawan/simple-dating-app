package pgsql

var (
	queryCreateSwipeHistoryInDB = `
		INSERT INTO
			swipe_history(swiper_id,swiped_id,swipe_direction)
		VALUES (
			$1,
			$2,
			$3
		)
	`

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
