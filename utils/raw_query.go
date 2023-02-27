package utils

const(
	//payment_method
	INSERT_PAYMENT_METHOD = "INSERT INTO payment_method (payment_method_id,method_name) values ($1,$2)"
	SELECT_PAYMENT_METHOD_LIST = "SELECT * FROM payment_method"
	SELECT_PAYMENT_METHOD_ID = "SELECT * FROM payment_method where payment_method_id = $1"
	//UPDATE_PAYMENT_METHOD = "UPDATE payment_method set payment_method_id=$1,method_name =$2 where id=$3"
	DELETE_PAYMENT_METHOD = "DELETE FROM payment_method where payment_method_id = $1"

	//transfer
)