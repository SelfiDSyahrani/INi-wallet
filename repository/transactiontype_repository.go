package repository

import (
	"INi-Wallet/model"

	"gorm.io/gorm"
)

type TransactionTypeRepository interface {
	CreateTransactionType(transactionType model.TransactionType) error
	GetTransactionTypeByID(transaction_Type_ID string) (model.TransactionType, error)
	GetAllTransactionTypes() ([]model.TransactionType, error)
}

type transactionTyperRepository struct {
	db *gorm.DB
}

// CreateTransactionType implements TransactionTypeRepository
func (t *transactionTyperRepository) CreateTransactionType(transactionType model.TransactionType) error {
	return t.db.Create(transactionType).Error
}

// GetAllTransactionTypes implements TransactionTypeRepository
func (t *transactionTyperRepository) GetAllTransactionTypes() ([]model.TransactionType, error) {
	var types []model.TransactionType
	err := t.db.Find(&types).Error
	return types, err
}

// GetTransactionTypeByID implements TransactionTypeRepository
func (t *transactionTyperRepository) GetTransactionTypeByID(id string) (model.TransactionType, error) {
	typeID := model.TransactionType{}
	err := t.db.First(typeID, id).Error
	return typeID, err
}

func NewTrasactionTypeRepository(db *gorm.DB) TransactionTypeRepository {
	return &transactionTyperRepository{
		db: db,
	}
}
