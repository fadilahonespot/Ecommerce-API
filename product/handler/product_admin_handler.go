package handler

import (
	"ecommerce/model"
	"ecommerce/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (e *ProductHandler) addCatagory(c *gin.Context) {
	user, err := e.userUsecase.AdminOnly(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	var catagory = model.Catagory{}
	err = c.Bind(&catagory)
	if err != nil {
		fmt.Printf("[productHandler.addCatagory] error bind data json %v \n", err)
		utils.HandleError(c, http.StatusInternalServerError, "Ooopss server someting wrong")
		return
	}
	catagories, err := e.productUsecase.ViewAllCatagory()
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	for i := 0; i < len(*catagories); i++ {
		if catagory.CatagoryName == (*catagories)[i].CatagoryName {
			utils.HandleError(c, http.StatusBadRequest, "failed add catagory, category already exis")
			return
		}
	}
	catagory.UserID = user.ID
	newCatagory, err := e.productUsecase.InsertCatagory(&catagory)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSucces(c, newCatagory)
}

func (e *ProductHandler) addSubCatagory(c *gin.Context) {
	user, err := e.userUsecase.AdminOnly(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	var subCatagory = model.SubCatagory{}
	err = c.Bind(&subCatagory)
	if err != nil {
		fmt.Printf("[productHandler.addSubCatagory] error bind data %v \n", err)
		utils.HandleError(c, http.StatusInternalServerError, "Oopsss internal server error")
		return
	}
	subCatagories, err := e.productUsecase.ViewAllSubCatagory()
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}

	for i := 0; i < len(*subCatagories); i++ {
		if (*subCatagories)[i].SubCatagoryName == subCatagory.SubCatagoryName {
			utils.HandleError(c, http.StatusBadRequest, "failed add sub catagory, sub category already exists")
			return
		}
	}
	subCatagory.UserID = user.ID
	newSubcatagory, err := e.productUsecase.InsertSubCatagory(&subCatagory)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSucces(c, newSubcatagory)
}

func (e *ProductHandler) deleteProduct(c *gin.Context) {
	_, err := e.userUsecase.AdminOnly(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	idStr := c.Param("idProduct")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, err.Error())
		return
	}
	products, err := e.productUsecase.ViewAllProduct()
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	err = e.productUsecase.DeleteProductShow(id, products)
	if err != nil {
		utils.HandleError(c, http.StatusForbidden, err.Error())
		return
	}
	utils.HandleSucces(c, "Success delete product")
}

func (e *ProductHandler) editCatagory(c *gin.Context) {
	user, err := e.userUsecase.AdminOnly(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	idStr := c.Param("idCatagory")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "Id catagor has be number")
		return
	}
	catagory := model.Catagory{}
	err = c.Bind(&catagory)
	if err != nil {
		fmt.Printf("[ProductHandler.editCatagory] error bind data catagory %v \n", err)
		utils.HandleError(c, http.StatusBadRequest, "unrecognized data structure")
		return
	}
	if catagory.ID != 0 {
		utils.HandleError(c, http.StatusBadRequest, "id categories cannot be edited")
		return
	}
	catagory.UserID = user.ID
	upCatagory, err := e.productUsecase.UpdateCatagoryById(id, &catagory)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, err.Error())
		return
	}
	utils.HandleSucces(c, upCatagory)
}

func (e *ProductHandler) editSubcatagory(c *gin.Context) {
	user, err := e.userUsecase.AdminOnly(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	var subcatagory = model.SubCatagory{}
	idStr := c.Param("idSubcatagory")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "id subcatagory has be number")
		return
	}
	err = c.Bind(&subcatagory)
	if err != nil {
		fmt.Printf("[ProductHandler.editeSubcatagory] error bind data %v \n", err)
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	if subcatagory.ID != 0 || subcatagory.CatagoryID != 0 {
		utils.HandleError(c, http.StatusBadRequest, "id subcategories or catagory cannot be edited")
		return
	}
	subcatagory.UserID = user.ID
	upSubcatagory, err := e.productUsecase.UpdateSubCatagory(id, &subcatagory)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSucces(c, upSubcatagory)
}

