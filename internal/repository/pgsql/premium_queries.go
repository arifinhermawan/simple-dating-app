package pgsql

var (
	queryCreatePurchaseHistory = `
		INSERT INTO 
			purchase_history(user_id,upgrade_type)
		VALUES (
			$1,
			$2
		)
	`
)
