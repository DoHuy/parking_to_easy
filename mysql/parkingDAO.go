package mysql

import (
	"encoding/json"
	"fmt"
	"github.com/DoHuy/parking_to_easy/model"
)

type ParkingDAO interface {
	CreateNewParkingOfOwner(newParking interface{})  error
	CreateNewParkingByAdmin(newParking interface{}) error
	ChangStatusParking(updatedParking model.Parking)  error
	FindParkingByID(id string) (model.Parking, error)
	FindParkingByOwnerId(ownerId string) (model.Owner, error)
	GetAllParking() ([]model.Parking, error)
}

func (db *DAO) CreateNewParkingOfOwner(newParking interface{})  error{
	var parking model.Parking
	raw,_ := json.Marshal(newParking)
	err := json.Unmarshal(raw, &parking)
	fmt.Println("PARKING:::", parking)
	if err != nil {
		return err
	}
	err = db.connection.Table("parkings").Create(&parking).Error
	if err != nil {
		fmt.Println("ERR in mysql:::", err)
		return err
	}
	fmt.Println("ERR in mysql:::", parking)
	return nil
}

func (db *DAO) CreateNewParkingByAdmin(newParking interface{}) error {
	var parking model.Parking
	raw,_ := json.Marshal(newParking)
	err := json.Unmarshal(raw, &parking)
	fmt.Println("Parking::::: ", parking)
	err  = db.connection.Table("parkings").Create(&parking).Error
	if err != nil {
		panic(err.Error())
		return err
	}
	return nil
}

func (db *DAO)ChangStatusParking(updatedParking model.Parking)  error {
	err := db.connection.Exec("UPDATE parkings SET `status`=?, modified_at=? WHERE id=?", updatedParking.Status,updatedParking.ModifiedAt, updatedParking.ID).Error
	if err != nil {
		fmt.Println("ERRR:::::", err)
		return err
	}
	return nil
}

func (db *DAO)FindParkingByID(id string) (model.Parking, error) {
	var parking model.Parking
	err := db.connection.Table("parkings").Raw("SELECT * FROM parkings WHERE parkings.id=?", id).Scan(&parking).Error

	if err != nil {
		return model.Parking{}, fmt.Errorf("Loi truy van database: %v", err.Error())
	}
	return parking, nil
}

func (db *DAO)FindParkingByOwnerId(ownerId string) (model.Owner, error) {
	var parkings []model.Parking
	var owner model.Owner

	err := db.connection.Raw("SELECT fullName, phoneNumber, address, cmndImage, status, created_at FROM owners WHERE credentialId=?", ownerId).Scan(&owner).Error
	if err != nil {
		return model.Owner{}, err
	}
	err  = db.connection.Raw("SELECT id, address, capacity, status FROM parkings WHERE ownerId=?", ownerId).Scan(&parkings).Error
	if err != nil {
		return model.Owner{}, err
	}
	owner.Parkings = parkings
	return owner, nil
}

func (db *DAO)GetAllParking() ([]model.Parking, error) {
	var parkings []model.Parking
	err := db.connection.Table("parkings").Raw("SELECT * FROM parkings").Scan(&parkings).Error
	if err != nil {
		return nil, fmt.Errorf("Loi truy van database: %v", err.Error())
	}
	if len(parkings) == 0 {
		return parkings, fmt.Errorf("records not found")
	}
	return parkings, nil
}