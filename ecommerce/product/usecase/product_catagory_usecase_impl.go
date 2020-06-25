package usecase

import "ecommerce/model"

func (e *ProductUsecaseImpl) ViewAllCatagory() (*[]model.Catagory, error) {
	return e.productRepo.ViewAllCatagory()
}

func (e *ProductUsecaseImpl) InsertCatagory(catagory *model.Catagory) (*model.Catagory, error) {
	return e.productRepo.InsertCatagory(catagory)
}

func (e *ProductUsecaseImpl) ViewCatagoryByIdUser(id int) (*model.Catagory, error) {
	return e.productRepo.ViewCatagoryByIdUser(id)
}

func (e *ProductUsecaseImpl) ViewAllSubCatagory() (*[]model.SubCatagory, error) {
	return e.productRepo.ViewAllSubCatagory()
}

func (e *ProductUsecaseImpl) InsertSubCatagory(subcatagory *model.SubCatagory) (*model.SubCatagory, error) {
	return e.productRepo.InsertSubCatagory(subcatagory)
}

func (e *ProductUsecaseImpl) ViewSubPlusCatagory(idOrNil int) (*[]model.SubPlusCatagory, error) {
	return e.productRepo.ViewSubPlusCatagory(idOrNil)
}

func (e *ProductUsecaseImpl) SubCatagoryById(id int) (*model.SubCatagory, error) {
	return e.productRepo.SubCatagoryById(id)
}

func (e *ProductUsecaseImpl) UpdateCatagoryById(id int, catagory *model.Catagory)(*model.Catagory, error) {
	return e.productRepo.UpdateCatagoryById(id, catagory)
}

func (e *ProductUsecaseImpl) UpdateSubCatagory(id int, subcatagory *model.SubCatagory)(*model.SubCatagory, error) {
	return e.productRepo.UpdateSubCatagory(id, subcatagory)
}
