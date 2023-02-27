package repository

import (
	"INi-Wallet/model"

	"gorm.io/gorm"
)

type PaymentMethodRepository interface {
	Create(paymentMethod model.PaymentMethod) error
	Update(paymentMethod model.PaymentMethod) error
	Delete(PaymentMethod_ID string) error
	GetByID(PaymentMethod_ID string) (model.PaymentMethod, error)
	GetAll() ([]model.PaymentMethod, error)
}

type paymentMethodRepository struct {
	db *gorm.DB
}

// Creates a new payment method
func (pmr *paymentMethodRepository) Create(paymentMethod model.PaymentMethod) error {
	return pmr.db.Create(paymentMethod).Error
}

// Update
func (pmr *paymentMethodRepository) Update(paymentMethod model.PaymentMethod) error {
	return pmr.db.Save(paymentMethod).Error
}

// Delete
func (pmr *paymentMethodRepository) Delete(PaymentMethod_ID string) error {
	paymentMethod := model.PaymentMethod{}
	paymentMethod.Payment_Method_ID = PaymentMethod_ID
	return pmr.db.Delete(paymentMethod).Error
}

//GetByID
func (pmr *paymentMethodRepository) GetByID(PaymentMethod_ID string) (model.PaymentMethod, error) {
	paymentMethod := model.PaymentMethod{}
	err := pmr.db.First(paymentMethod, PaymentMethod_ID).Error
	return paymentMethod, err
}

//GetAll
func (pmr *paymentMethodRepository)GetAll() ([]model.PaymentMethod, error)  {
	var paymentMethods []model.PaymentMethod
	err := pmr.db.Find(&paymentMethods).Error
	return paymentMethods,err
}

func NewPaymentMethodRepository(db * gorm.DB) PaymentMethodRepository {
	return &paymentMethodRepository{
		db: db,
	}
}