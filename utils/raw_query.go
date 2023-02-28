package utils

const(
	//payment_method
	INSERT_PAYMENT_METHOD = "INSERT INTO payment_method (payment_method_id,method_name) values ($1,$2)"
	SELECT_PAYMENT_METHOD_LIST = "SELECT * FROM payment_method"
	SELECT_PAYMENT_METHOD_ID = "SELECT * FROM payment_method where payment_method_id = $1"
	//UPDATE_PAYMENT_METHOD = "UPDATE payment_method set payment_method_id=$1,method_name =$2 where id=$3"
	DELETE_PAYMENT_METHOD = "DELETE FROM payment_method where payment_method_id = $1"

	//transfer

	//user
	INSERT_USER = "INSERT INTO user (userWallet_ID, name, email, phone, password, balance) values ($!,$2, $3, $4, $5, $6)"
	DELETE_USER = "DELETE FROM user where userWallet_ID = $1"
	SELECT_USER_ID = "SELECT * FROM user where userWallet_ID = $1"
	SELECT_USER_LIST = "SELECT * FROM payment_method"
	UPDATE_USER_BY_ID= "UPDATE user set set name=$1, balance=$2 where id=$3"
	UPDATE_USER_PASS= "UPDATE user set set password where id=$1 AND email=$2"


	SELECT_TRANSACTION = "SELECT * FROM transactions"
	SELECT_TRANSACTION_ID = "SELECT * FROM transactions where transaction_id = $1"
	DELETE_TRANSACTION = "DELETE FROM transactions where transaction_id = $1"

)
	//transaction


