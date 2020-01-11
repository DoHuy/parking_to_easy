package business_logic

import (
	"github.com/DoHuy/parking_to_easy/model"
	"github.com/DoHuy/parking_to_easy/mysql"
	"github.com/DoHuy/parking_to_easy/utils"
)

type OwnerService struct {
	Dao	*mysql.DAO
}

func NewOwnerService(dao *mysql.DAO) *OwnerService{
	return &OwnerService{
		Dao: dao,
	}
}

func (self *OwnerService)GetAllOwners(data interface{}) ([]model.Owner, error) {
	var ownerIface mysql.OwnerDAO
	ownerIface = self.Dao
	owners, err := ownerIface.FindAllOwners()
	if err != nil {
		return []model.Owner{}, err
	}
	return owners, nil
}

func (self *OwnerService)DisableOwner(data interface{}) error{
	var input model.DataStruct
	err := utils.BindRawStructToRespStruct(data, &input)
	if err != nil {
		return err
	}
	var ownerIface mysql.OwnerDAO
	ownerIface = self.Dao
	if err := ownerIface.ModifyOwner(input); err != nil {
		return err
	}
	return nil
}