package mysql

import (
	"encoding/json"
	"fmt"
	"github.com/DoHuy/parking_to_easy/model"
	"github.com/DoHuy/parking_to_easy/utils"
)

type ParkingDAO interface {
	CreateNewParkingOfOwner(newParking model.Parking)  error
	CreateNewParkingByAdmin(newParking model.NewParkingByAdmin) error
	ChangStatusParking(input model.VerifyingParkingInput)  error
	FindParkingByID(id string) (model.Parking, error)
	FindParkingByOwnerId(ownerId string) (model.Owner, error)
	GetAllParking() ([]model.Parking, error)
	ModifyParking(data interface{}) error
	DeleteParking(data interface{}) error
	FindParkingByLongLat(longitude, latitude string)([]model.Transaction, error)
	FindParkingByCreatedAt(created string)(model.Parking, error)
}

func (db *DAO) CreateNewParkingOfOwner(newParking model.Parking)  error{
	var parking model.Parking
	err := utils.BindRawStructToRespStruct(newParking, &parking)
	err = db.connection.Table("parkings").Create(&parking).Error
	if err != nil {
		fmt.Println("ERR in mysql:::", err)
		return err
	}
	fmt.Println("ERR in mysql:::", parking)
	return nil
}

func (db *DAO) CreateNewParkingByAdmin(newParking model.NewParkingByAdmin) error {

	err := db.connection.Table("parkings").Create(&newParking).Error
	if err != nil {
		panic(err.Error())
		return err
	}
	return nil
}

func (db *DAO)ChangStatusParking(input model.VerifyingParkingInput)  error {
	err := db.connection.Exec("UPDATE parkings SET `status`=?, modified_at=? WHERE id=?", input.Status, input.ModifiedAt, input.ID).Error
	if err != nil {
		fmt.Println("ERRR:::::", err)
		return err
	}
	return nil
}

func (db *DAO)FindParkingByID(id string) (model.Parking, error) {
	var parking model.Parking
	err := db.connection.Raw("SELECT * FROM parkings WHERE parkings.id=? AND deleted_at=\"\"", id).Scan(&parking).Error

	if err != nil {
		fmt.Println("eRRERE", err)
		return model.Parking{}, err
	}
	fmt.Println("parking::: ::: :::", parking)
	return parking, nil
}

func (db *DAO)FindParkingByOwnerId(ownerId string) (model.Owner, error) {
	var parkings []model.Parking
	var owner model.Owner

	err := db.connection.Raw("SELECT credentialId, fullName, phoneNumber, address, cmndImage, status, created_at FROM owners WHERE credentialId=?", ownerId).Scan(&owner).Error
	if err != nil {
		return model.Owner{}, err
	}
	err  = db.connection.Raw("SELECT id, address, capacity, certificateOfland, status, blockAmount, `describe`  FROM parkings WHERE ownerId=? AND deleted_at=\"\"", ownerId).Scan(&parkings).Error
	if err != nil {
		return model.Owner{}, err
	}
	owner.Parkings = parkings
	return owner, nil
}

func (db *DAO)GetAllParking() ([]model.Parking, error) {
	var parkings []model.Parking
	err := db.connection.Table("parkings").Raw("SELECT * FROM parkings WHERE deleted_at=\"\"").Scan(&parkings).Error
	if err != nil {
		return nil, fmt.Errorf("Loi truy van database: %v", err.Error())
	}
	if len(parkings) == 0 {
		return parkings, fmt.Errorf("records not found")
	}
	//fmt.Println("sdasdsd", parkings)
	return parkings, nil
}

func (db *DAO)ModifyParking(data interface{}) error {
	type UpdatedParking struct {
		ID			string	`json:"id"`
		Capacity	string	`json:"capacity"`
		ModifiedAt	string	`json:"modified_at"`
	}
	var updatedParking UpdatedParking
	raw,_ := json.Marshal(data)
	err := json.Unmarshal(raw, &updatedParking)
	if err != nil {
		return err
	}
	err = db.connection.Exec("UPDATE parkings SET capacity=?, modified_at=? WHERE id=?", updatedParking.Capacity, updatedParking.ModifiedAt, updatedParking.ID).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *DAO)DeleteParking(data interface{}) error {
	type UpdatedParking struct {
		ID			string	`json:"id"`
		DeletedAt	string	`json:"deleted_at"`
	}
	var updatedParking UpdatedParking
	raw,_ := json.Marshal(data)
	err := json.Unmarshal(raw, &updatedParking)
	if err != nil {
		return err
	}
	err = db.connection.Exec("UPDATE parkings SET deleted_at=? WHERE id=?",updatedParking.DeletedAt, updatedParking.ID).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *DAO)FindParkingByLongLat(longitude, latitude string)([]model.Transaction, error) {
	var transactions []model.Transaction
	err := db.connection.Raw("SELECT* FROM parkings WHERE longitude=? AND latitude=?", longitude, latitude).Scan(&transactions).Error
	if err != nil {
		fmt.Println("erereereIn DAo", err)
		return []model.Transaction{}, err
	}
	return transactions, nil
}

func (db *DAO)FindParkingByCreatedAt(created string)(model.Parking, error){
	var parking model.Parking
	err := db.connection.Raw("SELECT* FROM parkings WHERE created_at=?", created).Scan(&parking).Error
	if err != nil {
		return model.Parking{}, err
	}
	return parking, nil
}