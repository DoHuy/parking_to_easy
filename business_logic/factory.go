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
	return &DeviceService{
		Dao		: self.Dao,
		Redis	: self.Redis,
	}
}

func (self *ServiceFactory)GetOwnerService() *OwnerService{
	return &OwnerService{
		Dao		: self.Dao,
	}
}

func (self *ServiceFactory)GetParkingService() *ParkingService{
	return &ParkingService{
		Dao		: self.Dao,
	}
}

func (self *ServiceFactory)	GetRatingService() *RatingService{
	return &RatingService{
		Dao		: self.Dao,
	}
}

func (self *ServiceFactory)GetTransactionService() *TransactionService{
	return &TransactionService{
		Dao		: self.Dao,
		Redis	: self.Redis,
	}
}

func (self *ServiceFactory)GetAuthService() *auth.Auth{
	return &auth.Auth{
		Redis: self.Redis,
	}
}