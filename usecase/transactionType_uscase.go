package usecase

import (
	"dev_selfi/model"
	"dev_selfi/repository"
)

type TransactionTypeUsecase interface {
	TransactionTypeGetByID(transactionType_ID string) (model.TransactionType, error)
	TransactionTypeGetAll() ([]model.TransactionType, error)
}

type transactionTypeUsecase struct {
	TransactionTypeRepo repository.TransactionTypeRepository	
}

func  (ttu * transactionTypeUsecase) TransactionTypeGetByID(transactionType_ID string) (model.TransactionType, error){
	return ttu.TransactionTypeRepo.GetByID(transactionType_ID)
}

func (ttu * transactionTypeUsecase) TransactionTypeGetAll() ([]model.TransactionType, error){
	return ttu.TransactionTypeRepo.GetAll()
}