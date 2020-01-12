package business_logic

import (
	"fmt"
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

func (self *OwnerService)GetAllOwners(data interface{}) ([]model.GettingAllOwnersOutput, error) {
	var outputs []model.GettingAllOwnersOutput
	var ownerIface 		mysql.OwnerDAO
	var parkingIface	mysql.ParkingDAO
	var ratingIface		mysql.RatingDAO
	ownerIface = self.Dao
	parkingIface = self.Dao
	ratingIface = self.Dao
	owners, err := ownerIface.FindAllOwners()
	if err != nil {
		return []model.GettingAllOwnersOutput{}, err
	}
	//outputIndex := 0
	for i := 0 ; i< len(owners) ; i++ {
		var owner model.Owner
		owner, err = parkingIface.FindParkingByOwnerId(fmt.Sprintf("%d", owners[i].CredentialId))
		if err != nil {
			return []model.GettingAllOwnersOutput{}, err
		}
		var votes int
		var stars float64
		if len(owner.Parkings) == 0 || owner.Parkings == nil {
			votes = 0
			stars = 0

		} else {
			var avgs float64
			for j := 0 ; j < len(owner.Parkings) ; j++ {
				vote, _ 	:= ratingIface.CountVote(fmt.Sprintf("%d", owner.Parkings[j].ID))
				avg, _ 	:= ratingIface.AverageStarsOfParking(fmt.Sprintf("%d", owner.Parkings[j].ID))
				votes +=vote
				avgs +=avg
			}
			stars = avgs/float64(len(owner.Parkings))

		}
		outputs = append(outputs, model.GettingAllOwnersOutput{Owner: owner, Stars: stars, Votes: votes})
	}

	return outputs, nil
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