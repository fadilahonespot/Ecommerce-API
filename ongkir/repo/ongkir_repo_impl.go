package repo

import (
	"ecommerce/model"
	"ecommerce/ongkir"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

type OngkirRepoImpl struct {
	DB *gorm.DB
}

func CreateOngkirRepo(DB *gorm.DB) ongkir.OngkirRepo {
	return &OngkirRepoImpl{DB}
}

func (e *OngkirRepoImpl) BeginTrans() *gorm.DB {
	return e.DB.Begin()
}

func (r *OngkirRepoImpl) dataProcess(req *http.Request)([]byte, error) {
	req.Header.Add(viper.GetString("rajaongkir.key"), viper.GetString("rajaongkir.token"))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("[OngkirRepoImpl.dataProcess] Error http default client %v \n", err)
		return nil, fmt.Errorf("failed process data")
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("[OngkirRepoImpl.dataProcess] Error ioutil res Body %v \n", err)
		return nil, fmt.Errorf("Failed prosess data")
	}
	return body, nil
}

func (r *OngkirRepoImpl) GetCity()(*[]model.City, error) {
	req, err := http.NewRequest("GET", viper.GetString("rajaongkir.url") + "city", nil)
	if err != nil {
		fmt.Printf("[OngkirRepoImpl.getCity] Error execute url, %v \n", err)
		return nil, fmt.Errorf("Failed show data city")
	}
	
	body, err := r.dataProcess(req)
	if err != nil {
		return nil, err
	}

	var city = model.GetCity{}
	err = json.Unmarshal(body, &city)
	if err != nil {
		fmt.Printf("[OngkirRepoImpl.getCity] Error unmarshal data city %v \n", err)
		return nil, fmt.Errorf("failed show data city")
	}
	return &city.RajaOngkir.Result, nil
}

func (r *OngkirRepoImpl) GetCityById(id int) (*model.City, error) {
	req, err := http.NewRequest("GET", viper.GetString("rajaongkir.url") + "city?id=" + strconv.Itoa(id), nil)
	if err != nil {
		fmt.Printf("[OngkirRepoImpl.getById] Error execute url, %v \n", err)
		return nil, fmt.Errorf("id city is not found/exist")
	}

	body, err := r.dataProcess(req)
	if err != nil {
		return nil, err
	}

	var city = model.GetCityByID{}
	err = json.Unmarshal(body, &city)
	if err != nil {
		fmt.Printf("[OngkirRepoImpl.getById] Error unmarshal data city %v \n", err)
		return nil, fmt.Errorf("id is not exist")
	}
	return &city.RajaOngkir.Result, nil
}

func (e *OngkirRepoImpl) GetProvinces()(*[]model.Province, error) {
	req, err := http.NewRequest("GET", viper.GetString("rajaongkir.url") + "province", nil)
	if err != nil {
		fmt.Printf("[OngkirRepoImpl.GetProvince] Error execute query %v \n", err)
		return nil, fmt.Errorf("failed view all data province")
	}

	body, err := e.dataProcess(req)
	if err != nil {
		return nil, err
	}

	var provinces = model.Provinces{}
	err = json.Unmarshal(body, &provinces)
	if err != nil {
		fmt.Printf("[OngkirRepoImpl.GetProvinces] Error unmarshal data json %v \n", err)
		return nil, fmt.Errorf("failed view all data provinces")
	}
	return &provinces.RajaOngkir.Result, nil
}

func (e *OngkirRepoImpl) GetCityByProvince(id int)(*[]model.City, error) {
	req, err := http.NewRequest("GET", viper.GetString("rajaongkir.url") + "city?province=" + strconv.Itoa(id), nil)
	if err != nil {
		fmt.Printf("[OngkirRepoImpl.getById] Error execute url, %v \n", err)
		return nil, fmt.Errorf("id is not exist")
	}

	body, err := e.dataProcess(req)
	if err != nil {
		return nil, err
	}

	var city = model.GetCityByProvince{}
	err = json.Unmarshal(body, &city)
	if err != nil {
		fmt.Printf("[OngkirRepoImpl.getById] Error unmarshal data city %v \n", err)
		return nil, fmt.Errorf("id is not exist")
	}
	return &city.RajaOngkir.Result, nil
}

