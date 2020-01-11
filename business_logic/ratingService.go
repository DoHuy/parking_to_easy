package business_logic

import (
	"github.com/DoHuy/parking_to_easy/model"
	"github.com/DoHuy/parking_to_easy/mysql"
	"github.com/DoHuy/parking_to_easy/utils"
)

type RatingService struct {
	Dao *mysql.DAO
}

func NewRatingService(dao *mysql.DAO) *RatingService{
	return &RatingService{
		Dao: dao,
	}
}

func (self *RatingService)CreateVoteOfUser(rawRating interface{}) error{
	var rating model.Rating
	err := utils.BindRawStructToRespStruct(rawRating, &rating)
	if err != nil {
		return err
	}
	var ratingIface mysql.RatingDAO
	ratingIface = self.Dao
	if err := ratingIface.CreateRating(rating); err != nil {
		return err
	}
	return nil
}
