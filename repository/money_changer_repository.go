package repository

import (
	"dev_selfi/model"
	"dev_selfi/utils"

	"github.com/jmoiron/sqlx"
)

type MoneyChangerRepository interface {
	GetByID(MoneyChanger_ID string) (model.MoneyChanger, error)
	GetAll() ([]model.MoneyChanger, error)
}

type moneyChangerRepository struct {
	db *sqlx.DB
}

func (mc *moneyChangerRepository)GetByID(MoneyChanger_ID string) (model.MoneyChanger, error) {
	var MoneyChanger model.MoneyChanger
	err := mc.db.QueryRow(utils.SELECT_MONEY_CHANGER_ID,MoneyChanger_ID).Scan(
	 &MoneyChanger.Money_changer_ID,
	 &MoneyChanger.Name,
	 &MoneyChanger.Exchange_rate,
	 &MoneyChanger.Currency_type,
	 &MoneyChanger.Country,

	)
	if err != nil {
		return model.MoneyChanger{}, err
	}
	return MoneyChanger, nil
}

func (mc *moneyChangerRepository ) GetAll() ([]model.MoneyChanger, error) {
	var moneyChanger []model.MoneyChanger
	err := mc.db.Select(moneyChanger, utils.SELECT_TRANSACTION)
	if err != nil {
		return nil, err
	}
	return moneyChanger, nil

}