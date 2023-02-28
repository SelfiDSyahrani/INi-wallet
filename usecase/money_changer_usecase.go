package usecase

import (
	"dev_selfi/model"
	"dev_selfi/repository"

)

type MoneyChangerUsecase interface{
	MoneyChangerById(MoneyChanger_ID string) (model.MoneyChanger, error)
	MoneyChangerAll() ([]model.MoneyChanger, error)
}

type moneyChangerUsecase struct {
	moneyChangerRepo repository.MoneyChangerRepository
}

func (mcu * moneyChangerUsecase)MoneyChangerById(MoneyChanger_ID string) (model.MoneyChanger, error) {
	return mcu.moneyChangerRepo.GetByID(MoneyChanger_ID)
}

func (mcu * moneyChangerUsecase)MoneyChangerAll() ([]model.MoneyChanger, error)  {
	return mcu.moneyChangerRepo.GetAll()
}