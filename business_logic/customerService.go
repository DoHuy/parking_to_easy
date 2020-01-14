package business_logic

import (
	"fmt"
	"github.com/DoHuy/parking_to_easy/model"
	"github.com/DoHuy/parking_to_easy/mysql"
)

type CustomerService struct {
	Dao		*mysql.DAO
}

func NewCustomerService(dao *mysql.DAO)  *CustomerService{
	return &CustomerService{
		Dao: dao,
	}
}

func (this *CustomerService)SubPoints(credentialId, points int) error{
	var credentialIface mysql.CredentialDAO
	credentialIface = this.Dao
	err := credentialIface.ModifyCredential("SUB", points, credentialId)
	if err != nil {
		return err
	}
	return nil

}

func (this *CustomerService)CheckWallet(amount int, credentialId int) bool {
	var credential model.Credential
	var credentialIface mysql.CredentialDAO
	credentialIface = this.Dao
	credential, _ = credentialIface.FindCredentialByID(fmt.Sprintf("%d", credentialId))
	if credential.Points >= amount/1000 {
		return  true
	}
	return false
}