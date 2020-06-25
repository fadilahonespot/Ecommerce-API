package transaction

import (
	"ecommerce/model"
)

type TransactionUsecase interface {
	ViewAllBasket()(*[]model.Basket, error)
	ViewBasketByUserId(id int)(*[]model.Basket, error)
	ViewSubBasketByProductId(id int)(*model.SubBasket, error)
	UpdateSubBasketByProductId(id int, subbasket *model.SubBasket)(*model.SubBasket, error)
	UpdateSubBasketById(id int, subBasket *model.SubBasket)(*model.SubBasket, error) 
	ViewBasketById(id int) (*model.Basket, error)
	DeleteBasketById(id int) error
	DeleteSubBasketById(id int) error
	ViewBasketByMitraID(id int) (*model.Basket, error)
	ViewSubBasketByBasketId(id int)(*[]model.SubBasket, error)
	ViewBasketList(basket *[]model.Basket) (*[]model.BasketList, error)
	
	AddUserAddress(address *model.UserAddress)(*model.UserAddress, error)
	UpdateUserAddressById(id int, address *model.UserAddress)(*model.UserAddress, error)
	ViewUserAddressByUserId(id int)(*[]model.UserAddress, error)
	ViewUserAddressById(id int)(*model.UserAddress, error)
	DeleteUserAddresById(id int) error
	ViewTransactionByUserId(id int)(*[]model.Transaction, error)
	ViewTransactionById(id int)(*model.Transaction, error)
	ViewAllTransaction()(*[]model.Transaction, error)
	AddBasketSub(subBasket *model.SubBasket, user *model.User) (*model.SubBasket, error)
	// AddTransaction(inTransaction *model.TransactionInput, user *model.User, product *[]model.Product)(*model.TransactionShow, error)
}