package repo

import (
	"database/sql"
	"ecommerce/model"
	"fmt"
)

func (e *ProductRepoImpl) ViewAllCatagory()(*[]model.Catagory, error) {
	var catagory  []model.Catagory
	err := e.DB.Find(&catagory).Error
	if err != nil {
		fmt.Printf("[ProductRepoImpl.ViewAllCatagory] error execute query %v \n", err)
		return nil, fmt.Errorf("failed view data all catagory")
	}
	return &catagory, nil
}

func (e *ProductRepoImpl) InsertCatagory(catagory *model.Catagory) (*model.Catagory, error) {
	err := e.DB.Save(&catagory).Error
	if err != nil {
		fmt.Printf("[ProductRepoImpl.InsertCatagory] error execute query %v \n", err)
		return nil, fmt.Errorf("failed insert data catagory, id catagory is not exsis")
	}
	return catagory, nil
}

func (e *ProductRepoImpl) ViewCatagoryByIdUser(id int)(*model.Catagory, error) {
	var catagory = model.Catagory{}
	err := e.DB.Table("catagory").Where("user_id = ?", id).First(&catagory).Error
	if err != nil {
		fmt.Printf("[ProductRepoImpl.ViewCatagoryByIdUser] error execute query %v \n", err)
		return nil, fmt.Errorf("Id user is no exist in catagory")
	}
	return &catagory, nil
}

func (e *ProductRepoImpl) InsertSubCatagory(subcatagory *model.SubCatagory)(*model.SubCatagory, error) { 
	err := e.DB.Save(&subcatagory).Error
	if err != nil {
		fmt.Printf("[ProductRepoImpl.InsertSubCatagory] error execute query %v \n", err)
		return nil, fmt.Errorf("failed insert data sub catagory")
	}
	return subcatagory, nil
}

func (e *ProductRepoImpl) ViewAllSubCatagory()(*[]model.SubCatagory, error) {
	var subCatagory []model.SubCatagory
	err := e.DB.Find(&subCatagory).Error
	if err != nil {
		fmt.Printf("[ProductRepoImpl.ViewAllSubCatagory] error execute query %v \n", err)
		return nil, fmt.Errorf("failed view data sub catagory")
	}
	return &subCatagory, nil
}

func (e *ProductRepoImpl) ViewSubPlusCatagory(idOrNil int)(*[]model.SubPlusCatagory, error) {
	var catagories []model.SubPlusCatagory
	var rows *sql.Rows
	selectItem := "catagory.id, subcatagory.id, catagory.catagory_name, subcatagory.sub_catagory_name"
	query := e.DB.Table("catagory").Select(selectItem).Joins("join subcatagory on subcatagory.catagory_id = catagory.id")
	if idOrNil != 0 {
		rows, _ = query.Where("catagory.id = ?", idOrNil).Rows()
	} else {
		rows, _ = query.Rows()
	}
	defer rows.Close()
	var catagory = model.SubPlusCatagory{}
	for rows.Next() {
		rows.Scan(&catagory.CatagoryID, &catagory.SubCatagoryID, &catagory.CatagoryName, &catagory.SubCatagoryName)
		catagories = append(catagories, catagory)
	}
	if catagories == nil {
		fmt.Printf("[productRepoImpl.ViewSubPlusCatagory] error execute join query")
		return nil, fmt.Errorf("catagory data is empty")
	}
	return &catagories, nil
}

func (e *ProductRepoImpl) SubCatagoryById(id int)(*model.SubCatagory, error) {
	var subcatagory = model.SubCatagory{}
	err := e.DB.Table("subcatagory").Where("id = ?", id).First(&subcatagory).Error
	if err != nil {
		fmt.Printf("[productRepoImpl.SubCatagoryById] error execute query %v \n", err)
		return nil, fmt.Errorf("id sub catagory is not exist")
	}
	return &subcatagory, nil
}

func (e *ProductRepoImpl) CatagoryById(id int) (*model.Catagory, error) {
	var catagory = model.Catagory{}
	err := e.DB.Table("catagory").Where("id = ?", id).Find(&catagory).Error
	if err != nil {
		fmt.Printf("[ProductRepoImpl.CatagoryById] error execute query, %v \n", err)
		return nil, fmt.Errorf("id catagory is not exist")
	}
	return &catagory, nil
}

func (e *ProductRepoImpl) UpdateCatagoryById(id int, catagory *model.Catagory)(*model.Catagory, error) {
	var upCatagory = model.Catagory{}
	err := e.DB.Table("catagory").Where("id = ?", id).First(&upCatagory).Update(&catagory).Error
	if err != nil {
		fmt.Printf("[ProductRepo.UpdateCatagoryById] error execute query %v \n", err)
		return nil, fmt.Errorf("failed update catagory, id catagor not exsis")
	}
	return &upCatagory, nil
}

func (e *ProductRepoImpl) UpdateSubCatagory(id int, subcatagory *model.SubCatagory)(*model.SubCatagory, error) {
	var upSubcatagory = model.SubCatagory{}
	err := e.DB.Table("subcatagory").Where("id = ?", id).First(&upSubcatagory).Update(&subcatagory).Error
	if err != nil {
		fmt.Printf("[ProductRepoImpl.UpdateSubCatagory] error execute query %v \n", err)
		return nil, fmt.Errorf("failed update subcatagory, id subcatagory is not exist")
	}
	return &upSubcatagory, nil
}