package business_logic

import (
	"fmt"
	"github.com/DoHuy/parking_to_easy/firebase"
	"github.com/DoHuy/parking_to_easy/model"
	"github.com/DoHuy/parking_to_easy/mysql"
	"github.com/DoHuy/parking_to_easy/redis"
	"github.com/DoHuy/parking_to_easy/utils"
)

type ParkingService struct {
	Dao 		*mysql.DAO
	Redis 		*redis.Redis
	FireBase	*firebase.FireBaseService
}

func NewParkingService(dao *mysql.DAO, redis *redis.Redis) *ParkingService{
	firebaseService := firebase.NewFireBaseService(redis)
	return &ParkingService{
		Dao: dao,
		Redis: redis,
		FireBase: firebaseService,
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
	// gửi thông báo về phía owner thông báo đã được admin approve
	var parking model.Parking
	parking, err = parkingIface.FindParkingByID(input.ID)
	err = self.FireBase.SendNotifyToUserOfParking(parking.ID, parking.OwnerId, "THÔNG BÁO", fmt.Sprintf("Điểm đậu xe %s đã được phê duyệt", parking.ParkingName))
	// remove parking topic in redis
	err = self.Redis.DeleteParkingTopic(parking.ID)
	if err != nil {
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

func (self *ParkingService)CheckExistedLocation(longitude, latitude string) bool {
	var parkingIface mysql.ParkingDAO
	parkingIface = self.Dao
	_, err := parkingIface.FindParkingByLongLat(longitude, latitude)
	if err != nil {
		fmt.Println("ERR ::::", err)
		return false
	}
	return true
}

