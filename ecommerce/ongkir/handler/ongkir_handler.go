package handler

import (
	"ecommerce/middleware"
	"ecommerce/model"
	"ecommerce/ongkir"
	"ecommerce/user"
	"ecommerce/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OngkirHandler struct {
	ongkirUsecase ongkir.OngkirUsecase
	userUsecase   user.UserUsecase
}

func CreateOngkirHandler(r *gin.Engine, ongkirUsecase ongkir.OngkirUsecase, userUsecase user.UserUsecase) {
	ongkirHandler := OngkirHandler{ongkirUsecase, userUsecase}

	r.GET("/city", ongkirHandler.viewAllCities)
	r.GET("/city/:id", ongkirHandler.viewCitiById)
	r.GET("/province", ongkirHandler.viewAllProvinces)
	r.POST("/cost", ongkirHandler.calculateShipping)
	r.GET("/province/:id", ongkirHandler.getCityByProvince)
	r.POST("/tracking", ongkirHandler.trackingPackage)
	r.GET("/courier", ongkirHandler.viewAllCourier)

	r2 := r.Group("/admin").Use(middleware.TokenVerifikasiMiddleware())
	r2.POST("/courier", ongkirHandler.addCourier)
	r2.PUT("/courier/:idCourier", ongkirHandler.updateCourier)
	r2.DELETE("/courier/:idCourier", ongkirHandler.deleteCourier)

	r3 := r.Group("/mitra").Use(middleware.TokenVerifikasiMiddleware())
	r3.POST("/courier", ongkirHandler.addCourierMitra)
	r3.DELETE("/courier/:idCourierMitra", ongkirHandler.deleteCourierMitra)
	r3.GET("/courier", ongkirHandler.viewMitraCourier)

}

func (e *OngkirHandler) viewAllCities(c *gin.Context) {
	city, err := e.ongkirUsecase.GetCity()
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSucces(c, city)
}

func (e *OngkirHandler) viewCitiById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "id has be number")
		return
	}
	city, err := e.ongkirUsecase.GetCityById(id)
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	utils.HandleSucces(c, city)
}

func (e *OngkirHandler) viewAllProvinces(c *gin.Context) {
	provinces, err := e.ongkirUsecase.GetProvinces()
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSucces(c, provinces)
}

func (e *OngkirHandler) calculateShipping(c *gin.Context) {
	var bodyShip = model.QueryDetail{}
	err := c.Bind(&bodyShip)
	if err != nil {
		fmt.Printf("[userHandler.calculateShipping] error bind data body %v \n", err)
		utils.HandleError(c, http.StatusInternalServerError, "Ooppss internal server error")
		return
	}
	shipping, err := e.ongkirUsecase.CalculateShipping(&bodyShip)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	if len(shipping.Result) == 0 {
		utils.HandleError(c, http.StatusNotFound, "the courier didn't arrive at your place")
		return
	}
	utils.HandleSucces(c, shipping)
}

func (e *OngkirHandler) getCityByProvince(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "id has be number")
		return
	}
	city, err := e.ongkirUsecase.GetCityByProvince(id)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	if len(*city) == 0 {
		utils.HandleError(c, http.StatusNotFound, "id propince is not exsis")
		return
	}
	utils.HandleSucces(c, city)
}

func (e *OngkirHandler) trackingPackage(c *gin.Context) {
	var tracking = model.InputTracking{}
	err := c.Bind(&tracking)
	if err != nil {
		fmt.Printf("[OngkirHandler.TrackingPackage] Error bind data body %v \n", err)
		utils.HandleError(c, http.StatusInternalServerError, "Ooopss internal server error")
		return
	}
	detailTracking, err := e.ongkirUsecase.TrackResi(&tracking)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, err.Error())
		return
	}
	if len(detailTracking.Tracking) == 0 {
		utils.HandleError(c, http.StatusNotFound, "there is no tracking data yet")
		return
	}
	utils.HandleSucces(c, detailTracking)
}

