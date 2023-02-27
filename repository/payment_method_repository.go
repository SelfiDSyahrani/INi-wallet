package repository

import (
	"INi-Wallet/model"
	"INi-Wallet/utils"

	"github.com/jmoiron/sqlx"
	//"gorm.io/gorm"
	//"gorm.io/gorm/utils"
)

type PaymentMethodRepository interface {
	Insert(newPaymentMethod model.PaymentMethod) error
	Delete(PaymentMethod_ID string) error
	GetByID(PaymentMethod_ID string) (model.PaymentMethod, error)
	GetAll() ([]model.PaymentMethod, error)
	//Update(paymentMethod model.PaymentMethod) error
}

type paymentMethodRepository struct {
	db *sqlx.DB
}

// Creates a new payment method
func (pmr *paymentMethodRepository) Insert(newPaymentMethod model.PaymentMethod) error {
	_, err := pmr.db.NamedExec(utils.INSERT_PAYMENT_METHOD, newPaymentMethod)
	if err != nil {
		return err
	}
	return nil
}

// Delete
func (pmr *paymentMethodRepository) Delete(PaymentMethod_ID string) error {
	_, err := pmr.db.Exec(utils.DELETE_PAYMENT_METHOD, PaymentMethod_ID)
	if err != nil {
		return err
	}
	return nil
}

//GetByID
func (pmr *paymentMethodRepository) GetByID(PaymentMethod_ID string) (model.PaymentMethod, error) {
	var paymentMethod model.PaymentMethod
	err := pmr.db.QueryRow(utils.SELECT_PAYMENT_METHOD_ID, PaymentMethod_ID).Scan(
		&paymentMethod.Payment_Method_ID,
		&paymentMethod.Method_name,
	)
	if err != nil {
		return model.PaymentMethod{}, err
	}
	return paymentMethod, nil
}

//GetAll
func (pmr *paymentMethodRepository) GetAll() ([]model.PaymentMethod, error) {
	var paymentMethods []model.PaymentMethod
	err := pmr.db.Select(paymentMethods, utils.SELECT_PAYMENT_METHOD_LIST)
	if err != nil {
		return nil, err
	}
	return paymentMethods, nil
}

// Update
/* func (pmr *paymentMethodRepository) Update(paymentMethod model.PaymentMethod) error {
	_,err := pmr.db.NamedExec(utils.BELUM ADA,paymentMethod)
	if err != nil{
		return err
	}
	return nil
} */

//object
func NewPaymentMethodRepository(db *sqlx.DB) PaymentMethodRepository {
	return &paymentMethodRepository{
		db: db,
	}
}
