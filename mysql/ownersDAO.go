package mysql

import (
	"encoding/json"
	"fmt"
	"github.com/DoHuy/parking_to_easy/model"
	"strconv"
)

type OwnerDAO interface {
	CreateNewOwner(newOwner interface{}) error
	FindOwnerById(id interface{}) (model.Owner, error)
	GetAllOwners(limit, offset string) ([]model.Owner, int, error)
	ChangeStatusOwner(data interface{}) error
}

func (db *DAO) CreateNewOwner(newOwner interface{}) error{
	var owner model.Owner
	raw,_ := json.Marshal(newOwner)
	err := json.Unmarshal(raw, &owner)
	fmt.Println("OWNER:::", owner)
	if err != nil {
		return err
	}
	err = db.connection.Table("owners").Create(&owner).Scan(&owner).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *DAO) FindOwnerById(id interface{}) (model.Owner, error) {
	var owner model.Owner
	var ownerId string
	raw,_ := json.Marshal(id)
	err := json.Unmarshal(raw, &ownerId)
	if err != nil {
		return model.Owner{}, err
	}
	err = db.connection.Table("owners").Raw("SELECT fullName, phoneNumber, address, cmndImage, created_at FROM owners WHERE credentialId=?", ownerId).Scan(&owner).Error
	if err != nil {
		if err.Error() == "record not found"{
			return model.Owner{}, err
		}
		return model.Owner{}, fmt.Errorf("Lỗi truy vấn database %s", err.Error())
	}
	return owner, nil
}

func (db *DAO)GetAllOwners(limit, offset string) ([]model.Owner, int, error) {
	var owners []model.Owner
	// get total
	type Total struct {
		Total string	`json:"total" `
	}
	var total Total
	err := db.connection.Table("owners").Raw("SELECT count(*) AS total FROM owners").Scan(&total).Error
	totalRecord, _ := strconv.Atoi(total.Total)
	limitNum, _ := strconv.Atoi(limit)
	fmt.Println("total:::", totalRecord)
	//owners
	err = db.connection.Table("owners").Raw("SELECT credentialId, fullName, phoneNumber, address, cmndImage, status, created_at FROM owners LIMIT ? OFFSET ?", limit, offset).Scan(&owners).Error
	if err != nil {
		return owners, 0, err
	}
	return owners, totalRecord/limitNum, nil
}

func (db *DAO)ChangeStatusOwner(data interface{}) error{
	raw, _ := json.Marshal(data)
	type DataStruct struct {
		ID			string		`json:"id"`
		Status		string		`json:"status"`
		ModifiedAt	string		`json:"modified_at"`

	}
	var updateData DataStruct
	err := json.Unmarshal(raw, &updateData)
	if err != nil {
		return err
	}
	err = db.connection.Exec("UPDATE owners SET `status`=?, modified_at=? WHERE id=?", updateData.Status, updateData.ModifiedAt,updateData.ID).Error
	if err != nil {
		return err
	}
	return nil
}