func (e *OngkirHandler) addCourier(c *gin.Context) {
	_, err := e.userUsecase.AdminOnly(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	var courier = model.Courier{}
	err = c.Bind(&courier)
	if err != nil {
		fmt.Printf("[OngkirHandler.addCourier] error bind data body %v \n", err)
		utils.HandleError(c, http.StatusInternalServerError, "Oopsss server someting wrong")
		return
	}
	if courier.ID != 0 {
		utils.HandleError(c, http.StatusBadRequest, "inputs are not permitted")
		return
	}
	if courier.CourierName == "" {
		utils.HandleError(c, http.StatusBadRequest, "column cannot be empty")
		return
	}
	outCourier, err := e.ongkirUsecase.InsertCourier(&courier)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSucces(c, outCourier)
}

func (e *OngkirHandler) viewAllCourier(c *gin.Context) {
	couriers, err := e.ongkirUsecase.ViewAllCourier()
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSucces(c, couriers)
}

func (e *OngkirHandler) updateCourier(c *gin.Context) {
	_, err := e.userUsecase.AdminOnly(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	idStr := c.Param("idCourier")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "id has be number")
		return
	}
	var courier = model.Courier{}
	err = c.Bind(&courier)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, "Oopss internal server error")
		return
	}
	if courier.ID != 0 {
		utils.HandleError(c, http.StatusBadRequest, "inputs are not permitted")
		return
	}
	if courier.CourierName == "" {
		utils.HandleError(c, http.StatusBadRequest, "column cannot be empty")
		return
	}
	upCourier, err := e.ongkirUsecase.UpdateCourier(id, &courier)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSucces(c, upCourier)
}

func (e *OngkirHandler) deleteCourier(c *gin.Context) {
	_, err := e.userUsecase.AdminOnly(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	idStr := c.Param("idCourier")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "id has be number")
		return
	}
	err = e.ongkirUsecase.DeleteCourierById(id)
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	utils.HandleSucces(c, "Success delete courier")
}

func (e *OngkirHandler) addCourierMitra(c *gin.Context) {
	_, mitra, err := e.userUsecase.ValidMitra(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	var courierMitra = model.CourierMitra{}
	err = c.Bind(&courierMitra)
	if err != nil {
		fmt.Printf("[Ongkirhandler.addCourierMitra] error bind data body %v \n", err)
		utils.HandleError(c, http.StatusInternalServerError, "Ooppss server someting wrong")
		return
	}
	if courierMitra.ID != 0 || courierMitra.MitraID != 0 {
		utils.HandleError(c, http.StatusBadRequest, "inputs are not permitted")
		return
	}
	if courierMitra.CourierID == 0 {
		utils.HandleError(c, http.StatusBadRequest, "column cannot be empty")
		return
	}
	_, err = e.ongkirUsecase.ViewCourierById(int(courierMitra.CourierID))
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, "id courier not found")
		return
	}
	courierMitras, err := e.ongkirUsecase.ViewCourierMitraByMitraId(int(mitra.ID))
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	for i := 0; i < len(*courierMitras); i++ {
		if courierMitra.CourierID == (*courierMitras)[i].CourierID {
			utils.HandleError(c, http.StatusConflict, "courier mitra already exist")
			return
		}
	}
	courierMitra.MitraID = mitra.ID
	inCourier, err := e.ongkirUsecase.InsertCourierMitra(&courierMitra)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSucces(c, inCourier)
}

func (e *OngkirHandler) viewMitraCourier(c *gin.Context) {
	_, mitra, err := e.userUsecase.ValidMitra(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	courierMitras, err := e.ongkirUsecase.ViewCourierMitraByMitraId(int(mitra.ID))
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	if len(*courierMitras) == 0 {
		utils.HandleError(c, http.StatusNotFound, "you do not have a courier list")
		return
	}
	courierShow, err := e.ongkirUsecase.ViewCourierMitraShow(courierMitras)
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
	}
	utils.HandleSucces(c, courierShow)
}

func (e *OngkirHandler) deleteCourierMitra(c *gin.Context) {
	_, mitra, err := e.userUsecase.ValidMitra(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	idStr := c.Param("idCourierMitra")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "id has be number")
		return
	}
	mitras, err := e.ongkirUsecase.ViewCourierMitraByMitraId(int(mitra.ID))
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	for i := 0; i < len(*mitras); i++ {
		if (*mitras)[i].ID == uint(id) {
			err := e.ongkirUsecase.DeleteCourierMitraById(id)
			if err != nil {
				utils.HandleError(c, http.StatusNotFound, err.Error())
				return
			}
			utils.HandleSucces(c, "Success delete courier mitra")
			return
		}
	}
	utils.HandleError(c, http.StatusNotFound, "id courier mitra not found")
}
