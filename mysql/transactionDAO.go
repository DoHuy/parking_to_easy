package mysql

import "fmt"

type TransactionDAO interface {
	CalTotalAmountOfParking(id string) (int, error)
}

func (db *DAO)CalTotalAmountOfParking(id string) (int, error) {
	type Data struct {
		Total	int		`json:"total"`
	}
	var data Data
	status := "ACCEPTED"
	err := db.connection.Raw("SELECT SUM(amount) AS total FROM transactions WHERE parkingId=? AND `status`=?", id, status).Scan(&data).Error
	if err != nil {
		return 0, nil
	}
	fmt.Println("DATA:::",data)
	return data.Total, nil
}
