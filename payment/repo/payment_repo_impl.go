package repo

import (
	"ecommerce/model"
	"ecommerce/payment"
	"fmt"

	"github.com/jinzhu/gorm"
)

type PaymentRepoImpl struct {
	DB *gorm.DB
}

func CreatePaymetRepo(DB *gorm.DB) payment.PaymentRepo {
	return &PaymentRepoImpl{DB}
}

func (e *PaymentRepoImpl) InsertPayment(payment *model.Payment) (*model.Payment, error) {
	err := e.DB.Save(&payment).Error
	if err != nil {
		fmt.Printf("[PaymentRepoImpl.InsertPayment] error execute query %v \n", err)
		return nil, fmt.Errorf("failed inser data payment")
	}
	return payment, nil
}

func (e *PaymentRepoImpl) UpdatePayment(id int, payment *model.Payment) (*model.Payment, error) {
	var upPayment = model.Payment{}
	err := e.DB.Table("payment").Where("id = ?", id).First(&upPayment).Update(&payment).Error
	if err != nil {
		fmt.Printf("[PaymentRepoImpl.UpdatePayment] error execute query %v \n", err)
		return nil, fmt.Errorf("failed update payment, id payment not exist")
	}
	return &upPayment, nil
}

func (e *PaymentRepoImpl) ViewPaymentById(id int)(*model.Payment, error) {
	var payment = model.Payment{}
	err := e.DB.Table("payment").Where("id = ?", id).First(&payment).Error
	if err != nil {
		fmt.Printf("[PaymentRepoImpl.ViewpaymentById] error execute query %v \n", err)
		return nil, fmt.Errorf("id paymet is not exist")
	}
	return &payment, nil
}

func (e *PaymentRepoImpl) ViewAllPayment()(*[]model.Payment, error) {
	var payment []model.Payment
	err := e.DB.Table("payment").Find(&payment).Error
	if err != nil {
		fmt.Printf("[PaymentRepoImpl.ViewAllPayment] error execute query %v \n", err)
		return nil, fmt.Errorf("failed view all data payment")
	}
	return &payment, nil
}

func (e *PaymentRepoImpl) InsertPaymentMethod(payment *model.PaymentMethod)(*model.PaymentMethod, error) {
	err := e.DB.Save(&payment).Error
	if err != nil {
		fmt.Printf("[PaymentRepoImpl.InsertPaymentMethods] error execute query %v \n", err)
		return nil, fmt.Errorf("failed insert data payment methods")
	}
	return payment, nil
}

func (e *PaymentRepoImpl) UpdatePaymentMethodById(id int, payment *model.PaymentMethod)(*model.PaymentMethod, error) {
	var upPayment = model.PaymentMethod{}
	err := e.DB.Table("payment_method").Where("id = ?", id).First(&upPayment).Update(&payment).Error
	if err != nil {
		fmt.Printf("[PaymentRepoImpl.UpdatePaymentMethodsById] error execute query %v \n", err)
		return nil, fmt.Errorf("failed update payment methods %v \n", err)
	}
	return &upPayment, nil
}

func (e *PaymentRepoImpl) ViewPaymentMethodById(id int) (*model.PaymentMethod, error) {
	var payment = model.PaymentMethod{}
	err := e.DB.Table("payment_method").Where("id = ?", id).First(&payment).Error
	if err != nil {
		fmt.Printf("[PaymentRepoImpl.ViewpaymentMethodById] error execute query %v \n", err)
		return nil, fmt.Errorf("id payment methods is not exsis")
	}
	return &payment, nil
}

func (e *PaymentRepoImpl) ViewAllPaymentMethod()(*[]model.PaymentMethod, error) {
	var payment []model.PaymentMethod
	err := e.DB.Find(&payment).Error
	if err != nil {
		fmt.Printf("[PaymentRepoImpl.viewPaymentmethod] error execute query %v \n", err)
		return nil, fmt.Errorf("failed show data payment methods")
	}
	return &payment, nil
}

func (e *PaymentRepoImpl) DeletePaymentMethodById(id int) error {
	var payment = model.PaymentMethod{}
	err := e.DB.Table("payment_method").Where("id = ?", id).Delete(&payment).Error
	if err != nil {
		fmt.Printf("[PaymentRepoImpl.DeletePaymentMethodById] error execute query %v \n", err)
		return fmt.Errorf("id payment method is not exsis")
	}
	return nil
}

func (e *PaymentRepoImpl) ViewPaymentByPaymentMethodId(id int)(*[]model.Payment, error) {
	var payments []model.Payment
	err := e.DB.Table("payment").Where("payment_method_id = ?", id).Find(&payments).Error
	if err != nil {
		fmt.Printf("[PaymentRepoImpl.ViewPaymentByIdPaymentMethods] error execute query %v \n", err)
		return nil, fmt.Errorf("failed show data payment")
	}
	return &payments, nil
}

func (e *PaymentRepoImpl) DeletePaymentById(id int) error {
	var payment = model.Payment{}
	err := e.DB.Table("payment").Where("id = ?", id).Delete(&payment).Error
	if err != nil {
		fmt.Printf("[PaymentRepoImpl] error execute query %v \n", err)
		return fmt.Errorf("failed delete payment")
	}
	return nil
}



