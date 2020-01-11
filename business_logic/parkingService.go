package business_logic

import (
	"github.com/DoHuy/parking_to_easy/model"
	"github.com/DoHuy/parking_to_easy/mysql"
	"github.com/DoHuy/parking_to_easy/utils"
)

type ParkingService struct {
	Dao *mysql.DAO
}

func NewParkingService(dao *mysql.DAO) *ParkingService{
	return &ParkingService{
		Dao: dao,
	}
}

func (self *ParkingService)CalculateAmountParkingAndVote(id string)(model.CalculateAmountParkingResp, error){
	var ratingIface mysql.RatingDAO
	ratingIface = self.Dao
	avg, err := ratingIface.AverageStarsOfParking(id)
	if err != nil {
		return model.CalculateAmountParkingResp{}, err
	}
	var transactionIface mysql.TransactionDAO
	transactionIface = self.Dao
	points, err := transactionIface.CalTotalAmountOfParking(id)
	if err != nil {
		return model.CalculateAmountParkingResp{}, err
	}
	return model.CalculateAmountParkingResp{Points: points, Stars: avg}, nil
}

func (self *ParkingService)VerifyParking(updatedData interface{}) error {
	var input model.VerifyingParkingInput
	var parkingIface mysql.ParkingDAO
	err := utils.BindRawStructToRespStruct(updatedData, &input)
	if err != nil {
		return err
	}
	parkingIface = self.Dao
	if err := parkingIface.ChangStatusParking(input); err != nil {
		return err
	}
	return nil
}

func (self *ParkingService)CheckExistedParking(id string) bool {
	var parkingIface mysql.ParkingDAO
	parkingIface = self.Dao
	_, err := parkingIface.FindParkingByID(id)
	if err != nil {
		return false
	}
	return true
}
