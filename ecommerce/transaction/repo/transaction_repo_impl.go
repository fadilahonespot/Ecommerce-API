package repo

import (
	"ecommerce/model"
	"ecommerce/transaction"
	"fmt"

	"github.com/jinzhu/gorm"
)

type TransacationRepoImpl struct {
	DB *gorm.DB
}

func CreateTransactionRepo(DB *gorm.DB) transaction.TransactionRepo {
	return &TransacationRepoImpl{DB}
}

func (e *TransacationRepoImpl) ViewAllBasket()(*[]model.Basket, error) {
	var basket []model.Basket
	err := e.DB.Find(&basket).Error
	if err != nil {
		fmt.Printf("[TransactionRepoImpl.ViewAllBasket] error execute query %v \n", err)
		return nil, fmt.Errorf("failed view all data basket")
	}
	return &basket, nil
}

func (e *TransacationRepoImpl) ViewBasketByUserId(id int)(*[]model.Basket, error) {
	var basket []model.Basket
	err := e.DB.Table("basket").Where("user_id = ?", id).Find(&basket).Error
	if err != nil {
		fmt.Printf("[TransactionRepoImpl.ViewBasketByUserId] error execute query %v \n", err)
		return nil, fmt.Errorf("id user not found in basket")
	}
	return &basket, nil
}

func (e *TransacationRepoImpl) ViewSubBasketByProductId(id int)(*model.SubBasket, error) {
	var basket = model.SubBasket{}
	err := e.DB.Table("sub_basket").Where("product_id = ?", id).First(&basket).Error
	if err != nil {
		fmt.Printf("[TransactionRepoImpl.ViewBasketByProductId] error execute query %v \n", err)
		return nil, fmt.Errorf("id product not found in basket")
	}
	return &basket, nil
}

func (e *TransacationRepoImpl) InsertBasket(basket *model.Basket, tx *gorm.DB)(*model.Basket, error) {
	err := tx.Save(&basket).Error
	if err != nil {
		fmt.Printf("[TransactionRepoImpl.InsertBasket] error execute query %v \n", err)
		return nil, fmt.Errorf("failed insert data basket")
	}
	return basket, nil
}

func (e *TransacationRepoImpl) UpdateSubBasketByProductId(id int, subbasket *model.SubBasket)(*model.SubBasket, error) {
	var upBasket = model.SubBasket{}
	err := e.DB.Table("sub_basket").Where("product_id = ?", id).First(&upBasket).Update(&subbasket).Error
	if err != nil {
		fmt.Printf("[TransactionRepoImpl.Updatebasket] error execute query %v \n", err)
		return nil, fmt.Errorf("failed update data basket")
	}
	return &upBasket, nil
}

func (e *TransacationRepoImpl) UpdateSubBasketById(id int, subBasket *model.SubBasket)(*model.SubBasket, error) {
	var upSubbasket = model.SubBasket{}
	err := e.DB.Table("sub_basket").Where("id = ?", id).First(&upSubbasket).Update(&subBasket).Error
	if err != nil {
		fmt.Printf("[TransactionRepoImpl.UpdateSubBasketById] error execute query %v \n", err)
		return nil, fmt.Errorf("failed update sub basket, id is not exsis")
	}
	return &upSubbasket, nil
}

func (e *TransacationRepoImpl) ViewBasketByMitraID(id int) (*model.Basket, error) {
	var basket = model.Basket{}
	err := e.DB.Table("basket").Where("mitra_id = ?", id).First(&basket).Error
	if err != nil {
		fmt.Printf("[TransactionRepoImpl.ViewBasketByMitraId] error execute query %v \n", err)
		return nil, fmt.Errorf("id mitra is not exsis in basket")
	}
	return &basket, nil
}

func (e *TransacationRepoImpl) ViewSubBasketByBasketId(id int)(*[]model.SubBasket, error) {
	var subbasket []model.SubBasket
	err := e.DB.Table("sub_basket").Where("basket_id = ?", id).Find(&subbasket).Error
	if err != nil {
		fmt.Printf("[TransactionRepoImpl.ViewSubBasketByBasketId] error execute query %v \n", err)
		return nil, fmt.Errorf("id basket is not exsis")
	}
	return &subbasket, nil
}

func (e *TransacationRepoImpl) InsertSubBasket(subbasket *model.SubBasket, tx *gorm.DB)(*model.SubBasket, error) {
	err := tx.Save(&subbasket).Error
	if err != nil {
		fmt.Printf("[TransactionRepoImpl.InsertSubBasket] error execute query %v \n", err)
		return nil, fmt.Errorf("failed add data to sub basket")
	}
	return subbasket, nil
}

