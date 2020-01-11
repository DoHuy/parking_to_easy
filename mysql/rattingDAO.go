package mysql

import (
	"fmt"
	"github.com/DoHuy/parking_to_easy/model"
)

type RatingDAO interface {
	AverageStarsOfParking(id string) (float64, error)
	CreateRating(rating model.Rating) error
}

func (db *DAO)AverageStarsOfParking(id string) (float64, error) {
	type Data struct {
		Average		float64		`json:"average"`
	}
	var data Data
	err := db.connection.Raw("SELECT AVG(stars) AS average FROM ratings WHERE parkingId=?", id).Scan(&data).Error
	if err != nil {
		return 0.0, nil
	}
	fmt.Println("DATA:::",data)
	return data.Average, nil
}

func (db *DAO)CreateRating(rating model.Rating) error {
	err := db.connection.Create(&rating).Error
	if err != nil {
		return err
	}
	return nil
}