package mysql

import (
	"fmt"
	"github.com/DoHuy/parking_to_easy/model"
)

type CredentialDAO interface {
	FindCredentialByName(username string) (model.Credential, error)
	FindCredentialByMail(email string) (model.Credential, error)
	FindCredentialByID(id string) (model.Credential, error)
	FindCredentialByNameAndPassword(name , pwd string) (model.Credential, error)
	FindAllCredential(limit, offset string) ([]model.Credential, error)
	CreateCredential(newUser model.Credential) error
	ModifyCredential(operator string, points int, credentialId int) error
}

func (db *DAO)FindCredentialByName(username string) (model.Credential, error){
	var credential model.Credential
	err := db.connection.Table("credentials").Raw("SELECT username, email FROM credentials WHERE username=?", username).
		Scan(&credential).Error
	if err != nil {
		return model.Credential{}, fmt.Errorf("Loi truy van co so du lieu %s", err.Error())
	}
	return credential, nil
}

func (db *DAO) FindCredentialByMail(email string) (model.Credential, error){
	var credential model.Credential
	err := db.connection.Table("credentials").Raw("SELECT username, email FROM credentials WHERE email=?", email).
		Scan(&credential).Error
	if err != nil {
		return model.Credential{}, fmt.Errorf("Loi truy van co so du lieu %s", err.Error())
	}
	return credential, nil
}

func (db *DAO) FindCredentialByID(id string) (model.Credential, error) {
	var credential model.Credential
	err := db.connection.Table("credentials").Raw("SELECT username, password, email, points, role, token, expired, created_at FROM credentials where id=?", id).Scan(&credential).Error
	if err != nil {
		if err.Error() == "record not found" {
			return credential, err
		}
		return credential, fmt.Errorf("Lỗi truy vấn db %s", err)
	}
	return credential, nil
}

func (db *DAO) FindCredentialByNameAndPassword(name , pwd string) (model.Credential, error) {
	var credential model.Credential
	err := db.connection.Table("credentials").Raw("SELECT id, username, password, role FROM credentials where username=? AND password=?", name, pwd).Scan(&credential).Error
	if err != nil {
		return model.Credential{}, fmt.Errorf("Lỗi query db %v", err)
	}
	return credential, nil
}

func (db *DAO) FindAllCredential(limit, offset string) ([]model.Credential, error) {
	var credentials []model.Credential
	err := db.connection.Raw("SELECT username, password, email, points, role, token, expired, created_at, modified_at, deleted_at FROM credentials LIMIT ? OFFSET ?", limit, offset).Scan(&credentials).Error
	if err != nil {
		fmt.Println("ERR: ", err)
		return credentials, err
	}
	return credentials, nil
}

func (db *DAO) CreateCredential(newUser model.Credential) error {
	err := db.connection.Create(&newUser).Error
	//err  = connection.Raw("SELECT username, email, points, role, created_at FROM credentials WHERE username=?", newUser.Username).Scan(&credential).Error
	if err != nil {
		return fmt.Errorf("Loi truy van database: %s", err.Error())
	}
	return nil
}

func (db *DAO)ModifyCredential(operator string, points int, credentialId int) error{
	var credential model.Credential
	err := db.connection.Raw("SELECT* FROM credentials WHERE id=?", credentialId).Scan(&credential).Error
	if operator == "SUB" {
		err := db.connection.Exec("UPDATE credentials SET points=? WHERE id=?", credential.Points - points, credentialId).Error
		if err != nil {
			return err
		}
	} else {
		//var owner model.Credential
		//err := db.connection.Raw("SELECT* FROM credentials WHERE id=?", credentialId).Scan(&owner).Error
		err  = db.connection.Exec("UPDATE credentials SET points=? WHERE id=?", credential.Points + points, credentialId).Error
		if err != nil {
			return err
		}
	}

	return nil
}