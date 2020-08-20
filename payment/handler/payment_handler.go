package handler

import (
	"ecommerce/middleware"
	"ecommerce/model"
	"ecommerce/payment"
	"ecommerce/user"
	"ecommerce/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	paymentUsecase payment.PaymentUsecase
	userUsecase user.UserUsecase
}

func CreatePaymentHandler(r *gin.Engine, paymentUsecase payment.PaymentUsecase, userUsecase user.UserUsecase) {
	paymentHandler := PaymentHandler{paymentUsecase, userUsecase}

	r2 := r.Group("/admin").Use(middleware.TokenVerifikasiMiddleware())
	r2.POST("/payments", paymentHandler.addPaymentMethod)
	r2.PUT("/payments/:idPaymentMethod", paymentHandler.editPaymentMethod)
	r2.GET("/payments", paymentHandler.viewPaymentMethods)
	r2.DELETE("/payments/:idPaymentMethod", paymentHandler.deletePaymentMethod)
	r2.POST("/payment", paymentHandler.addPayment)
	r2.PUT("/payment/:idPayment", paymentHandler.editPayment)
	r2.GET("/payment", paymentHandler.viewAllPayment)
	r2.DELETE("/payment/:idPayment", paymentHandler.deletePayment)
}

func (e *PaymentHandler) addPaymentMethod(c *gin.Context) {
	_, err := e.userUsecase.AdminOnly(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	var payment = model.PaymentMethod{}
	err = c.Bind(&payment)
	if err != nil {
		fmt.Printf("[PaymentHandler.addPaymentMethods] error bind data body %v \n", err)
		utils.HandleError(c, http.StatusInternalServerError, "Oopss server someting wrong")
		return
	}
	if payment.PaymentType == "" {
		utils.HandleError(c, http.StatusBadRequest, "colum cannot be empty")
		return
	}
	inPayment, err := e.paymentUsecase.InsertPaymentMethod(&payment)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSucces(c, inPayment)
}

func (e *PaymentHandler) editPaymentMethod(c *gin.Context) {
	_, err := e.userUsecase.AdminOnly(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	idStr := c.Param("idPaymentMethod")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "id has be number")
		return
	}
	var payment = model.PaymentMethod{}
	err = c.Bind(&payment)
	if err != nil {
		fmt.Printf("[PaymentHandler.updatePaymentMethod] error bind data body %v \n", err)
		utils.HandleError(c, http.StatusInternalServerError, "Oopss server someting wrong")
		return
	}
	if payment.ID != 0 {
		utils.HandleError(c, http.StatusBadRequest, "input are not permitted")
		return
	}
	if payment.PaymentType == "" {
		utils.HandleError(c, http.StatusBadRequest, "column cannot be empty")
		return
	}
	_, err = e.paymentUsecase.ViewPaymentMethodById(id)
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	upPayment, err := e.paymentUsecase.UpdatePaymentMethodById(id, &payment)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSucces(c, upPayment)
}

func (e *PaymentHandler) viewPaymentMethods(c *gin.Context) {
	_, err := e.userUsecase.AdminOnly(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	payments, err := e.paymentUsecase.ViewAllPaymentMethod()
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
	}
	if len(*payments) == 0 {
		utils.HandleError(c, http.StatusNotFound, "payment method list is empty")
		return
	}
	utils.HandleSucces(c, payments)
}

func (e *PaymentHandler) deletePaymentMethod(c *gin.Context) {
	_, err := e.userUsecase.AdminOnly(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	idStr := c.Param("idPaymentMethod")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "id has be number")
		return
	}
	_, err = e.paymentUsecase.ViewPaymentMethodById(id)
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	err = e.paymentUsecase.DeletePaymentMethodById(id)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSucces(c, "Succes delete data payment methods")
}

func (e *PaymentHandler) addPayment(c *gin.Context) {
	_, err := e.userUsecase.AdminOnly(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	var inPayment = model.Payment{} 
	err = c.Bind(&inPayment)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, "Oppss server someting wrong")
		return
	}
	if inPayment.RekeningAccount == "" || inPayment.RekeningName == "" || inPayment.RekeningNumber == "" {
		utils.HandleError(c, http.StatusBadRequest, "column cannot be empty")
		return
	}
	_, err = e.paymentUsecase.ViewPaymentMethodById(int(inPayment.PaymentMethodID))
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	payment, err := e.paymentUsecase.InsertPayment(&inPayment)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSucces(c, payment)
}

func (e *PaymentHandler) editPayment(c *gin.Context) {
	_, err := e.userUsecase.AdminOnly(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	idStr := c.Param("idPayment")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "id has be number")
		return
	}
	var inPayment = model.Payment{}
	err = c.Bind(&inPayment)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, "Oopss server someting wrong")
		return
	}
	if inPayment.ID != 0 || inPayment.PaymentMethodID != 0 {
		utils.HandleError(c, http.StatusBadRequest, "inputs are not permitted")
		return
	}
	if inPayment.RekeningAccount == "" || inPayment.RekeningName == "" || inPayment.RekeningNumber == "" {
		utils.HandleError(c, http.StatusBadRequest, "column cannot be empty")
		return
	}
	_, err = e.paymentUsecase.ViewPaymentById(id)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, err.Error())
		return
	}
	payment, err := e.paymentUsecase.UpdatePayment(id, &inPayment)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, err.Error())
		return
	}
	utils.HandleSucces(c, payment)
}

func (e *PaymentHandler) viewAllPayment(c *gin.Context) {
	_, err := e.userUsecase.AdminOnly(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	paymentMethods, err := e.paymentUsecase.ViewAllPaymentMethod()
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	if len(*paymentMethods) == 0 {
		utils.HandleError(c, http.StatusNotFound, "no list can be displayed")
		return
	}
	var paymentShow = model.PaymentShow{}
	var paymentsShow []model.PaymentShow
	var subPaymentShow = model.SubPaymentShow{}
	for i := 0; i < len(*paymentMethods); i++ {
		paymentShow.PaymentMethodID = (*paymentMethods)[i].ID
		paymentShow.PaymentType = (*paymentMethods)[i].PaymentType
		payments, err := e.paymentUsecase.ViewPaymentByPaymentMethodId(int((*paymentMethods)[i].ID))
		if err != nil {
			utils.HandleError(c, http.StatusNotFound, err.Error())
			return
		}
		for k := 0; k < len(*payments); k++ {
			subPaymentShow = model.SubPaymentShow{
				PaymentID: (*payments)[k].ID,
				RekeningAccount: (*payments)[k].RekeningAccount,
				RekeningName: (*payments)[k].RekeningName,
				RekeningNumber: (*payments)[k].RekeningNumber,
			}
			paymentShow.Payments = append(paymentShow.Payments, subPaymentShow)
		}
		paymentsShow = append(paymentsShow, paymentShow)
		paymentShow.Payments = nil
	}
	utils.HandleSucces(c, paymentsShow)
}

func (e *PaymentHandler) deletePayment(c *gin.Context) {
	_, err := e.userUsecase.AdminOnly(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	idStr := c.Param("idPayment")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "id has be number")
		return
	}
	_, err = e.paymentUsecase.ViewPaymentById(id)
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	err = e.paymentUsecase.DeletePaymentById(id)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSucces(c, "Success delete payment")
}

