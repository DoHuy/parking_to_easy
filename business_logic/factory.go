package business_logic

import (
	"github.com/DoHuy/parking_to_easy/business_logic/auth"
	"github.com/DoHuy/parking_to_easy/mysql"
	"github.com/DoHuy/parking_to_easy/redis"
)

type ServiceFactory struct {
	Dao		*mysql.DAO
	Redis	*redis.Redis
}

func NewFactory() (*ServiceFactory, error){
	msql, err := mysql.NewDAO()
	if err != nil {
		return nil, err
	}
	redis := redis.NewRedis()
	return &ServiceFactory{
		Dao: 	msql,
		Redis:	redis,
	}, nil
}

func (self *ServiceFactory)GetDeViceService() *DeviceService{
	return NewDeviceService(self.Dao, self.Redis)
}

func (self *ServiceFactory)GetOwnerService() *OwnerService{
	return NewOwnerService(self.Dao, self.Redis)
}

func (self *ServiceFactory)GetParkingService() *ParkingService{
	return NewParkingService(self.Dao, self.Redis)

}

func (self *ServiceFactory)	GetRatingService() *RatingService{
	return NewRatingService(self.Dao)
}

func (self *ServiceFactory)GetTransactionService() *TransactionService{
	return NewService(self.Dao, self.Redis)
}

func (self *ServiceFactory)GetAuthService() *auth.Auth{
	return auth.NewAuth(self.Redis)
}

func (self *ServiceFactory)GetCustomerService() *CustomerService{
	return NewCustomerService(self.Dao)
}