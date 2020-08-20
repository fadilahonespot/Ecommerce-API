package transaction

import (
	"ecommerce/model"
	"github.com/jinzhu/gorm"
)

type TransactionRepo interface {
	ViewAllBasket()(*[]model.Basket, error)
	ViewBasketByUserId(id int)(*[]model.Basket, error)
	DeleteBasketById(id int) error
	ViewBasketByMitraID(id int) (*model.Basket, error)
	ViewSubBasketByBasketId(id int)(*[]model.SubBasket, error)
	InsertBasket(basket *model.Basket, tx *gorm.DB)(*model.Basket, error)
	InsertSubBasket(subbasket *model.SubBasket, tx *gorm.DB)(*model.SubBasket, error)
	ViewSubBasketByProductId(id int)(*model.SubBasket, error)
	UpdateSubBasketByProductId(id int, subbasket *model.SubBasket)(*model.SubBasket, error)
	UpdateSubBasketById(id int, subBasket *model.SubBasket)(*model.SubBasket, error) 
	ViewBasketById(id int) (*model.Basket, error)
	DeleteSubBasketByProductId(id int, tx *gorm.DB) error
	DeleteSubBasketById(id int) error
	AddUserAddress(address *model.UserAddress, tx *gorm.DB)(*model.UserAddress, error)
	ViewUserAddressById(id int)(*model.UserAddress, error)
	UpdateUserAddressById(id int, address *model.UserAddress, tx *gorm.DB)(*model.UserAddress, error)
	ViewUserAddressByUserId(id int)(*[]model.UserAddress, error)
	DeleteUserAddresById(id int) error
	ViewTransactionByUserId(id int)(*[]model.Transaction, error)
	ViewTransactionById(id int)(*model.Transaction, error)
	InsertTransaction(transaction *model.Transaction, tx *gorm.DB)(*model.Transaction, error)
	UpdateTransaction(id int, transaction*model.Transaction, tx *gorm.DB)(*model.Transaction, error)
	ViewAllTransaction()(*[]model.Transaction, error)
	ViewMitraTransactionById(id int)(*model.MitraTransaction, error)
	InsertMitraTransaction(transaction *model.MitraTransaction, tx *gorm.DB) (*model.MitraTransaction, error)
	UpdateMitraTransaction(id int, transaction *model.MitraTransaction, tx *gorm.DB) (*model.MitraTransaction, error)
	ViewSubTransactionById(id int)(*model.SubTransaction, error)
	InsertSubTransaction(subtransaction *model.SubTransaction, tx *gorm.DB)(*model.SubTransaction, error)
	UpdateSubTransaction(id int, subtransaction *model.SubTransaction, tx *gorm.DB) (*model.SubTransaction, error)
}