package usecase

import (
	"ecommerce/model"
	"ecommerce/ongkir"
	"ecommerce/product"
	"ecommerce/transaction"
	"ecommerce/user"
	"fmt"
	"strconv"
	"strings"
)

type TransactionUsecaseImpl struct {
	transactionRepo transaction.TransactionRepo
	productRepo     product.ProductRepo
	userRepo        user.UserRepo
	ongkirRepo      ongkir.OngkirRepo
}

func CreateTransactionUsecase(transactionRepo transaction.TransactionRepo, productRepo product.ProductRepo, userRepo user.UserRepo, ongkirRepo ongkir.OngkirRepo) transaction.TransactionUsecase {
	return &TransactionUsecaseImpl{transactionRepo, productRepo, userRepo, ongkirRepo}
}

func (e *TransactionUsecaseImpl) ViewAllBasket() (*[]model.Basket, error) {
	return e.transactionRepo.ViewAllBasket()
}

func (e *TransactionUsecaseImpl) ViewBasketByUserId(id int) (*[]model.Basket, error) {
	return e.transactionRepo.ViewBasketByUserId(id)
}

func (e *TransactionUsecaseImpl) ViewSubBasketByProductId(id int) (*model.SubBasket, error) {
	return e.transactionRepo.ViewSubBasketByProductId(id)
}

func (e *TransactionUsecaseImpl) ViewBasketByMitraID(id int) (*model.Basket, error) {
	return e.transactionRepo.ViewBasketByMitraID(id)
}

func (e *TransactionUsecaseImpl) DeleteBasketById(id int) error {
	return e.transactionRepo.DeleteBasketById(id)
}

func (e *TransactionUsecaseImpl) DeleteSubBasketById(id int) error {
	return e.transactionRepo.DeleteSubBasketById(id)
}

func (e *TransactionUsecaseImpl) UpdateSubBasketByProductId(id int, subbasket *model.SubBasket) (*model.SubBasket, error) {
	return e.transactionRepo.UpdateSubBasketByProductId(id, subbasket)
}

func (e *TransactionUsecaseImpl) ViewBasketById(id int) (*model.Basket, error) {
	return e.transactionRepo.ViewBasketById(id)
}

func (e *TransactionUsecaseImpl) ViewSubBasketByBasketId(id int) (*[]model.SubBasket, error) {
	return e.transactionRepo.ViewSubBasketByBasketId(id)
}

func (e *TransactionUsecaseImpl) ViewBasketList(basket *[]model.Basket) (*[]model.BasketList, error) {
	if len(*basket) == 0 {
		return nil, fmt.Errorf("your basket is empty")
	}
	var basketLists []model.BasketList
	var basketList = model.BasketList{}
	for i := 0; i < len(*basket); i++ {
		subBaskets, err := e.transactionRepo.ViewSubBasketByBasketId(int((*basket)[i].ID))
		if err != nil {
			return nil, err
		}
		mitra, err := e.userRepo.ViewMitraById(int((*basket)[i].MitraID))
		if err != nil {
			return nil, err
		}
		basketList.BasketID = (*basket)[i].ID
		basketList.MitraID = (*basket)[i].MitraID
		basketList.MitraName = mitra.StoreName
		var subBasketLists []model.BasketSubList
		var subBasketList = model.BasketSubList{}
		for j := 0; j < len(*subBaskets); j++ {
			product, err := e.productRepo.ViewProductById(int((*subBaskets)[j].ProductID))
			if err != nil {
				return nil, err
			}
			images, err := e.productRepo.ImageByProductId(int(product.ID))
			if err != nil {
				return nil, err
			}
			subBasketList = model.BasketSubList{
				SubBasketID: (*subBaskets)[j].ID,
				ProductID:   (*subBaskets)[j].ProductID,
				ProductName: product.Name,
				Prince:      product.Prince,
				Quantinty:   (*subBaskets)[j].Quantity,
				SubTotal:    (*subBaskets)[j].SubTotal,
				Stock:       product.Stock,
				Image:       (*images)[0].PartImg,
			}
			subBasketLists = append(subBasketLists, subBasketList)
			basketList.Total += subBasketList.SubTotal
		}
		basketList.ProductsList = subBasketLists
		basketLists = append(basketLists, basketList)
		basketList.Total = 0
	}
	return &basketLists, nil
}

func (e *TransactionUsecaseImpl) AddUserAddress(address *model.UserAddress) (*model.UserAddress, error) {
	tx := e.ongkirRepo.BeginTrans()
	city, err := e.ongkirRepo.GetCityById(address.CityID)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(address.MobileNumber, "0") == true {
		address.MobileNumber = strings.Replace(address.MobileNumber, "0", "62", 1)
	}
	if strings.HasPrefix(address.MobileNumber, "62") == false {
		address.MobileNumber = "62" + address.MobileNumber
	}
	address.CityName = city.Type + " " + city.CityName
	address.Province = city.Province
	address.PostalCode = city.PostalCode
	inAddress, err := e.transactionRepo.AddUserAddress(address, tx)
	if err != nil {
		return nil, err
	}
	tx.Commit()
	return inAddress, nil
}