func (e *TransacationRepoImpl) DeleteBasketById(id int) error {
	var basket = model.Basket{}
	err := e.DB.Table("basket").Where("id = ?", id).Delete(&basket).Error
	if err != nil {
		fmt.Printf("[TransaktionRepoImpl.DeleteBasketById] error execute query %v \n", err)
		return fmt.Errorf("id is not exist")
	}
	return nil
}

func (e *TransacationRepoImpl) DeleteSubBasketByProductId(id int, tx *gorm.DB) error {
	var basket = model.SubBasket{}
	err := tx.Table("sub_basket").Where("product_id = ?", id).Delete(&basket).Error
	if err != nil {
		fmt.Printf("[TransactionRepoImpl.DeleteBasketByProductId] error execute query %v \n", err)
		return fmt.Errorf("id product in basket not exist")
	}
	return nil
}

func (e *TransacationRepoImpl) DeleteSubBasketById(id int) error {
	var subbasket = model.SubBasket{}
	err := e.DB.Table("sub_basket").Where("id = ?", id).Delete(&subbasket).Error
	if err != nil {
		fmt.Printf("[TransactionRepoImpl.DeleteSubBasketById] error execute query %v \n", err)
		return fmt.Errorf("failed delete data, id basket is not exsis")
	}
	return nil
}

func (e *TransacationRepoImpl) ViewBasketById(id int) (*model.Basket, error) {
	var basket = model.Basket{}
	err := e.DB.Table("basket").Where("id = ?", id).First(&basket).Error
	if err != nil {
		fmt.Printf("[TransactionRepoImpl.ViewBasketById] error execute query %v \n", err)
		return nil, fmt.Errorf("id basket not exist")
	}
	return &basket, nil
}

func (e *TransacationRepoImpl) AddUserAddress(address *model.UserAddress, tx *gorm.DB)(*model.UserAddress, error) {
	err := tx.Save(&address).Error
	if err != nil {
		fmt.Printf("[TransactionRepoImpl.AddUserAddress] error execute query %v \n", err)
		return nil, fmt.Errorf("failed insert data user addres")
	}
	return address, nil
}

func (e *TransacationRepoImpl) UpdateUserAddressById(id int, address *model.UserAddress, tx *gorm.DB)(*model.UserAddress, error) {
	var upAddress = model.UserAddress{}
	err := tx.Table("user_address").Where("id = ?", id).First(&upAddress).Update(&address).Error
	if err != nil {
		fmt.Printf("[TransactionRepoImpl.UpdateUserAddressByUserId] error execute query %v \n", err)
		return nil, fmt.Errorf("failed update data user address")
	}
	return &upAddress, nil
}

func (e *TransacationRepoImpl) ViewUserAddressById(id int)(*model.UserAddress, error) {
	var address = model.UserAddress{}
	err := e.DB.Table("user_address").Where("id = ?", id).First(&address).Error
	if err != nil {
		fmt.Print("[TransactionRepoImpl.ViewUseraddressById] error execute query %v \n", err)
		return nil, fmt.Errorf("id user address not exsis")
	}
	return &address, nil
}

func (e *TransacationRepoImpl) ViewUserAddressByUserId(id int)(*[]model.UserAddress, error) {
	var address []model.UserAddress
	err := e.DB.Table("user_address").Where("user_id = ?", id).Find(&address).Error
	if err != nil {
		fmt.Printf("[TransactionRepo.ViewUserAddressByUserId] error execute query %v \n", err)
		return nil, fmt.Errorf("failed view user address, user id is not exsis")
	}
	return &address, nil
}

func (e *TransacationRepoImpl) DeleteUserAddresById(id int) error {
	var address = model.UserAddress{}
	err := e.DB.Table("user_address").Where("id = ?", id).Delete(&address).Error
	if err != nil {
		fmt.Printf("[TransactionRepoImpl.DeleteUserAddresById] error execute query %v \n", err)
		return fmt.Errorf("failed delete data, id is not exist")
	}
	return nil
}

func (e *TransacationRepoImpl) ViewAllTransaction()(*[]model.Transaction, error) {
	var transaction []model.Transaction
	err := e.DB.Find(&transaction).Error
	if err != nil {
		fmt.Printf("[TransactionRepoImpl.ViewAllTransaction] error execute query %v \n", err)
		return nil, fmt.Errorf("failed view all data transaction")
	}
	return &transaction, nil
}

func (e *TransacationRepoImpl) ViewTransactionByUserId(id int)(*[]model.Transaction, error) {
	var transaction []model.Transaction
	err := e.DB.Table("transaction").Where("user_id = ?", id).Find(&transaction).Error
	if err != nil {
		fmt.Printf("[TransactionRepoImpl.ViewTransactionbyUserId] error execute query %v \n", err)
		return nil, fmt.Errorf("failed view data transaction, user id is not exist")
	}
	return &transaction, nil
}

