package handler

import (
	"ecommerce/middleware"
	"ecommerce/model"
	"ecommerce/ongkir"
	"ecommerce/product"
	"ecommerce/transaction"
	"ecommerce/user"
	"ecommerce/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	transactionUsecase transaction.TransactionUsecase
	userUsecase        user.UserUsecase
	productUsecase     product.ProductUsecase
	ongkirUsecase      ongkir.OngkirUsecase
}

func CreateTransactionHandler(r *gin.Engine, transactionUsecase transaction.TransactionUsecase, userUsecase user.UserUsecase, productUsecase product.ProductUsecase, ongkirUsecase ongkir.OngkirUsecase) {
	transactionHandler := TransactionHandler{transactionUsecase, userUsecase, productUsecase, ongkirUsecase}

	r.POST("/shipping", transactionHandler.checkShippingStore)

	r2 := r.Group("/user").Use(middleware.TokenVerifikasiMiddleware())
	r2.POST("/basket", transactionHandler.addBasket)
	r2.PUT("/basket", transactionHandler.editSubBasket)
	r2.GET("/basket", transactionHandler.viewBasket)
	r2.DELETE("/basket/:idSubBasket", transactionHandler.deleteSubBasket)
	r2.GET("/address", transactionHandler.viewMyAddress)
	r2.POST("/address", transactionHandler.addUserAddress)
	r2.PUT("/address/:idAddress", transactionHandler.editeUserAddress)
	r2.DELETE("/address/:idAddress", transactionHandler.removeMyAddress)
}

func (e *TransactionHandler) addBasket(c *gin.Context) {
	user, err := e.userUsecase.ValidUser(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	var subbasket = model.SubBasket{}
	err = c.Bind(&subbasket)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, "Oopss server someting wrong")
		return
	}
	if subbasket.ID != 0 || subbasket.SubTotal != 0 {
		utils.HandleError(c, http.StatusBadRequest, "inputs are not permitted")
		return
	}
	inBasket, err := e.transactionUsecase.AddBasketSub(&subbasket, user)
	if err != nil {
		utils.HandleError(c, http.StatusForbidden, err.Error())
		return
	}
	utils.HandleSucces(c, inBasket)
}

func (e *TransactionHandler) editSubBasket(c *gin.Context) {
	user, err := e.userUsecase.ValidUser(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	var inBasket = model.SubBasket{}
	err = c.Bind(&inBasket)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, "Oopss server someting wrong")
		return
	}
	if inBasket.ProductID != 0 || inBasket.SubTotal != 0 {
		utils.HandleError(c, http.StatusBadRequest, "inputs are not permitted")
		return
	}
	baskets, err := e.transactionUsecase.ViewBasketByUserId(int(user.ID))
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	if len(*baskets) == 0 {
		utils.HandleError(c, http.StatusNotFound, "you do not have a basket list")
		return
	}
	for i := 0; i < len(*baskets); i++ {
		subBaskets, err := e.transactionUsecase.ViewSubBasketByBasketId(int((*baskets)[i].ID))
		if err != nil {
			utils.HandleError(c, http.StatusNotFound, err.Error())
			return
		}
		for k := 0; k < len(*subBaskets); k++ {
			if (*subBaskets)[k].ID == inBasket.ID {
				product, err := e.productUsecase.ViewProductById(int((*subBaskets)[k].ProductID))
				if err != nil {
					utils.HandleError(c, http.StatusNotFound, err.Error())
					return
				}
				var subBasket = model.SubBasket{
					Quantity: inBasket.Quantity,
					SubTotal: inBasket.Quantity * product.Prince,
				}
				upSubbasket, err := e.transactionUsecase.UpdateSubBasketById(int(inBasket.ID), &subBasket)
				if err != nil {
					utils.HandleError(c, http.StatusNotFound, err.Error())
					return
				}
				utils.HandleSucces(c, upSubbasket)
				return
			}
		}
	}
	utils.HandleError(c, http.StatusBadRequest, "id sub basket is not found")
}

func (e *TransactionHandler) deleteSubBasket(c *gin.Context) {
	user, err := e.userUsecase.ValidUser(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	idStr := c.Param("idSubBasket")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "id has be number")
		return
	}
	baskets, err := e.transactionUsecase.ViewBasketByUserId(int(user.ID))
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
	}
	if len(*baskets) == 0 {
		utils.HandleError(c, http.StatusNotFound, "you don't have a basket list")
		return
	}
	for i := 0; i < len(*baskets); i++ {
		subBaskets, err := e.transactionUsecase.ViewSubBasketByBasketId(int((*baskets)[i].ID))
		if err != nil {
			utils.HandleError(c, http.StatusNotFound, err.Error())
			return
		}
		for k := 0; k < len(*subBaskets); k++ {
			if uint(id) == (*subBaskets)[k].ID {
				err := e.transactionUsecase.DeleteSubBasketById(id)
				if err != nil {
					utils.HandleError(c, http.StatusNotFound, err.Error())
					return
				}
				upBaskets, err := e.transactionUsecase.ViewSubBasketByBasketId(int((*baskets)[i].ID))
				if err != nil {
					utils.HandleError(c, http.StatusNotFound, err.Error())
					return
				}
				if len(*upBaskets) == 0 {
					err := e.transactionUsecase.DeleteBasketById(int((*baskets)[i].ID))
					if err != nil {
						utils.HandleError(c, http.StatusNotFound, err.Error())
						return
					}
				}
				utils.HandleSucces(c, "success delete sub basket")
				return
			}
		} 
	}
	utils.HandleError(c, http.StatusNotFound, "sub basket id is not found")
}

