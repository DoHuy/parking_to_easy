package mysql

import "fmt"

type RatingDAO interface {
	AverageStarsOfParking(id string) (float64, error)
	CreateVoteOfUser(data interface{}) error
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

func (db *DAO)CreateVoteOfUser(data interface{}) error {
	err := db.connection.Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}