package utils

import (
	"ecommerce/model"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PaginationProductShow(c *gin.Context, product *[]model.ProductShow) (*[]model.ProductShow, error) {
	limitStr := c.Request.URL.Query().Get("limit")
	pageStr := c.Request.URL.Query().Get("page")
	if limitStr == "" || pageStr == "" {
		return product, nil
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return nil, fmt.Errorf("page value has be number")
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return nil, fmt.Errorf("limit value has be number")
	}
	top := page * limit
	bottom := top - limit
	var counter = 0
	var result []model.ProductShow
	for r := 0; r < len(*product); r++ {
		if counter >= bottom && counter < top {
			result = append(result, (*product)[r])
		}
		counter++
	}
	if result == nil {
		return nil, fmt.Errorf("no data, the page exceeds the limit")
	}
	return &result, nil
}
