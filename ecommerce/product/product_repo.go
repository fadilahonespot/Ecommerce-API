package product

import (
	"ecommerce/model"

	"github.com/jinzhu/gorm"
)

type ProductRepo interface {
	ViewAllCatagory()(*[]model.Catagory, error)
	InsertCatagory(catagory *model.Catagory) (*model.Catagory, error)
	ViewCatagoryByIdUser(id int)(*model.Catagory, error)
	ViewAllSubCatagory()(*[]model.SubCatagory, error)
	InsertSubCatagory(subcatagory *model.SubCatagory)(*model.SubCatagory, error)
	ViewSubPlusCatagory(idOrNil int)(*[]model.SubPlusCatagory, error)
	InsertProduct(product *model.Product, tx *gorm.DB)(*model.Product, error)
	ViewAllProduct()(*[]model.Product, error) 
	InsertProductImg(productImg *model.ProductImg, tx *gorm.DB)(*model.ProductImg, error)
	ProductByMitraId(id int)(*[]model.Product, error)
	ImageByProductId(id int)(*[]model.ProductImg, error)
	SubCatagoryById(id int)(*model.SubCatagory, error)
	CatagoryById(id int) (*model.Catagory, error)
	ProductReviewByProductId(id int)(*[]model.ProductReview, error)
	UpdateProduct(id int, tx *gorm.DB, product *model.Product)(*model.Product, error)
	DeleteProductImgByProductId(id int, tx *gorm.DB) error 
	DeleteProductById(id int, tx *gorm.DB) error
	AddWachlist(wachlist *model.Watchlist)(*model.Watchlist, error)
	ViewProductById(id int)(*model.Product, error)
	ViewWachlistByUserId(id int)(*[]model.Watchlist, error)
	ViewWachlistByProductId(id int)(*model.Watchlist, error)
	DeleteWachlistByProductId(id int, tx *gorm.DB) error 
	DeleteWatchlistById(id int) error
	UpdateCatagoryById(id int, catagory *model.Catagory)(*model.Catagory, error) 
	UpdateSubCatagory(id int, subcatagory *model.SubCatagory)(*model.SubCatagory, error)
	InsertProductDiscussin(discussion *model.ProductDiscussion)(*model.ProductDiscussion, error)
	ViewProductDiscussionByProductId(id int)(*[]model.ProductDiscussion, error)
}