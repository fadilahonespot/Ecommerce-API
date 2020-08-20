package product

import (
	"ecommerce/model"

	"github.com/gin-gonic/gin"
)

type ProductUsecase interface {
	ViewAllCatagory()(*[]model.Catagory, error)
	InsertCatagory(catagory *model.Catagory) (*model.Catagory, error)
	ViewCatagoryByIdUser(id int)(*model.Catagory, error)
	ViewAllSubCatagory()(*[]model.SubCatagory, error)
	InsertSubCatagory(subcatagory *model.SubCatagory)(*model.SubCatagory, error)
	ViewSubPlusCatagory(idOrNil int)(*[]model.SubPlusCatagory, error)
	ViewAllProduct()(*[]model.Product, error) 
	ProductByMitraId(id int)(*[]model.Product, error)
	ImageByProductId(id int)(*[]model.ProductImg, error)
	SubCatagoryById(id int)(*model.SubCatagory, error)
	ViewProductShow(product *[]model.Product, mitra *model.Mitra) (*[]model.ProductShow, error)
	AddAndUpdateProductShow(idMitra int, idProduct int, c *gin.Context)(*model.Product, []string, error)
	DeleteProductShow(idProduct int, products *[]model.Product) error
	AddWachlist(wachlist *model.Watchlist)(*model.Watchlist, error)
	ViewProductById(id int)(*model.Product, error)
	ViewWachlistByUserId(id int)(*[]model.Watchlist, error)
	ViewWachlistByProductId(id int)(*model.Watchlist, error)
	DeleteWatchlistById(id int) error
	UpdateCatagoryById(id int, catagory *model.Catagory)(*model.Catagory, error) 
	UpdateSubCatagory(id int, subcatagory *model.SubCatagory)(*model.SubCatagory, error)
	DeleteWachlistByProductId(id int) error
	InsertProductDiscussin(discussion *model.ProductDiscussion)(*model.ProductDiscussion, error)
	ViewProductDiscussionByProductId(id int)(*[]model.ProductDiscussion, error)
}