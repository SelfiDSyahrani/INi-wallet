package utils

const(
	//payment_method
	INSERT_PAYMENT_METHOD = "INSERT INTO payment_method (payment_method_id,method_name) values ($1,$2)"
	SELECT_PAYMENT_METHOD_LIST = "SELECT * FROM payment_method"
	SELECT_PAYMENT_METHOD_ID = "SELECT * FROM payment_method where payment_method_id = $1"
	//UPDATE_PAYMENT_METHOD = "UPDATE payment_method set payment_method_id=$1,method_name =$2 where id=$3"
	DELETE_PAYMENT_METHOD = "DELETE FROM payment_method where payment_method_id = $1"

<<<<<<< HEAD
	//transfer

	//
	INSERT_USER = "INSERT INTO user (userWallet_ID, name, email, phone, password) values ($!,$2, $3, $4, $5)"
	DELETE_USER = "DELETE FROM user where userWallet_ID = $1"
	SELECT_USER_ID = "SELECT * FROM user where userWallet_ID = $1"
	SELECT_USER_LIST = "SELECT * FROM payment_method"
	UPDATE_USER= "UPDATE payment_method set "




)
=======
	//transaction
	SELECT_TRANSACTION = "SELECT * FROM transactions"
	SELECT_TRANSACTION_ID = "SELECT * FROM transactions where transaction_id = $1"
	DELETE_TRANSACTION = "DELETE FROM transactions where transaction_id = $1"
)
>>>>>>> 70dbb88e8dcdea7190f7104e570f79f11b616a32
