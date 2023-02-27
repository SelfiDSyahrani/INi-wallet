package repository

import (
	"INi-Wallet/model"
	"INi-Wallet/utils"
	"github.com/jmoiron/sqlx"
	//"gorm.io/gorm"
)

type TransactionRepository interface {
	GetByID(transaction_ID string) (model.Transaction, error)
	GetAll() ([]model.Transaction, error)
	Delete(transaction_ID string) error
}

type transactionRepository struct {
	db *sqlx.DB
}

//GetByID
func (tr *transactionRepository) GetByID(transaction_ID string) (model.Transaction, error) {
	var transaction model.Transaction
	err := tr.db.QueryRow(utils.SELECT_TRANSACTION_ID, transaction_ID).Scan(
		&transaction.Transaction_ID,
		&transaction.User_ID,
		&transaction.Money_Changer_ID,
		&transaction.Transaction_Type_ID,
		&transaction.Payment_method_id,
		&transaction.Amount,
		&transaction.Status,
		&transaction.Date_Time,
	)
	if err != nil {
		return model.Transaction{}, err
	}
	return transaction, nil
}

// GetAll
func (tr *transactionRepository) GetAll() ([]model.Transaction, error) {
	var transactions []model.Transaction
	err := tr.db.Select(transactions, utils.SELECT_TRANSACTION)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

// delete
func (tr *transactionRepository) Delete(transaction_ID string) error {
	_, err := tr.db.Exec(utils.DELETE_TRANSACTION, transaction_ID)
	if err != nil {
		return err
	}
	return nil
}

//object
func NewTransactionRepository(db *sqlx.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}
