package business_logic

import (
	"github.com/DoHuy/parking_to_easy/model"
	"github.com/DoHuy/parking_to_easy/mysql"
	"github.com/DoHuy/parking_to_easy/utils"
	redis "github.com/DoHuy/parking_to_easy/redis"
)

type DeviceService struct {
	Dao		*mysql.DAO
	Redis	*redis.Redis
}

func NewDeviceService(dao *mysql.DAO) *DeviceService{
	var redisPool *redis.Redis
	redisPool = redis.NewRedis()
	return &DeviceService{
		Dao:dao,
		Redis:redisPool,
	}
}

func (self *DeviceService)SaveDeviceTokenOfUser(input interface{}) error{
	var devToken model.UserDevice
	err := utils.BindRawStructToRespStruct(input, &devToken)
	if err != nil {
		return err
	}
	var deviceIface mysql.DeviceDAO
	deviceIface = self.Dao
	existed, err := deviceIface.FindToken(devToken)
	if  err != nil && err.Error() == "record not found" {
		if err = deviceIface.CreateNewToken(devToken); err != nil {
			return err
		}
	}
	if existed.ID != 0 {
		if err := deviceIface.DeleteToken(existed); err != nil {
			return err
		}

		if err = deviceIface.CreateNewToken(devToken); err != nil {
			return err
		}

	}
	return nil
}

func (self *DeviceService)RemoveDeviceTokenOfUser(input interface{}) error{
	var devToken model.UserDevice
	err := utils.BindRawStructToRespStruct(input, &devToken)
	if err != nil {
		return err
	}
	var deviceIface mysql.DeviceDAO
	deviceIface = self.Dao
	if err = deviceIface.DeleteToken(devToken); err != nil {
		return err
	}
	return nil
}

func (self *DeviceService)GetTokenListOfUser(userId int) ([]model.UserDevice, error){
	var deviceIface mysql.DeviceDAO
	deviceIface = self.Dao
	list, err := deviceIface.FindTokensOfUser(userId)
	if err != nil {
		return []model.UserDevice{}, err
	}
	return list, nil
}

func (self *DeviceService)CreateTransactionTopic(input model.TransactionTopicInput) error {
	return nil
}
