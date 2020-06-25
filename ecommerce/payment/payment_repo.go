package payment

import "ecommerce/model"

type PaymentRepo interface {
	InsertPayment(payment *model.Payment) (*model.Payment, error)
	ViewAllPayment()(*[]model.Payment, error)
	UpdatePayment(id int, payment *model.Payment) (*model.Payment, error)
	ViewPaymentById(id int)(*model.Payment, error)
	InsertPaymentMethod(payment *model.PaymentMethod)(*model.PaymentMethod, error)
	UpdatePaymentMethodById(id int, payment *model.PaymentMethod)(*model.PaymentMethod, error)
	ViewPaymentMethodById(id int) (*model.PaymentMethod, error)
	DeletePaymentMethodById(id int) error
	ViewAllPaymentMethod()(*[]model.PaymentMethod, error)
	DeletePaymentById(id int) error
	ViewPaymentByPaymentMethodId(id int)(*[]model.Payment, error)
}