func (e *TransactionHandler) viewBasket(c *gin.Context) {
	user, err := e.userUsecase.ValidUser(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	baskets, err := e.transactionUsecase.ViewBasketByUserId(int(user.ID))
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	basketList, err := e.transactionUsecase.ViewBasketList(baskets)
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	utils.HandleSucces(c, basketList)
}

func (e *TransactionHandler) checkShippingStore(c *gin.Context) {
	var shipping = model.CheckShippingInput{}
	err := c.Bind(&shipping)
	if err != nil {
		fmt.Printf("[TransactionHandler.checkShippingStore] error bind data %v \n", err)
		utils.HandleError(c, http.StatusInternalServerError, "Oopss server someting wrong")
		return
	}
	fmt.Println(shipping)
	product, err := e.productUsecase.ViewProductById(shipping.ProductID)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, err.Error())
		return
	}
	mitra, err := e.userUsecase.ViewMitraById(int(product.MitraID))
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	courier, err := e.ongkirUsecase.ViewCourierMitraByMitraId(int(mitra.ID))
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	courierShow, err := e.ongkirUsecase.ViewCourierMitraShow(courier)
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	var couriers []model.DataShipping
	var shippingDetail = &model.Shipping{}
	for i := 0; i < len(*courierShow); i++ {
		inShipping := model.QueryDetail{
			Origin:      strconv.Itoa(shipping.CityID),
			Destination: strconv.Itoa(mitra.CityID),
			Weight:      product.Weight * shipping.Quantity,
			Courier:     (*courierShow)[i].CourierName,
		}
		shippingDetail, err = e.ongkirUsecase.CalculateShipping(&inShipping)
		if err != nil {
			utils.HandleError(c, http.StatusInternalServerError, err.Error())
			return
		}
		if len(shippingDetail.Result) != 0 && len(shippingDetail.Result[0].Costs) != 0 {
			couriers = append(couriers, shippingDetail.Result[0])
		}
	}
	shippingDetail.Result = couriers
	utils.HandleSucces(c, shippingDetail)
}

func (e *TransactionHandler) addUserAddress(c *gin.Context) {
	user, err := e.userUsecase.ValidUser(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	var address = model.UserAddress{}
	err = c.Bind(&address)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, "Oopss server someting wrong")
		return
	}
	if address.ID != 0 || address.UserID != 0 {
		utils.HandleError(c, http.StatusBadRequest, "Inputs are not permitted")
		return
	}
	address.UserID = user.ID
	inAddress, err := e.transactionUsecase.AddUserAddress(&address)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSucces(c, inAddress)
}

func (e *TransactionHandler) editeUserAddress(c *gin.Context) {
	user, err := e.userUsecase.ValidUser(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	idStr := c.Param("idAddress")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "id has be number")
		return
	}
	var address = model.UserAddress{}
	err = c.Bind(&address)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, "Oopss server someting wrong")
		return
	}
	if address.ID != 0 || address.UserID != 0 {
		utils.HandleError(c, http.StatusBadRequest, "inputs are not permitted")
		return
	}
	allAddress, err := e.transactionUsecase.ViewUserAddressByUserId(int(user.ID))
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	if len(*allAddress) == 0 {
		utils.HandleError(c, http.StatusForbidden, "you don't have an address yet")
		return
	}

	for i := 0; i < len(*allAddress); i++ {
		if uint(id) == (*allAddress)[i].ID {
			if address.CityID == 0 {
				address.CityID = (*allAddress)[i].CityID
			}
			upAddress, err := e.transactionUsecase.UpdateUserAddressById(id, &address)
			if err != nil {
				utils.HandleError(c, http.StatusInternalServerError, err.Error())
				return
			}
			utils.HandleSucces(c, upAddress)
			return
		}
	}
	utils.HandleError(c, http.StatusNotFound, "id adress not exsis")
}

func (e *TransactionHandler) viewMyAddress(c *gin.Context) {
	user, err := e.userUsecase.ValidUser(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	address, err := e.transactionUsecase.ViewUserAddressByUserId(int(user.ID))
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	if len(*address) == 0 {
		utils.HandleError(c, http.StatusNotFound, "you do not have a recipient address")
		return
	}
	utils.HandleSucces(c, address)
}

func (e *TransactionHandler) removeMyAddress(c *gin.Context) {
	user, err := e.userUsecase.ValidUser(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	idStr := c.Param("idAddress")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "id has be number")
		return
	}
	address, err := e.transactionUsecase.ViewUserAddressByUserId(int(user.ID))
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	if (*address)[0].ID == uint(id) {
		utils.HandleError(c, http.StatusForbidden, "cannot delete default address data")
		return
	}
	for i := 0; i < len(*address); i++ {
		for (*address)[i].ID == uint(id) {
			err := e.transactionUsecase.DeleteUserAddresById(id)
			if err != nil {
				utils.HandleError(c, http.StatusNotFound, err.Error())
				return
			}
			utils.HandleSucces(c, "Success delete data address")
			return
		}
	}
	utils.HandleError(c, http.StatusNotFound, "Id address not found")
}
