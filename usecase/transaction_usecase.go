package usecase

import (
	"dev_selfi/model"
	"dev_selfi/repository"
)

type TransactionUscase interface {
	Transfer(transaction model.Transaction,user model.User) error
	TopuP(transaction model.Transaction) error
	Payment(transaction model.Transaction) error
}

type transactionUscase struct {
	transactionRepo repository.TransactionRepository
}

func (t *transactionUscase) Transfer(transaction model.Transaction,user model.User) error {
	return t.transactionRepo.InsertTransactionTransfer(transaction)
	if transaction.Amount >= user.Balance {
		return error.Error("dddddddddddd")
	}
	//Proses dengan transactions
	//1. transctionRepo.InsertIntoTabelTransction
	// bila sukses ambil amount ransaction dan Update saldo user menjadi saldo + amount_transfer

	//3. select data saldo terakhir setelah di update where ID_USER = id

	return nil
}

func (t *transactionUscase) TopuP(transaction model.Transaction) error {
	return t.transactionRepo.InsertTransactionTopUp(transaction)

	//Proses dengan transactions
	//1. transctionRepo.InsertIntoTabelTransction
	// bila sukses ambil amount ransaction dan Update saldo user menjadi saldo + amount_transfer

	//3. select data saldo terakhir setelah di update where ID_USER = id

	return nil
}

func (t *transactionUscase) Payment(transaction model.Transaction) error {
	return t.transactionRepo.InsertTransactionPayment(transaction)

	return nil
}