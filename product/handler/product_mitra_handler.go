package handler

import (
	"ecommerce/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (e *ProductHandler) addProduct(c *gin.Context) {
	_, mitra, err := e.userUsecase.ValidMitra(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	newProduct, newProductImg, err := e.productUsecase.AddAndUpdateProductShow(int(mitra.ID), 0, c)
	if err != nil {
		utils.HandleError(c, http.StatusForbidden, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"Succes": true, "Message": "Success", "Data": newProduct, "Images": newProductImg})
}

func (e *ProductHandler) updateMyProduct(c *gin.Context) {
	_, mitra, err := e.userUsecase.ValidMitra(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	idStr := c.Param("idProduct")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "id has be number")
		return
	}
	product, err := e.productUsecase.ProductByMitraId(int(mitra.ID))
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	for i := 0; i < len(*product); i++ {
		if id == int((*product)[i].ID) {
			newProduct, newProductImg, err := e.productUsecase.AddAndUpdateProductShow(int(mitra.ID), id, c)
			if err != nil {
				utils.HandleError(c, http.StatusForbidden, err.Error())
				return
			}
			c.JSON(http.StatusOK, gin.H{"Succes": true, "Message": "Success", "Data": newProduct, "Images": newProductImg})
			return
		}
	}
	utils.HandleError(c, http.StatusNotFound, "id product not found")
}

func (e *ProductHandler) deleteMyProduct(c *gin.Context) {
	_, mitra, err := e.userUsecase.ValidMitra(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	idStr := c.Param("idProduct")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "Id has be number")
		return
	}
	products, err := e.productUsecase.ProductByMitraId(int(mitra.ID))
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	err = e.productUsecase.DeleteProductShow(id, products)
	if err != nil {
		utils.HandleError(c, http.StatusForbidden, err.Error())
		return
	}
	utils.HandleSucces(c, "Success delete product")
}

// myProduct godoc
// @Summary Show all My Product
// @Description get all product mitra in mitra
// @ID get-all myproduct mitra
// @Accept  json
// @Produce  json
// @Header 200 {string} Token "qwerty"
// @Success 200 {object} model.Respons{data=[]proto.Product}
// @Failure 400 {object} model.ResponsFalse
// @Failure 404 {object} model.ResponsFalse
// @Failure 500 {object} model.ResponsFalse
// @Security ApiKeyAuth
// @Router /mitra/product [get]
func (e *ProductHandler) myProduct(c *gin.Context) {
	_, mitra, err := e.userUsecase.ValidMitra(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	product, err := e.productUsecase.ProductByMitraId(int(mitra.ID))
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	products, err := e.productUsecase.ViewProductShow(product, mitra)
	if err != nil {
		utils.HandleError(c, http.StatusForbidden, err.Error())
		return
	}
	result, err := utils.PaginationProductShow(c, products)
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	utils.HandleSucces(c, result)
}

func (e *ProductHandler) myProductById(c *gin.Context) {
	_, mitra, err := e.userUsecase.ValidMitra(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	idStr := c.Param("idProduct")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "Id has be number")
		return
	}
	product, err := e.productUsecase.ProductByMitraId(int(mitra.ID))
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	products, err := e.productUsecase.ViewProductShow(product, mitra)
	if err != nil {
		utils.HandleError(c, http.StatusForbidden, err.Error())
		return
	}
	for i := 0; i < len(*products); i++ {
		if uint(id) == (*products)[i].ProductID {
			utils.HandleSucces(c, (*products)[i])
			return
		}
	}
	utils.HandleError(c, http.StatusNotFound, "id product is not exist")
}