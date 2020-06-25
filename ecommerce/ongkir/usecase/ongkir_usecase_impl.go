package usecase

import (
	"ecommerce/model"
	"ecommerce/ongkir"
)

type OngkirUsecaseImpl struct {
	ongkirRepo ongkir.OngkirRepo
}

func CreateOngkirUsecase(ongkirRepo ongkir.OngkirRepo) ongkir.OngkirUsecase {
	return &OngkirUsecaseImpl{ongkirRepo}
}

func (e *OngkirUsecaseImpl) GetCity()(*[]model.City, error) {
	return e.ongkirRepo.GetCity()
}

func (e *OngkirUsecaseImpl) GetCityById(id int) (*model.City, error) {
	return e.ongkirRepo.GetCityById(id)
}

func (e *OngkirUsecaseImpl) GetProvinces()(*[]model.Province, error) {
	return e.ongkirRepo.GetProvinces()
}

func (e *OngkirUsecaseImpl) CalculateShipping(shipping *model.QueryDetail)(*model.Shipping, error) {
	return e.ongkirRepo.CalculateShipping(shipping)
}

func (e *OngkirUsecaseImpl) GetCityByProvince(id int)(*[]model.City, error) {
	return e.ongkirRepo.GetCityByProvince(id)
}

func (e *OngkirUsecaseImpl) TrackResi(tracking *model.InputTracking) (*model.DetailTracking, error) {
	return e.ongkirRepo.TrackResi(tracking)
}

func (e *OngkirUsecaseImpl) ViewAllCourier()(*[]model.Courier, error) {
	return e.ongkirRepo.ViewAllCourier()
}

func (e *OngkirUsecaseImpl) InsertCourier(courier *model.Courier)(*model.Courier, error) {
	tx := e.ongkirRepo.BeginTrans()
	courier, err := e.ongkirRepo.InsertCourier(courier, tx)
	if err != nil {
		return nil, err
	}
	tx.Commit()
	return courier, nil
}

func (e *OngkirUsecaseImpl) UpdateCourier(id int, courier *model.Courier) (*model.Courier, error) {
	tx := e.ongkirRepo.BeginTrans()
	courier, err := e.ongkirRepo.UpdateCourier(id, tx, courier)
	if err != nil {
		return nil, err
	}
	tx.Commit()
	return courier, nil
}

func (e *OngkirUsecaseImpl) DeleteCourierById(id int) error {
	return e.ongkirRepo.DeleteCourierById(id)
}

func (e * OngkirUsecaseImpl) ViewCourierById(id int)(*model.Courier, error) {
	return e.ongkirRepo.ViewCourierById(id)
}

func (e *OngkirUsecaseImpl) InsertCourierMitra(courierMitra *model.CourierMitra)(*model.CourierMitra, error) {
	tx := e.ongkirRepo.BeginTrans()
	courier, err := e.ongkirRepo.InsertCourierMitra(courierMitra, tx)
	if err != nil {
		return nil, err
	}
	tx.Commit()
	return courier, nil
}

func (e *OngkirUsecaseImpl) ViewCourierMitraById(id int) (*model.CourierMitra, error) {
	return e.ongkirRepo.ViewCourierMitraById(id)
}

func (e *OngkirUsecaseImpl) ViewCourierMitraByMitraId(id int) (*[]model.CourierMitra, error) {
	return e.ongkirRepo.ViewCourierMitraByMitraId(id)
}

func (e *OngkirUsecaseImpl) DeleteCourierMitraById(id int) error {
	return e.ongkirRepo.DeleteCourierMitraById(id)
}

func (e *OngkirUsecaseImpl) ViewCourierMitraShow(courierMitras *[]model.CourierMitra)(*[]model.CourierMitraShow, error) {
	var courierShow []model.CourierMitraShow
	for i := 0; i < len(*courierMitras); i++ {
		courier, err := e.ongkirRepo.ViewCourierById(int((*courierMitras)[i].CourierID))
		if err != nil {
			return nil, err
		}
		var temCourier = model.CourierMitraShow{
			MitraID: (*courierMitras)[i].MitraID,
			CourierID: courier.ID,
			CourierName: courier.CourierName,
		}
		courierShow = append(courierShow, temCourier)
	}
	return &courierShow, nil
}


