package mysql

import (
	"fmt"
	"github.com/DoHuy/parking_to_easy/config"
	"github.com/jinzhu/gorm"
	"github.com/DoHuy/parking_to_easy/model"
)
var connection *gorm.DB

func ConnectDatabase() (*gorm.DB, error) {
	config := config.GetDatabaseConfig()
	var err error
	connectionInfo := fmt.Sprintf(`%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local`, config.Database.Username, config.Database.Password, config.Database.Address, config.Database.DatabaseName)
	connection, err = gorm.Open("mysql", connectionInfo)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("error in connectDatabase(): %s", err.Error())
	}
	return connection, nil
}

// Thao tác với Credential model
func FindCredentialByID(id int) model.Credential {
	var credential model.Credential
	connection.Raw("SELECT username, password, email, points, role, token, expired, created_at, modified_at, deleted_at FROM credentials where id=?", id).Scan(&credential)
	return credential
}

func FindCredentialByNameAndPassword(name , pwd string) (model.Credential, error) {
	var credential model.Credential
	err := connection.Table("credentials").Raw("SELECT id, username, password FROM credentials where username=? AND password=?", name, pwd).Scan(&credential).Error
	if err != nil {
		return model.Credential{}, fmt.Errorf("Lỗi query db %v", err)
	}
	return credential, nil
}

func FindAllCredential(limit, offset int) []model.Credential {
	var credentials []model.Credential
	rows, _ := connection.Raw("SELECT username, password, email, points, role, token, expired, created_at, modified_at, deleted_at FROM credentials LIMIT ? OFFSET ?", limit, offset).Rows()
	defer rows.Close()
	index := 0
	for rows.Next() {
		rows.Scan(&credentials[index])
		index++
	}
	return credentials
}

func CreateCredential(newUser model.Credential) (error, model.Credential) {
	var credential model.Credential
	err := connection.Create(&newUser).Scan(&credential).Error
	err  = connection.Raw("SELECT username, email, points, role, created_at FROM credentials WHERE username=?", newUser.Username).Scan(&credential).Error
	if err != nil {
		return fmt.Errorf("Loi truy van database: %s", err.Error()), model.Credential{}
	}
	return nil, credential
}
////////////////////////////////////////////////////

// Thao tac vs Parking Model
func CreateOwnerAndParking(data interface{}) (map[string]interface{}, error) {
	var parking model.Parking
	var owner	model.Owner

	// init transaction
	tx := connection.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return nil, err
	}
	err := tx.Table("owners").Create(&data).Scan(&owner).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Table("parking").Create(&data).Scan(&parking)
	if err := tx.Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	result := map[string]interface{}{
		"message": "Tạo bãi xe thành công",
		"parking":parking,
		"owner": owner,
	}
	return result, nil

}

func CreateNewParkingByAdmin(newParking interface{}) (error, interface{}) {
	var parking model.Parking
	err := connection.Table("parkings").Create(&newParking).Scan(&parking).Error
	if err != nil {
		panic(err.Error())
		return err, nil
	}
	result := map[string]interface{}{
		"message":"Tạo bãi xe thành công",
		"parking": parking,
	}
	return nil, result
}

func ModifyParking(updatedParking model.Parking) model.Parking {
	var parking model.Parking
	connection.Model(&parking).Updates(updatedParking).Scan(&parking)
	return parking
}

func FindParkingByID(id string) ([]model.Parking, error) {
	var parkings []model.Parking
	err := connection.Table("parkings").Raw("SELECT * FROM parkings INNER JOIN owners ON owners.credentialId=parkings.ownerId WHERE parkings.id=?", id).Scan(&parkings).Error
	if err != nil {
		return nil, fmt.Errorf("Loi truy van database: %v", err.Error())
	}
	for i, _ := range parkings {
		err := connection.Table("parkings").Raw("SELECT*from owners WHERE credentialId=?", parkings[i].OwnerId).Scan(&parkings[i].Owner).Error

		if err != nil {
			return nil, fmt.Errorf("Loi truy van database: %v", err.Error())
		}
	}

	return parkings, nil
}
//
func GetAllParking(limit, offset string) ([]model.Parking, error){
	var parkings []model.Parking
	err := connection.Table("parkings").Raw("SELECT * FROM parkings LIMIT ? OFFSET ?", limit, offset).Scan(&parkings).Error
	if err != nil {
		return nil, fmt.Errorf("Loi truy van database: %v", err.Error())
	}
	for i, _ := range parkings {
		err := connection.Table("parkings").Raw("SELECT*from owners WHERE credentialId=?", parkings[i].OwnerId).Scan(&parkings[i].Owner).Error

		if err != nil {
			return nil, fmt.Errorf("Loi truy van database: %v", err.Error())
		}
	}
	return parkings, nil
}