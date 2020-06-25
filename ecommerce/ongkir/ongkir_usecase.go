package ongkir

import "ecommerce/model"

type OngkirUsecase interface {
	GetCity()(*[]model.City, error)
	GetCityById(id int) (*model.City, error)
	GetProvinces()(*[]model.Province, error)
	CalculateShipping(shipping *model.QueryDetail)(*model.Shipping, error)
	GetCityByProvince(id int)(*[]model.City, error) 
	TrackResi(tracking *model.InputTracking) (*model.DetailTracking, error)
	ViewAllCourier()(*[]model.Courier, error)
	InsertCourier(courier *model.Courier)(*model.Courier, error)
	UpdateCourier(id int, courier *model.Courier) (*model.Courier, error)
	DeleteCourierById(id int) error
	ViewCourierById(id int)(*model.Courier, error)
	InsertCourierMitra(courierMitra *model.CourierMitra)(*model.CourierMitra, error)
	ViewCourierMitraById(id int) (*model.CourierMitra, error)
	ViewCourierMitraByMitraId(id int) (*[]model.CourierMitra, error)
	DeleteCourierMitraById(id int) error
	ViewCourierMitraShow(courierMitras *[]model.CourierMitra)(*[]model.CourierMitraShow, error)
}