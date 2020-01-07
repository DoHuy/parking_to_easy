package mysql

import (
	"encoding/json"
	"fmt"
	"github.com/DoHuy/parking_to_easy/model"
)

type TransactionDAO interface {
	CalTotalAmountOfParking(id string) (int, error)
	FindTransactionOfUser(id interface{}) ([]model.Transaction, error)
	FindAllTransaction() ([]model.Transaction, error)
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

func (db *DAO)FindTransactionOfUser(id interface{}) ([]model.Transaction, error) {
	raw, _ := json.Marshal(id)
	var userId int
	err := json.Unmarshal(raw, &userId)
	if err != nil {
		return []model.Transaction{}, err
	}
	var transactions []model.Transaction
	err = db.connection.Raw("SELECT * FROM transactions WHERE credentialId=?", userId).Scan(&transactions).Error
	if err != nil {
		return []model.Transaction{}, err
	}
	return transactions, nil

}

func (db *DAO)FindAllTransaction() ([]model.Transaction, error) {
	var transactions []model.Transaction
	err := db.connection.Raw("SELECT * FROM transactions").Scan(&transactions).Error
	if err != nil {
		return []model.Transaction{}, err
	}
	return transactions, nil
}