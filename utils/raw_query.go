package utils

const(
	//payment_method
	INSERT_PAYMENT_METHOD = "INSERT INTO payment_method (payment_method_id,method_name) values ($1,$2)"
	SELECT_PAYMENT_METHOD_LIST = "SELECT * FROM payment_method"
	SELECT_PAYMENT_METHOD_ID = "SELECT * FROM payment_method where payment_method_id = $1"
	DELETE_PAYMENT_METHOD = "DELETE FROM payment_method where payment_method_id = $1"

	//transaction
	SELECT_TRANSACTION = "SELECT * FROM transactions" 
	SELECT_TRANSACTION_ID = "SELECT * FROM transactions where transaction_id = $1"
	DELETE_TRANSACTION = "DELETE FROM transactions where transaction_id = $1"
	
	//transfer 
	INSERT_RECORDS_TRANSFER = "INSERT INTO transactions (transaction_id,userwallet_id,money_changer_id,trasaction_type_id,amount,status,date_time) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)"
	UPDATE_BALANCE_TRANSFER_USER = "UPDATE users SET balance= balance - (SELECT * amount FROM transactions WHERE transactions_ID=$1) WHERE userwallet_id=$2"
	BALANCE_TRANSFER = "SELECT *FROM users WHERE userWallet_ID" 

	//topup
	INSERT_RECORDS_TOPUP = "INSERT INTO transactions (transaction_id,userwallet_id,money_changer_id,trasaction_type_id,amount,status,date_time) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)"
	UPDATE_BALANCE_TOPUP = "UPDATE users SET balance= balance + (SELECT amount FROM transactions WHERE transactions_ID=$1) WHERE userwallet_id=$2"
	BALANCE_TOPUP = "SELECT *FROM users WHERE userWallet_ID" 

	//payment
	INSERT_RECORDS_PAYMENT = "INSERT INTO transactions (transaction_id,userwallet_id,money_changer_id,trasaction_type_id,amount,status,date_time) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)"
	UPDATE_BALANCE_PAYMENT ="UPDATE users SET balance= balance - (SELECT amont SELECT amount FROM transactions WHERE transactions_ID=$1"
	BALANCE_PAYMENT = "SELECT *FROM users WHERE userWallet_ID" 

)
//