func (e *TransactionUsecaseImpl) UpdateUserAddressById(id int, address *model.UserAddress) (*model.UserAddress, error) {
	tx := e.ongkirRepo.BeginTrans()
	city, err := e.ongkirRepo.GetCityById(address.CityID)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(address.MobileNumber, "0") == true {
		address.MobileNumber = strings.Replace(address.MobileNumber, "0", "62", 1)
	}
	if strings.HasPrefix(address.MobileNumber, "62") == false {
		address.MobileNumber = "62" + address.MobileNumber
	}
	address.CityName = city.Type + " " + city.CityName
	address.Province = city.Province
	address.PostalCode = city.PostalCode
	upAddress, err := e.transactionRepo.UpdateUserAddressById(id, address, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	allAddress, err := e.transactionRepo.ViewUserAddressByUserId(int(upAddress.UserID))
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if (*allAddress)[0].ID == upAddress.ID {
		mobileNumber, err := strconv.Atoi(upAddress.MobileNumber)
		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("Mobile number has be number")
		}
		var user = model.User{
			Nama:     upAddress.ReceipedName,
			NoHP:     mobileNumber,
			IDKota:   upAddress.CityID,
			Kota:     upAddress.CityName,
			Provinsi: upAddress.Province,
			KodePos:  upAddress.PostalCode,
			Alamat:   upAddress.Address,
		}
		_, err = e.userRepo.UpdateUser(int(upAddress.UserID), &user, tx)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	tx.Commit()
	return upAddress, nil
}

func (e *TransactionUsecaseImpl) ViewUserAddressByUserId(id int) (*[]model.UserAddress, error) {
	return e.transactionRepo.ViewUserAddressByUserId(id)
}

func (e *TransactionUsecaseImpl) ViewUserAddressById(id int) (*model.UserAddress, error) {
	return e.transactionRepo.ViewUserAddressById(id)
}

func (e *TransactionUsecaseImpl) DeleteUserAddresById(id int) error {
	return e.transactionRepo.DeleteUserAddresById(id)
}

func (e *TransactionUsecaseImpl) ViewTransactionByUserId(id int) (*[]model.Transaction, error) {
	return e.transactionRepo.ViewTransactionByUserId(id)
}

func (e *TransactionUsecaseImpl) ViewTransactionById(id int) (*model.Transaction, error) {
	return e.transactionRepo.ViewTransactionById(id)
}

func (e *TransactionUsecaseImpl) ViewAllTransaction() (*[]model.Transaction, error) {
	return e.transactionRepo.ViewAllTransaction()
}

func (e *TransactionUsecaseImpl) UpdateSubBasketById(id int, subBasket *model.SubBasket)(*model.SubBasket, error) {
	return e.transactionRepo.UpdateSubBasketById(id, subBasket)
}

func (e *TransactionUsecaseImpl) AddBasketSub(subBasket *model.SubBasket, user *model.User) (*model.SubBasket, error) {
	tx := e.ongkirRepo.BeginTrans()
	mitra, err := e.userRepo.ViewMitraByUserId(int(user.ID))
	if err == nil {
		products, err := e.productRepo.ProductByMitraId(int(mitra.ID))
		if err == nil {
			for i := 0; i < len(*products); i++ {
				if (*products)[i].ID == subBasket.ProductID {
					return nil, fmt.Errorf("can't add products from the store itself")
				}
			}
		}
	}
	product, err := e.productRepo.ViewProductById(int(subBasket.ProductID))
	if err != nil {
		return nil, err
	}
	baskets, err := e.transactionRepo.ViewBasketByUserId(int(user.ID))
	if err != nil {
		return nil, err
	}
	if len(*baskets) != 0 {
		for i := 0; i < len(*baskets); i++ {
			if (*baskets)[i].MitraID == product.MitraID {
				subBaskets, err := e.transactionRepo.ViewSubBasketByBasketId(int((*baskets)[i].ID))
				if err != nil {
					return nil, err
				}
				for k := 0; k < len(*subBaskets); k++ {
					if (*subBaskets)[k].ProductID == subBasket.ProductID {
						subBasket.Quantity += (*subBaskets)[k].Quantity
						subBasket.SubTotal = subBasket.Quantity * product.Prince
						upBasket, err := e.transactionRepo.UpdateSubBasketByProductId(int(subBasket.ProductID), subBasket)
						if err != nil {
							return nil, err
						}
						return upBasket, nil
					}
				}
				var subBasket = model.SubBasket{
					BasketID:  (*baskets)[i].ID,
					ProductID: product.ID,
					Quantity:  subBasket.Quantity,
					SubTotal:  subBasket.Quantity * product.Prince,
				}
				inSubbasket, err := e.transactionRepo.InsertSubBasket(&subBasket, tx)
				if err != nil {
					return nil, err
				}
				tx.Commit()
				return inSubbasket, nil
			}
		}
	}

	var basket = model.Basket{
		UserID:  user.ID,
		MitraID: product.MitraID,
	}
	inBasket, err := e.transactionRepo.InsertBasket(&basket, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	var subbasket = model.SubBasket{
		BasketID:  inBasket.ID,
		ProductID: product.ID,
		Quantity:  subBasket.Quantity,
		SubTotal:  subBasket.Quantity * product.Prince,
	}
	inSubBasket, err := e.transactionRepo.InsertSubBasket(&subbasket, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return inSubBasket, nil
}