func (e *TransacationRepoImpl) ViewTransactionById(id int)(*model.Transaction, error) {
	var transaction = model.Transaction{}
	err := e.DB.Table("transaction").Where("id = ?", id).First(&transaction).Error
	if err != nil {
		fmt.Printf("[TransactionRepoImpl.ViewTransactionById] error execute query %v \n", err)
		return nil, fmt.Errorf("failed view data transaction, id is not exist")
	}
	return &transaction, nil
}

func (e *TransacationRepoImpl) InsertTransaction(transaction *model.Transaction, tx *gorm.DB)(*model.Transaction, error) {
	err := tx.Save(&transaction).Error
	if err != nil {
		fmt.Printf("[TransactionRepoImpl.InsertTransaction] error execute query %v \n", err)
		return nil, fmt.Errorf("failed insert data transaction")
	}
	return transaction, nil
}

func (e *TransacationRepoImpl) UpdateTransaction(id int, transaction*model.Transaction, tx *gorm.DB)(*model.Transaction, error) {
	upTransaction := model.Transaction{}
	err := tx.Table("transaction").Where("id = ?", id).First(&upTransaction).Update(&transaction).Error
	if err != nil {
		fmt.Printf("[TransationRepoImpl.UpdateTransaction] error execute query %v \n", err)
		return nil, fmt.Errorf("failed update transaction, id is not exsis")
	}
	return &upTransaction, nil
}

func (e *TransacationRepoImpl) ViewMitraTransactionById(id int)(*model.MitraTransaction, error) {
	var mitraTransaction = model.MitraTransaction{}
	err := e.DB.Table("mitra_transaction").Where("id = ?", id).First(&mitraTransaction).Error
	if err != nil {
		fmt.Printf("[TransactionRepoImpl.ViewMitraTransactionById] error execute query %v \n", err)
		return nil, fmt.Errorf("id mitra transaction is not exsis")
	} 
	return &mitraTransaction, nil
}

func (e *TransacationRepoImpl) InsertMitraTransaction(transaction *model.MitraTransaction, tx *gorm.DB) (*model.MitraTransaction, error) {
	err := tx.Save(&transaction).Error
	if err != nil {
		fmt.Printf("[TransactionRepoImpl.InsertMitraTransaction] error execute query %v \n", err)
		return nil, fmt.Errorf("failed insert data in mitra transaction")
	}
	return transaction, nil
}

func (e *TransacationRepoImpl) UpdateMitraTransaction(id int, transaction *model.MitraTransaction, tx *gorm.DB) (*model.MitraTransaction, error) {
	var upTransaction = model.MitraTransaction{}
	err := tx.Table("mitra_transaction").Where("id = ?", id).First(&upTransaction).Update(&transaction).Error
	if err != nil {
		fmt.Printf("[TransactionRepoImpl.UpdateMitraTransaction] error execute query %v \n", err)
		return nil, fmt.Errorf("failed update data mitra transaction")
	}
	return &upTransaction, nil
}

func (e *TransacationRepoImpl) ViewSubTransactionById(id int)(*model.SubTransaction, error) {
	var subtransaction = model.SubTransaction{}
	err := e.DB.Table("subtransaction").Where("id = ?", id).First(&subtransaction).Error
	if err != nil {
		fmt.Printf("[TransactionRepoImpl.ViewSubTransactionById] error execute query %v \n", err)
		return nil, fmt.Errorf("id sub transaction is not exsis")
	}
	return &subtransaction, nil
}

func (e *TransacationRepoImpl) InsertSubTransaction(subtransaction *model.SubTransaction, tx *gorm.DB)(*model.SubTransaction, error) {
	err := tx.Save(&subtransaction).Error
	if err != nil {
		fmt.Printf("[TransactionRepoImpl.InsertSubTransaction] error execute query %v \n", err)
		return nil, fmt.Errorf("failed insert data sub transaction")
	}
	return subtransaction, nil
}

func (e *TransacationRepoImpl) UpdateSubTransaction(id int, subtransaction *model.SubTransaction, tx *gorm.DB) (*model.SubTransaction, error) {
	var upSubtransaction = model.SubTransaction{}
	err := tx.Table("subtransaction").Where("id = ?", id).First(&upSubtransaction).Update(subtransaction).Error
	if err != nil {
		fmt.Printf("[TransactionRepoImpl.UpdateSubtransaction] error execute query %v \n", err)
		return nil, fmt.Errorf("failed update subtransaction")
	}
	return &upSubtransaction, nil
}

