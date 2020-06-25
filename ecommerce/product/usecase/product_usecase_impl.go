package usecase

import (
	"ecommerce/model"
	"ecommerce/ongkir"
	"ecommerce/product"
	"ecommerce/transaction"
	"ecommerce/user"
	"ecommerce/utils"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/xlzd/gotp"
)

type ProductUsecaseImpl struct {
	productRepo product.ProductRepo
	userRepo    user.UserRepo
	ongkirRepo  ongkir.OngkirRepo
	transactionRepo transaction.TransactionRepo
}

func CreateProductUsecase(productRepo product.ProductRepo, userRepo user.UserRepo, ongkirRepo ongkir.OngkirRepo, transactionRepo transaction.TransactionRepo) product.ProductUsecase {
	return &ProductUsecaseImpl{productRepo, userRepo, ongkirRepo, transactionRepo}
}

func (e *ProductUsecaseImpl) ViewAllProduct() (*[]model.Product, error) {
	return e.productRepo.ViewAllProduct()
}

func (e *ProductUsecaseImpl) ProductByMitraId(id int) (*[]model.Product, error) {
	return e.productRepo.ProductByMitraId(id)
}

func (e *ProductUsecaseImpl) ImageByProductId(id int) (*[]model.ProductImg, error) {
	return e.productRepo.ImageByProductId(id)
}

func (e *ProductUsecaseImpl) ViewProductById(id int)(*model.Product, error) {
	return e.productRepo.ViewProductById(id)
}

func (e *ProductUsecaseImpl) ViewProductShow(product *[]model.Product, mitra *model.Mitra) (*[]model.ProductShow, error) {
	if len(*product) == 0 {
		return nil, fmt.Errorf("your shop doesn't have a product yet")
	}
	var imgProducts []string
	var showProducts []model.ProductShow
	for i := 0; i < len(*product); i++ {
		imgPart, err := e.productRepo.ImageByProductId(int((*product)[i].ID))
		if err != nil {
			return nil, err
		}
		for j := 0; j < len(*imgPart); j++ {
			imgProducts = append(imgProducts, (*imgPart)[j].PartImg)
		}
		subcatagory, err := e.productRepo.SubCatagoryById(int((*product)[i].SubcatagoryID))
		if err != nil {
			return nil, err
		}
		catagory, err := e.productRepo.CatagoryById(int(subcatagory.CatagoryID))
		if err != nil {
			return nil, err
		}
		productRev, err := e.productRepo.ProductReviewByProductId(int((*product)[i].ID))
		if err != nil {
			return nil, err
		}
		var rating, counter int
		var result float64
		for k := 0; k < len(*productRev); k++ {
			rating = rating + (*productRev)[k].Rating
			counter++
		}
		if rating != 0 || counter != 0 {
			result = float64(rating / counter)
		} else {
			result = 0.00001
		}

		var showProduct = model.ProductShow{
			DateCreate:      (*product)[i].CreatedAt.String(),
			Name:            (*product)[i].Name,
			ProductID:       (*product)[i].ID,
			Prince:          (*product)[i].Prince,
			Stock:           (*product)[i].Stock,
			Weight:          (*product)[i].Weight,
			Condition:       string((*product)[i].Condition),
			Brand:           (*product)[i].Brand,
			Description:     (*product)[i].Description,
			MinPurchase:     (*product)[i].MinPurchase,
			Sold:            (*product)[i].Sold,
			MitraID:         (*product)[i].MitraID,
			CatagoryID:      catagory.ID,
			SubcatagoryID:   (*product)[i].SubcatagoryID,
			MitraName:       mitra.StoreName,
			MitraCity:       mitra.City,
			MitraCityID:     mitra.CityID,
			CatagoryName:    catagory.CatagoryName,
			SubCatagoryName: subcatagory.SubCatagoryName,
			Rating:          fmt.Sprintf("%.1f", result),
			Images:          imgProducts,
		}
		showProducts = append(showProducts, showProduct)
		imgProducts = nil
	}
	return &showProducts, nil
}

