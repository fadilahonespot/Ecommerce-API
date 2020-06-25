package handler

import (
	"ecommerce/middleware"
	"ecommerce/model"
	"ecommerce/product"
	"ecommerce/user"
	"ecommerce/utils"
	"fmt"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	// "github.com/swaggo/swag/example/celler/httputil"
	// "github.com/swaggo/swag/example/celler/model"
)

type ProductHandler struct {
	productUsecase product.ProductUsecase
	userUsecase    user.UserUsecase
}

func CreateProductHandler(r *gin.Engine, productUsecase product.ProductUsecase, userUsecase user.UserUsecase) {
	productHandler := ProductHandler{productUsecase, userUsecase}

	r.GET("/catagory", productHandler.viewAllCatagory)
	r.GET("/subcatagory", productHandler.viewAllSubCatagory)
	r.GET("catagories", productHandler.viewAllSubPlusCatagory)
	r.GET("/catagories/:idCatagory", productHandler.viewSubPlusCatagoryById)
	r.GET("/product", productHandler.viewAllProduct)
	r.GET("/product/discussion/:idProduct", productHandler.viewPoductDiscussion)

	r2 := r.Group("/admin").Use(middleware.TokenVerifikasiMiddleware())
	r2.POST("/catagory", productHandler.addCatagory)
	r2.PUT("/catagory/:idCatagory", productHandler.editCatagory)
	r2.POST("/subcatagory", productHandler.addSubCatagory)
	r2.PUT("/subcatagory/:idSubcatagory", productHandler.editSubcatagory)
	r2.DELETE("/product/:idProduct", productHandler.deleteProduct)

	r3 := r.Group("/mitra").Use(middleware.TokenVerifikasiMiddleware())
	r3.POST("/product", productHandler.addProduct)
	r3.GET("/product", productHandler.myProduct)
	r3.GET("/product/:idProduct", productHandler.myProductById)
	r3.PUT("/product/:idProduct", productHandler.updateMyProduct)
	r3.DELETE("/product/:idProduct", productHandler.deleteMyProduct)

	r4 := r.Group("/user").Use(middleware.TokenVerifikasiMiddleware())
	r4.POST("/wacthlist/:idProduct", productHandler.addWachlist)
	r4.GET("/wacthlist", productHandler.viewMyWachlist)
	r4.DELETE("/wacthlist/:idProduct", productHandler.deleteWachlist)
	r4.POST("/product/discussion", productHandler.addProductDisscusion)
}

// viewAllCatagory godoc
// @Summary Show all catagory product ecommerce
// @Description get all catagory
// @ID get-allcatagory
// @Accept  json
// @Produce  json
// @Success 200 {object} model.ResponsCat
// @Failure 400 {object} model.ResponsFalse{message:"Id is not exsis"}
// @Failure 404 {object} model.ResponsFalse
// @Failure 500 {object} model.ResponsFalse{data=string}
// @Router /catagory [get]
func (e *ProductHandler) viewAllCatagory(c *gin.Context) {
	catagory, err := e.productUsecase.ViewAllCatagory()
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	if len(*catagory) == 0 {
		utils.HandleError(c, http.StatusNotFound, "catagory is empty")
		return
	}
	utils.HandleSucces(c, catagory)
}

func (e *ProductHandler) viewAllSubCatagory(c *gin.Context) {
	subCatagories, err := e.productUsecase.ViewAllSubCatagory()
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	if len(*subCatagories) == 0 {
		utils.HandleError(c, http.StatusNotFound, "catagory is empty")
		return
	}
	utils.HandleSucces(c, subCatagories)
}

func (e *ProductHandler) viewAllSubPlusCatagory(c *gin.Context) {
	catagories, err := e.productUsecase.ViewSubPlusCatagory(0)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSucces(c, catagories)
}

func (e *ProductHandler) viewSubPlusCatagoryById(c *gin.Context) {
	idStr := c.Param("idCatagory")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "id has be number")
		return
	}
	catagories, err := e.productUsecase.ViewSubPlusCatagory(id)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, err.Error())
		return
	}
	utils.HandleSucces(c, catagories)
}

func (e *ProductHandler) viewAllProduct(c *gin.Context) {
	allMitra, err := e.userUsecase.ViewAllMitra()
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	var products *[]model.Product
	var productShow *[]model.ProductShow
	var upProduct []model.ProductShow
	for i := 0; i < len(*allMitra); i++ {
		products, err = e.productUsecase.ProductByMitraId(int((*allMitra)[i].ID))
		if err != nil {
			utils.HandleError(c, http.StatusNotFound, err.Error())
			return
		}
		mitra, err := e.userUsecase.ViewMitraById(int((*allMitra)[i].ID))
		if err != nil {
			utils.HandleError(c, http.StatusNotFound, err.Error())
			return
		}
		if len(*products) > 0 {
			productShow, err = e.productUsecase.ViewProductShow(products, mitra)
			if err != nil {
				utils.HandleError(c, http.StatusForbidden, err.Error())
				return
			}
			for j := 0; j < len(*productShow); j++ {
				upProduct = append(upProduct, (*productShow)[j])
			}
		}
	}
	if len(upProduct) == 0 {
		utils.HandleError(c, http.StatusNotFound, "there are no products in the list")
		return
	}
	result, err := utils.PaginationProductShow(c, &upProduct)
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	utils.HandleSucces(c, result)
}

