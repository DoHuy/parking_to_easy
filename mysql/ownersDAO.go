package mysql

import (
	"encoding/json"
	"fmt"
	"github.com/DoHuy/parking_to_easy/model"
)

type OwnerDAO interface {
	CreateNewOwner(newOwner interface{}) error
	FindOwnerById(id interface{}) (model.Owner, error)
	FindAllOwners() ([]model.Owner, error)
	ModifyOwner(data model.DataStruct) error
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

func (db *DAO)FindAllOwners() ([]model.Owner, error) {
	var owners []model.Owner
	//owners
	err := db.connection.Table("owners").Raw("SELECT credentialId, fullName, phoneNumber, address, cmndImage, status, created_at FROM owners").Scan(&owners).Error
	if err != nil {
		return owners, err
	}
	return owners, nil
}

func (db *DAO)ModifyOwner(updateData model.DataStruct) error{
	err := db.connection.Exec("UPDATE owners SET `status`=?, modified_at=? WHERE credentialId=?", updateData.Status, updateData.ModifiedAt,updateData.CredentialId).Error
	if err != nil {
		return err
	}
	return nil
}