func (e *ProductUsecaseImpl) AddAndUpdateProductShow(idMitra int, idProduct int, c *gin.Context) (*model.Product, []string, error) {
	tx := e.ongkirRepo.BeginTrans()
	var partImg []string
	var param = []string{"subcatagory_id", "minimal_purchase", "prince", "stock", "weight", "name", "condition", "brand", "description", "courier"}
	var dataStr []string
	var dataInt []int
	for i := 0; i < len(param); i++ {
		data := c.PostForm(param[i])
		if i < 5 {
			temp, err := strconv.Atoi(data)
			if err != nil {
				utils.ValidationRollbackImage(partImg)
				return nil, nil, fmt.Errorf("input has be number")
			}
			dataInt = append(dataInt, temp)
		} else {
			dataStr = append(dataStr, data)
		}
	}

	form, err := c.MultipartForm()
	files := form.File["images"]
	var counter = 0

	for _, file := range files {
		if counter >= 5 {
			utils.ValidationRollbackImage(partImg)
			return nil, nil, fmt.Errorf("maximum of 5 image files")
		}
		err = utils.FileImgValidation(file, partImg)
		if err != nil {
			return nil, nil, err
		}
		name := gotp.RandomSecret(12)
		path := viper.GetString("part.images") + name + ".jpg"
		err = c.SaveUploadedFile(file, path)
		if err != nil {
			fmt.Printf("[productUsecase.addDetailProduct] error save image %v \n", err)
			utils.ValidationRollbackImage(partImg)
			return nil, nil, fmt.Errorf("Oopss internal server error")
		}
		partImg = append(partImg, path)
		counter++
	}

	_, err = e.productRepo.SubCatagoryById(dataInt[0])
	if err != nil {
		utils.ValidationRollbackImage(partImg)
		return nil, nil, fmt.Errorf("subcatagory not found")
	}
	if dataInt[3] < 0 {
		utils.ValidationRollbackImage(partImg)
		return nil, nil, fmt.Errorf("stock must not be minus")
	}
	var product = model.Product{
		Name:          dataStr[0],
		Prince:        dataInt[2],
		Stock:         strconv.Itoa(dataInt[3]),
		Weight:        dataInt[4],
		Condition:     dataStr[1],
		Brand:         dataStr[2],
		Description:   dataStr[3],
		MinPurchase:   dataInt[1],
		MitraID:       uint(idMitra),
		SubcatagoryID: uint(dataInt[0]),
	}

	var newProduct = &model.Product{}
	if idProduct != 0 {
		newProduct, err = e.productRepo.UpdateProduct(idProduct, tx, &product)
		if err != nil {
			tx.Rollback()
			utils.ValidationRollbackImage(partImg)
			return nil, nil, err
		}
	} else {
		newProduct, err = e.productRepo.InsertProduct(&product, tx)
		if err != nil {
			tx.Rollback()
			utils.ValidationRollbackImage(partImg)
			return nil, nil, err
		}
	}

	var newImages []string
	if idProduct != 0 {
		productImg, err := e.productRepo.ImageByProductId(idProduct)
		if err != nil {
			utils.ValidationRollbackImage(partImg)
			return nil, nil, err
		}
		var partImg []string
		for i := 0; i < len(*productImg); i++ {
			partImg = append(partImg, (*productImg)[i].PartImg)
		}
		utils.ValidationRollbackImage(partImg)

		err = e.productRepo.DeleteProductImgByProductId(idProduct, tx)
		if err != nil {
			tx.Rollback()
			utils.ValidationRollbackImage(partImg)
			return nil, nil, err
		}
	}
	for y := 0; y < len(partImg); y++ {
		var productImg = model.ProductImg{
			ProductID: newProduct.ID,
			PartImg:   partImg[y],
		}
		newImg, err := e.productRepo.InsertProductImg(&productImg, tx)
		if err != nil {
			tx.Rollback()
			utils.ValidationRollbackImage(partImg)
			return nil, nil, err
		}
		newImages = append(newImages, newImg.PartImg)
	}
	tx.Commit()
	return newProduct, newImages, nil
}

func (e *ProductUsecaseImpl) DeleteProductShow(idProduct int, products *[]model.Product) error {
	if len(*products) == 0 {
		return fmt.Errorf("your product is empty")
	}
	tx := e.ongkirRepo.BeginTrans()
	for i := 0; i < len(*products); i++ {
		if int((*products)[i].ID) == idProduct {
			err := e.productRepo.DeleteProductById(idProduct, tx)
			if err != nil {
				tx.Rollback()
				return err
			}
			break
		} 
		if i+1 == len(*products) {
			return fmt.Errorf("failed delete product, product does not exist")
		}
	}
	productImg, err := e.productRepo.ImageByProductId(idProduct)
	if err != nil {
		tx.Rollback()
		return err
	}
	var partImg []string
	for j := 0; j < len(*productImg); j++ {
		partImg = append(partImg, (*productImg)[j].PartImg)
	}
	err = e.productRepo.DeleteProductImgByProductId(idProduct, tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = e.productRepo.ViewWachlistByProductId(idProduct)
	if err == nil {
		err := e.productRepo.DeleteWachlistByProductId(idProduct, tx)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	_, err = e.transactionRepo.ViewSubBasketByProductId(idProduct)
	if err == nil {
		err := e.transactionRepo.DeleteSubBasketByProductId(idProduct, tx)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	utils.ValidationRollbackImage(partImg)
	tx.Commit()
	return nil
}

func (e *ProductUsecaseImpl) InsertProductDiscussin(discussion *model.ProductDiscussion)(*model.ProductDiscussion, error) {
	return e.productRepo.InsertProductDiscussin(discussion)
}

func (e *ProductUsecaseImpl) ViewProductDiscussionByProductId(id int)(*[]model.ProductDiscussion, error) {
	return e.productRepo.ViewProductDiscussionByProductId(id)
}
