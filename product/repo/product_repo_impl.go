package repo

import (
	"ecommerce/model"
	"ecommerce/product"
	"fmt"

	"github.com/jinzhu/gorm"
)

type ProductRepoImpl struct {
	DB *gorm.DB
}

func CreateProducRepo(DB *gorm.DB) product.ProductRepo {
	return &ProductRepoImpl{DB}
}

func (e *ProductRepoImpl) InsertProduct(product *model.Product, tx *gorm.DB)(*model.Product, error) {
	err := tx.Save(&product).Error
	if err != nil {
		fmt.Printf("[ProductRepoImpl.InsertProduct] error execute query %v \n", err)
		return nil, fmt.Errorf("failed insert data product")
	}
	return product, nil
}

func (e *ProductRepoImpl) ViewAllProduct()(*[]model.Product, error) {
	var product []model.Product
	err := e.DB.Find(&product).Error
	if err != nil {
		fmt.Printf("[ProductRepoImpl.ViewAllProduct] error execute query %v \n", err)
		return nil, fmt.Errorf("failed view all product")
	}
	return &product, nil
}

func (e *ProductRepoImpl) InsertProductImg(productImg *model.ProductImg, tx *gorm.DB)(*model.ProductImg, error) {
	err := tx.Save(&productImg).Error
	if err != nil {
		fmt.Printf("[ProductRepoImpl.UploadProductImg] error execute query %v \n", err)
		return nil, fmt.Errorf("failed upload image product")
	}
	return productImg, nil
}

func (e *ProductRepoImpl) DeleteProductImgByProductId(id int, tx *gorm.DB) error {
	var productImg = model.ProductImg{}
	err := tx.Table("product_img").Where("product_id = ?", id).Delete(&productImg).Error
	if err != nil {
		fmt.Printf("[ProductRepoImpl.DeleteProductImgByProductId] error execute query %v \n", err)
		return fmt.Errorf("failed delete data product img, id is not exist")
	}
	return nil
}

func (e *ProductRepoImpl) ProductByMitraId(id int)(*[]model.Product, error) {
	var products []model.Product
	err := e.DB.Table("product").Where("mitra_id = ?", id).Find(&products).Error
	if err != nil {
		fmt.Printf("[ProductRepoImpl.ProductByMitraId] error execute query %v \n", err)
		return nil, fmt.Errorf("id mitra is not exist")
	}
	return &products, nil
}

func (e *ProductRepoImpl) ImageByProductId(id int)(*[]model.ProductImg, error) {
	var imgProduct []model.ProductImg
	err := e.DB.Table("product_img").Where("product_id = ?", id).Find(&imgProduct).Error
	if err != nil {
		fmt.Printf("[ProductRepoImpl.ImageByProductId] error execute query %v \n", err)
		return nil, fmt.Errorf("id product is not exist")
	}
	return &imgProduct, nil
}

func (e *ProductRepoImpl) ProductReviewByProductId(id int)(*[]model.ProductReview, error) {
	var productRev []model.ProductReview
	err := e.DB.Table("product_review").Where("product_id = ?", id).Find(&productRev).Error
	if err != nil {
		fmt.Printf("[productRepo.ViewProductReview] error execute query %v \n", err)
		return nil, fmt.Errorf("failed view data product review")
	}
	return &productRev, nil 
}

func (e *ProductRepoImpl) UpdateProduct(id int, tx *gorm.DB, product *model.Product)(*model.Product, error) {
	var upProduct = model.Product{}
	err := tx.Table("product").Where("id = ?", id).First(&upProduct).Update(&product).Error
	if err != nil {
		fmt.Printf("[ProductRepoImpl.UpdateProduct] error execute query %v \n", err)
		return nil, fmt.Errorf("failed update product")
	}
	return &upProduct, nil
}

func (e *ProductRepoImpl) DeleteProductById(id int, tx *gorm.DB) error {
	var product = model.Product{}
	err := tx.Table("product").Where("id = ?", id).Delete(&product).Error
	if err != nil {
		fmt.Printf("[ProductRepoImpl.DeleteProductById] error execute query %v \n", err)
		return fmt.Errorf("failed delete product, id is not exsis")
	}
	return nil
}

func (e *ProductRepoImpl) ViewProductById(id int)(*model.Product, error) {
	var product = model.Product{}
	err := e.DB.Table("product").Where("id = ?", id).First(&product).Error
	if err != nil {
		fmt.Printf("[ProductRepoImpl.ViewProductById] error execute query %v \n", err)
		return nil, fmt.Errorf("id product is not exist")
	}
	return &product, nil
}

func (e *ProductRepoImpl) AddWachlist(wachlist *model.Watchlist)(*model.Watchlist, error) {
	err := e.DB.Save(&wachlist).Error
	if err != nil {
		fmt.Printf("[ProductRepoImpl.AddWachlist] error execute query %v \n", err)
		return nil, fmt.Errorf("failed add product at wachlist")
	}
	return wachlist, nil
}

func (e *ProductRepoImpl) ViewWachlistByUserId(id int)(*[]model.Watchlist, error) {
	var wachlist []model.Watchlist
	err := e.DB.Table("wachlist").Where("user_id = ?", id).Find(&wachlist).Error
	if err != nil {
		fmt.Printf("[ProductRepoImpl.ViewWachlistByUserId] error execute query %v \n", err)
		return nil, fmt.Errorf("id user not exist in wachlist")
	}
	return &wachlist, nil
}

func (e *ProductRepoImpl) ViewWachlistByProductId(id int)(*model.Watchlist, error) {
	var wachlist = model.Watchlist{}
	err := e.DB.Table("wachlist").Where("product_id = ?", id).First(&wachlist).Error
	if err != nil {
		fmt.Printf("[ProductRepoImpl] error execute query %v \n", err)
		return nil, fmt.Errorf("id product in wachlist is not exist")
	}
	return &wachlist, nil
}

func (e *ProductRepoImpl) DeleteWachlistByProductId(id int, tx *gorm.DB) error {
	var wachlist = model.Watchlist{}
	err := tx.Table("wachlist").Where("product_id = ?", id).Delete(&wachlist).Error
	if err != nil {
		fmt.Printf("[ProductRepoImpl] error execute query %v \n", err)
		return fmt.Errorf("failed delete wachlist, id product not exsis")
	}
	return nil
}

func (e *ProductRepoImpl) DeleteWatchlistById(id int) error {
	var wacthlist = model.Watchlist{}
	err := e.DB.Table("wachlist").Where("id = ?", id).Delete(&wacthlist).Error
	if err != nil {
		fmt.Printf("[ProductRepoImpl.DeleteWatclistById] error execute query %v \n", err)
		return fmt.Errorf("failed delete watchlist, id is not exist")
	}
	return nil
}

func (e *ProductRepoImpl) InsertProductDiscussin(discussion *model.ProductDiscussion)(*model.ProductDiscussion, error) {
	err := e.DB.Save(&discussion).Error
	if err != nil {
		fmt.Printf("[ProductRepoImpl.InsertProductDiscussio] error execute query %v \n", err)
		return nil, fmt.Errorf("failed add data to product discussion")
	}
	return discussion, nil
}

func (e *ProductRepoImpl) ViewProductDiscussionByProductId(id int)(*[]model.ProductDiscussion, error) {
	var disscusion []model.ProductDiscussion
	err := e.DB.Table("product_discussion").Where("product_id = ?", id).Find(&disscusion).Error
	if err != nil {
		fmt.Printf("[ProductRepoImpl] error execute query %v \n", err)
		return nil, fmt.Errorf("failed view data product discussion")
	}
	return &disscusion, nil
}
