package utils

import (
	"ecommerce/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleSucces(c *gin.Context, data interface{}) {
	var returnData = model.Respons{
		Success: true,
		Message: "Success",
		Data: data,
	}
	c.JSON(http.StatusOK, returnData)
}

func HandleError(c *gin.Context, status int, message string) {
	var returnData = model.Respons{
		Success: false,
		Message: message,
	}
	c.JSON(status, returnData)
}