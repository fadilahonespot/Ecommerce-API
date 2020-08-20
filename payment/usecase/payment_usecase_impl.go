package usecase

import (
	"ecommerce/model"
	"ecommerce/payment"
)

type PaymentUsecaseImpl struct {
	paymentRepo payment.PaymentRepo
}

func CreatePaymentUsecase(paymentRepo payment.PaymentRepo) payment.PaymentUsecase {
	return &PaymentUsecaseImpl{paymentRepo}
}

func (e *PaymentUsecaseImpl) ViewAllPayment() (*[]model.Payment, error) {
	return e.paymentRepo.ViewAllPayment()
}

func (e *PaymentUsecaseImpl) ViewPaymentById(id int) (*model.Payment, error) {
	return e.paymentRepo.ViewPaymentById(id)
}

func (e *PaymentUsecaseImpl) InsertPaymentMethod(payment *model.PaymentMethod)(*model.PaymentMethod, error) {
	return e.paymentRepo.InsertPaymentMethod(payment)
}

func (e *PaymentUsecaseImpl) UpdatePaymentMethodById(id int, payment *model.PaymentMethod)(*model.PaymentMethod, error) {
	return e.paymentRepo.UpdatePaymentMethodById(id, payment)
}

func (e *PaymentUsecaseImpl) ViewPaymentMethodById(id int) (*model.PaymentMethod, error) {
	return e.paymentRepo.ViewPaymentMethodById(id)
}

func (e *PaymentUsecaseImpl) ViewAllPaymentMethod()(*[]model.PaymentMethod, error) {
	return e.paymentRepo.ViewAllPaymentMethod()
}

func (e *PaymentUsecaseImpl) DeletePaymentMethodById(id int) error {
	return e.paymentRepo.DeletePaymentMethodById(id)
}

func (e *PaymentUsecaseImpl) ViewPaymentByPaymentMethodId(id int)(*[]model.Payment, error) {
	return e.paymentRepo.ViewPaymentByPaymentMethodId(id)
}

func (e *PaymentUsecaseImpl) DeletePaymentById(id int) error {
	return e.paymentRepo.DeletePaymentById(id)
}

func (e *PaymentUsecaseImpl) InsertPayment(payment *model.Payment) (*model.Payment, error) {
	return e.paymentRepo.InsertPayment(payment)
}

func (e *PaymentUsecaseImpl) UpdatePayment(id int, payment *model.Payment) (*model.Payment, error) {
	return e.paymentRepo.UpdatePayment(id, payment)
}
