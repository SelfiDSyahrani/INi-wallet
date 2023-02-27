package repository

import (
	"INi-Wallet/model"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(transaction model.Transaction) error
	GetByID(transaction_ID string) (model.Transaction, error)
	GetAll() ([]model.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

//create a new transaction
func (tr *transactionRepository) Create(transaction model.Transaction) error {
	return tr.db.Create(transaction).Error
}

func (tr * transactionRepository) GetByID(transaction_ID string) (model.Transaction, error) {
	 transaction := model.Transaction{}
	 err := tr.db.First(transaction,transaction_ID).Error
	 return transaction,err
}

func (tr *transactionRepository)GetAll() ([]model.Transaction, error)  {
	var transactions []model.Transaction
	err :=tr.db.Find(&transactions).Error
	return transactions,err
}