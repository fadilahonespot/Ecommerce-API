package config

import (
	"ecommerce/model"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

func DbConnect() *gorm.DB {
	DB, err := gorm.Open("mysql", viper.GetString("database.mysql"))
	if err != nil {
		log.Fatal(err)
	}

	DB.Debug().AutoMigrate(
		model.User{},
		model.UserAuth{},
		model.Newletter{},
		model.Catagory{},
		model.SubCatagory{},
		model.Mitra{},
		model.Product{},
		model.ProductReview{},
		model.ProductImg{},
		model.ProductDiscussion{},
		model.Courier{},
		model.CourierMitra{},
		model.Watchlist{},
		model.ProductSubDiscussion{},
		model.Basket{},
		model.SubBasket{},
		model.UserAddress{},
		// model.Transaction{},
		// model.MitraTransaction{},
		// model.SubTransaction{},
		model.Payment{},
		model.PaymentMethod{},
		// model.PaymentDetail{},
		// model.CourierTrx{},
	)

	DB.Model(&model.UserAuth{}).AddForeignKey("user_id", "user(id)", "CASCADE", "CASCADE")
	DB.Model(&model.Newletter{}).AddForeignKey("user_id", "user(id)", "CASCADE", "CASCADE")
	DB.Model(&model.Catagory{}).AddForeignKey("user_id", "user(id)", "CASCADE", "CASCADE")
	DB.Model(&model.SubCatagory{}).AddForeignKey("catagory_id", "catagory(id)", "CASCADE", "CASCADE")
	DB.Model(&model.Mitra{}).AddForeignKey("user_id", "user(id)", "CASCADE", "CASCADE")
	DB.Model(&model.SubCatagory{}).AddForeignKey("user_id", "user(id)", "CASCADE", "CASCADE")
	DB.Model(&model.Product{}).AddForeignKey("mitra_id", "mitra(id)", "CASCADE", "CASCADE")
	DB.Model(&model.Product{}).AddForeignKey("subcatagory_id", "subcatagory(id)", "CASCADE", "CASCADE")
	DB.Model(&model.ProductReview{}).AddForeignKey("user_id", "user(id)", "CASCADE", "CASCADE")
	DB.Model(&model.ProductReview{}).AddForeignKey("product_id", "product(id)", "CASCADE", "CASCADE")
	DB.Model(&model.ProductImg{}).AddForeignKey("product_id", "product(id)", "CASCADE", "CASCADE")
	DB.Model(&model.ProductDiscussion{}).AddForeignKey("product_id", "product(id)", "CASCADE", "CASCADE")
	DB.Model(&model.ProductDiscussion{}).AddForeignKey("user_id", "user(id)", "CASCADE", "CASCADE")
	DB.Model(&model.CourierMitra{}).AddForeignKey("courier_id", "courier(id)", "CASCADE", "CASCADE")
	DB.Model(&model.CourierMitra{}).AddForeignKey("mitra_id", "mitra(id)", "CASCADE", "CASCADE")
	DB.Model(&model.Watchlist{}).AddForeignKey("product_id", "product(id)", "CASCADE", "CASCADE")
	DB.Model(&model.Watchlist{}).AddForeignKey("user_id", "user(id)", "CASCADE", "CASCADE")
	DB.Model(&model.ProductSubDiscussion{}).AddForeignKey("product_discussion_id", "product_discussion(id)", "CASCADE", "CASCADE")
	DB.Model(&model.ProductSubDiscussion{}).AddForeignKey("user_id", "user(id)", "CASCADE", "CASCADE")
	DB.Model(&model.Basket{}).AddForeignKey("user_id", "user(id)", "CASCADE", "CASCADE")
	DB.Model(&model.Basket{}).AddForeignKey("mitra_id", "mitra(id)", "CASCADE", "CASCADE")
	DB.Model(&model.SubBasket{}).AddForeignKey("basket_id", "basket(id)", "CASCADE", "CASCADE")
	DB.Model(&model.SubBasket{}).AddForeignKey("product_id", "product(id)", "CASCADE", "CASCADE")
	DB.Model(&model.UserAddress{}).AddForeignKey("user_id", "user(id)", "CASCADE", "CASCADE")
	DB.Model(&model.Payment{}).AddForeignKey("payment_method_id", "payment_method(id)", "CASCADE", "CASCADE")
	// DB.Model(&model.Transaction{}).AddForeignKey("user_id", "user(id)", "CASCADE", "CASCADE")
	// DB.Model(&model.MitraTransaction{}).AddForeignKey("mitra_id", "mitra(id)", "CASCADE", "CASCADE")
	// DB.Model(&model.MitraTransaction{}).AddForeignKey("transaction_id", "transaction(id)", "CASCADE", "CASCADE")
	// DB.Model(&model.SubTransaction{}).AddForeignKey("mitra_transaction_id", "mitra_transaction(id)", "CASCADE", "CASCADE")
	// DB.Model(&model.SubTransaction{}).AddForeignKey("product_id", "product(id)", "CASCADE", "CASCADE")
	// DB.Model(&model.PaymentDetail{}).AddForeignKey("payment_id", "payment(id)", "CASCADE", "CASCADE")
	// DB.Model(&model.PaymentDetail{}).AddForeignKey("transaction_id", "transaction(id)", "CASCADE", "CASCADE")
	// DB.Model(&model.CourierTrx{}).AddForeignKey("mitra_transaction_id", "mitra_transaction(id)", "CASCADE", "CASCADE")
	// DB.Model(&model.CourierTrx{}).AddForeignKey("user_address_id", "user_address(id)", "CASCADE", "CASCADE")

	return DB
}
