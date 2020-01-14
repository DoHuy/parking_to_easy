package business_logic

import (
	"encoding/json"
	"fmt"
	"github.com/DoHuy/parking_to_easy/firebase"
	"github.com/DoHuy/parking_to_easy/model"
	"github.com/DoHuy/parking_to_easy/mysql"
	"github.com/DoHuy/parking_to_easy/redis"
	"github.com/DoHuy/parking_to_easy/utils"
)

type OwnerService struct {
	Dao			*mysql.DAO
	Redis		*redis.Redis
	FireBase	*firebase.FireBaseService
}

func NewOwnerService(dao *mysql.DAO, redis *redis.Redis) *OwnerService{
	firebaseService := firebase.NewFireBaseService(redis)
	return &OwnerService{
		Dao: dao,
		Redis: redis,
		FireBase: firebaseService,
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

func (this *OwnerService)AddPoints(credentialId, points int) error{
	var credentialIface mysql.CredentialDAO
	credentialIface = this.Dao
	err := credentialIface.ModifyCredential("ADD", points, credentialId)
	if err != nil {
		return err
	}
	return nil

}

func (this *OwnerService)ShareParking(input interface{}) error{
	var parking model.Parking
	raw,_ := json.Marshal(input)
	err := json.Unmarshal(raw, &parking)
	fmt.Println("PARKING:::", parking)
	if err != nil {
		return err
	}
	var parkingIface mysql.ParkingDAO
	parkingIface = this.Dao
	err = parkingIface.CreateNewParkingOfOwner(parking)
	if err != nil {
		return err
	}
	// tao cau truc luu tru trong redis
	park, err := parkingIface.FindParkingByCreatedAt(parking.CreatedAt)
	if err != nil {
		return err
	}
	// lay toan bo token cua owner
	var userDeviceIface mysql.DeviceDAO
	userDeviceIface = this.Dao
	devices, _ := userDeviceIface.FindTokensOfUser(park.OwnerId)
	var tokens []string
	for _, device := range devices {
		tokens = append(tokens, device.DeviceToken)
	}
	// luu lai tokens trong redis cho viec gui thong bao
	err = this.Redis.SetTokenListParking(park.ID, park.OwnerId, tokens)
	if err != nil {
		return err
	}
	// gui thong bao toi admin co chia se diem dau moi
	//
	return nil
}