func (e *OngkirRepoImpl) CalculateShipping(shipping *model.QueryDetail)(*model.Shipping, error) {
	payload := strings.NewReader("origin=" + shipping.Origin + "&destination=" + shipping.Destination + "&weight=" + strconv.Itoa(shipping.Weight) +"&courier=" + shipping.Courier)
	req, err := http.NewRequest("POST", viper.GetString("rajaongkir.url") + "cost", payload)
	if err != nil {
		fmt.Printf("[OngkirRepoImpl.CalculateShipping] Error execute query %v \n", err)
		return nil, fmt.Errorf("failed view all data province")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	body, err := e.dataProcess(req)
	if err != nil {
		return nil, err
	}

	var reqDetailShipping = model.DetailShippingMap{}
	err = json.Unmarshal(body, &reqDetailShipping)
	if err != nil {
		fmt.Printf("[OngkirRepoImpl.CalculateShipping] Error json unmarshal %v \n", err)
		return nil, fmt.Errorf("Oopss internal server error")
	}

	var detailShipping = model.Shipping{
		OriginDetail: reqDetailShipping.RajaOngkir.OriginDetail,
		DestinationDetail: reqDetailShipping.RajaOngkir.DestinationDetail,
		Result: reqDetailShipping.RajaOngkir.Result,
	}
	
	return &detailShipping, nil
}

func (e *OngkirRepoImpl) TrackResi(tracking *model.InputTracking) (*model.DetailTracking, error) {
	req, err := http.NewRequest("GET", viper.GetString("binderbyte.url") + "cekresi?awb=" + tracking.AWB +"&api_key="+ viper.GetString("binderbyte.token") +"&courier=" + tracking.Courier, nil)
	if err != nil {
		fmt.Printf("[OngkirRepo.TrackingResi] Error request method %v \n", err)
		return nil, fmt.Errorf("Receipt number not found")
	}
	body, err := e.dataProcess(req)
	if err != nil {
		return nil, err
	}
	var packageTracking = model.TrackingPackage{}
	err = json.Unmarshal(body, &packageTracking)
	if err != nil {
		fmt.Printf("[OngkirRepo.TracResi] Error unmarshal json body %v \n", err)
		return nil, fmt.Errorf("Failed tracking resi")
	}
	return &packageTracking.Data, nil
}

func (e *OngkirRepoImpl) ViewAllCourier()(*[]model.Courier, error) {
	var corier []model.Courier
	err := e.DB.Table("courier").Find(&corier).Error
	if err != nil {
		fmt.Printf("[OngkirRepoImpl] error execute query %v \n", err)
		return nil, fmt.Errorf("id is not exist")
	}
	return &corier, nil
}

func (e *OngkirRepoImpl) ViewCourierById(id int)(*model.Courier, error) {
	var courier = model.Courier{}
	err := e.DB.Table("courier").Where("id = ?", id).First(&courier).Error
	if err != nil {
		fmt.Printf("[OngkirRepoImpl.ViewCourierById] error execute query %v \n", err)
		return nil, fmt.Errorf("id is not exsis")
	}
	return &courier, nil
}

func (e *OngkirRepoImpl) InsertCourier(courier *model.Courier, tx *gorm.DB)(*model.Courier, error) {
	err := tx.Save(&courier).Error
	if err != nil {
		fmt.Println("[OngkirRepoImpl.InsertCourier] error execute query %v \n", err)
		return nil, fmt.Errorf("failed insert data courier")
	}
	return courier, nil
}

func (e *OngkirRepoImpl) UpdateCourier(id int, tx *gorm.DB, courier *model.Courier)(*model.Courier, error) {
	var upCourier = model.Courier{}
	err := tx.Table("courier").Where("id = ?", id).First(&upCourier).Update(&courier).Error
	if err != nil {
		fmt.Printf("[OngkirRepoImpl.UpdateCourier] error execute query %v \n", err)
		return nil, fmt.Errorf("failed update data courier")
	}
	return &upCourier, nil
}

func (e *OngkirRepoImpl) DeleteCourierById(id int) error {
	var courier = model.Courier{}
	err := e.DB.Table("courier").Where("id = ?", id).Delete(&courier).Error
	if err != nil {
		fmt.Printf("[OngkirRepoImpl] error execute query %v \n", err)
		return fmt.Errorf("failed delete data, id is not exsis")
	}
	return nil
} 

func (e *OngkirRepoImpl) InsertCourierMitra(courierMitra *model.CourierMitra, tx *gorm.DB)(*model.CourierMitra, error) {
	err := tx.Save(&courierMitra).Error
	if err != nil {
		fmt.Printf("[OngkirRepoImpl.InsertCourierMitra] error execute query %v \n", err)
		return nil, fmt.Errorf("failed insert data courier mitra")
	}
	return courierMitra, nil
}

func (e *OngkirRepoImpl) ViewCourierMitraById(id int) (*model.CourierMitra, error) {
	var courier = model.CourierMitra{}
	err := e.DB.Table("courier_mitra").Where("id = ?", id).First(&courier).Error
	if err != nil {
		fmt.Printf("[OngkirRepoImpl.ViewCourierMitraById] error execute query %v \n", err)
		return nil, fmt.Errorf("id courier is not exsis")
	}
	return &courier, nil
}

func (e *OngkirRepoImpl) ViewCourierMitraByMitraId(id int) (*[]model.CourierMitra, error) {
	var couriers []model.CourierMitra
	err := e.DB.Table("courier_mitra").Where("mitra_id = ?", id).Find(&couriers).Error
	if err != nil {
		fmt.Printf("[OngkirRepoImpl.ViewCourierByMitraId] error execute query %v \n", err)
		return nil, fmt.Errorf("failed fiew courier mitra, id mitra is not exsis")
	}
	return &couriers, nil
}

func (e OngkirRepoImpl) DeleteCourierMitraById(id int) error {
	var courier = model.CourierMitra{}
	err := e.DB.Table("courier_mitra").Where("id = ?", id).Delete(&courier).Error
	if err != nil {
		fmt.Printf("[OngkirRepoImpl.DeleteCourierMitraById] error execute query %v \n", err)
		return fmt.Errorf("failed delete courier mitra, id is not exsis")
	}
	return nil
}



