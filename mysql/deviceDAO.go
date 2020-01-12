package mysql

import "github.com/DoHuy/parking_to_easy/model"

type DeviceDAO interface {
	CreateNewToken(input model.UserDevice) error
	DeleteToken(input model.UserDevice) error
	FindToken(input model.UserDevice) (model.UserDevice, error)
	FindTokensOfUser(userId int) ([]model.UserDevice, error)
}

func (self *DAO)CreateNewToken(input model.UserDevice) error {
	err := self.connection.Exec("INSERT INTO userDevices (deviceToken, credentialId) VALUES(?, ?)", input.DeviceToken, input.CredentialId).Error
	if err != nil {
		return err
	}
	return nil
}

func (self *DAO)DeleteToken(input model.UserDevice) error {
	err := self.connection.Exec("DELETE FROM userDevices WHERE deviceToken=? AND credentialId=?", input.DeviceToken, input.CredentialId).Error
	if err != nil {
		return err
	}
	return nil
}

func  (self *DAO)FindToken(input model.UserDevice) (model.UserDevice, error) {
	var output model.UserDevice
	err := self.connection.Raw("SELECT* FROM userDevices WHERE deviceToken=? AND credentialId=?", input.DeviceToken, input.CredentialId).Scan(&output).Error
	if err != nil {
		return model.UserDevice{}, err
	}
	return output, nil
}
func (self *DAO)FindTokensOfUser(userId int) ([]model.UserDevice, error){
	var outputs []model.UserDevice
	err := self.connection.Raw("SELECT* FROM userDevices WHERE credentialId=?", userId).Scan(&outputs).Error
	if err != nil {
		return []model.UserDevice{}, err
	}
	return outputs, nil
}