func (e *ProductHandler) addWachlist(c *gin.Context) {
	user, err := e.userUsecase.ValidUser(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	idStr := c.Param("idProduct")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "id product has be number")
		return
	}
	_, err = e.productUsecase.ViewProductById(id)
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	wachlistUser, err := e.productUsecase.ViewWachlistByUserId(int(user.ID))
	if err == nil {
		for i := 0; i < len(*wachlistUser); i++ {
			if (*wachlistUser)[i].ProductID == uint(id) {
				utils.HandleError(c, http.StatusForbidden, "the product is already on the watchlist")
				return
			}
		}
	}
	mitra, err := e.userUsecase.ViewMitraByUserId(int(user.ID))
	if err == nil {
		products, err := e.productUsecase.ProductByMitraId(int(mitra.ID))
		if err == nil {
			for i := 0; i < len(*products); i++ {
				if uint(id) == (*products)[i].ID {
					utils.HandleError(c, http.StatusBadRequest, "can't add products from the store itself")
					return
				} 
			}
		}
	}
	var inWachlist = model.Watchlist{
		ProductID: uint(id),
		UserID: user.ID,
	}
	wachlist, err := e.productUsecase.AddWachlist(&inWachlist)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSucces(c, wachlist)
}

func (e *ProductHandler) viewMyWachlist(c *gin.Context) {
	user, err := e.userUsecase.ValidUser(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	wachlist, err := e.productUsecase.ViewWachlistByUserId(int(user.ID))
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	var products []model.Product
	var productShow []model.ProductShow
	for i := 0; i < len(*wachlist); i++ {
		product, err := e.productUsecase.ViewProductById(int((*wachlist)[i].ProductID))
		if err != nil {
			utils.HandleError(c, http.StatusNotFound, err.Error())
			return
		}
		mitra, err := e.userUsecase.ViewMitraById(int(product.MitraID))
		if err != nil {
			utils.HandleError(c, http.StatusNotFound, err.Error())
			return
		}
		products = append(products, *product)
		outProduct, err := e.productUsecase.ViewProductShow(&products, mitra)
		if err != nil {
			utils.HandleError(c, http.StatusForbidden, err.Error())
			return
		}
		for j := 0; j < len(*outProduct); j++ {
			if len(*outProduct) - 1 == j {
				productShow = append(productShow, (*outProduct)[j])
			}
		}
	}
	if len(productShow) == 0 {
		utils.HandleError(c, http.StatusNotFound, "your wacthlist is empty")
		return
	}
	utils.HandleSucces(c, productShow)
}

func (e *ProductHandler) deleteWachlist(c *gin.Context) {
	user, err := e.userUsecase.ValidUser(c)
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
	wacthlist, err := e.productUsecase.ViewWachlistByUserId(int(user.ID))
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, "You don't have a wacthlist product list")
		return
	}
	for i := 0; i < len(*wacthlist); i++ {
		if (*wacthlist)[i].ProductID == uint(id) {
			err = e.productUsecase.DeleteWachlistByProductId(id)
			if err != nil {
				utils.HandleError(c, http.StatusNotFound, err.Error())
				return
			}
			utils.HandleSucces(c, "Success delete wachlist")
			return
		}
	}
	utils.HandleError(c, http.StatusNotFound, "watchlist id not found, product id is not on the watchlist")
}

func (e *ProductHandler) addProductDisscusion(c *gin.Context) {
	user, err := e.userUsecase.ValidUser(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	var disscusion = model.ProductDiscussion{}
	err = c.Bind(&disscusion)
	if err != nil {
		fmt.Printf("[productHandler.addProductDisscusion] error bind data %v \n", err)
		utils.HandleError(c, http.StatusInternalServerError, "Oopss server someting wrong")
		return
	}
	if disscusion.UserID != 0 || disscusion.ID != 0 {
		utils.HandleError(c, http.StatusBadRequest, "inputs are not permitted")
		return
	}
	disscusion.UserID = user.ID
	_, err = e.productUsecase.ViewProductById(int(disscusion.ProductID))
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, err.Error())
		return
	}
	inDisscusion, err := e.productUsecase.InsertProductDiscussin(&disscusion)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSucces(c, inDisscusion)
}

func (e *ProductHandler) viewPoductDiscussion(c *gin.Context) {
	idStr := c.Param("idProduct")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "id product has be number")
		return
	}
	discussion, err := e.productUsecase.ViewProductDiscussionByProductId(id)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	if len(*discussion) == 0 {
		utils.HandleError(c, http.StatusNotFound, "there is no discussion on this product")
		return
	}
	utils.HandleSucces(c, discussion)
}
