package repository

import (
	"INi-Wallet/model"

	"gorm.io/gorm"
)

type MoneyChangerRepository interface {
	Create(Mc model.MoneyChanger) error
	GetByID(id string) (*model.MoneyChanger, error)
	GetAll() ([]model.TransactionType, error)
	Update()
}

type moneyChangerRepository struct {
	db *gorm.DB
}

func (m *moneyChangerRepository) Create(Mc model.MoneyChanger) error {
	return m.db.Create(Mc).Error
}

func (m *moneyChangerRepository) Delete(id string) error {
	money := model.MoneyChanger{}
	money.Money_Changer_ID = id
	return m.db.Where("id = ?", id).Delete(&money).Error
}

func (m *moneyChangerRepository) Update(moneyChanger model.MoneyChanger) error {
	return m.db.Save(moneyChanger).Error
}

func (m *moneyChangerRepository) GetByID(id string) (*model.MoneyChanger, error) {
	money := model.MoneyChanger{}
	err := m.db.Where("id = ?", id).First(&money).Error
	if err != nil {
		return nil, err
	}
	return &money, err
}
