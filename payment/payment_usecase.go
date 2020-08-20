package payment

import "ecommerce/model"

type PaymentUsecase interface {
	InsertPayment(payment *model.Payment) (*model.Payment, error)
	UpdatePayment(id int, payment *model.Payment) (*model.Payment, error)
	ViewAllPayment()(*[]model.Payment, error)
	ViewPaymentById(id int)(*model.Payment, error)
	DeletePaymentMethodById(id int) error
	ViewPaymentByPaymentMethodId(id int)(*[]model.Payment, error)
	ViewPaymentMethodById(id int) (*model.PaymentMethod, error)
	DeletePaymentById(id int) error
	ViewAllPaymentMethod()(*[]model.PaymentMethod, error)
	InsertPaymentMethod(payment *model.PaymentMethod)(*model.PaymentMethod, error)
	UpdatePaymentMethodById(id int, payment *model.PaymentMethod)(*model.PaymentMethod, error)
}