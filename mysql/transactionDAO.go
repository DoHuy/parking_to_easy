package mysql

import (
	"fmt"
	"github.com/DoHuy/parking_to_easy/model"
)

type TransactionDAO interface {
	CalTotalAmountOfParking(id string) (int, error)
	FindTransactionOfUser(id int) ([]model.Transaction, error)
	FindAllTransaction() ([]model.Transaction, error)
	CreateTransaction(tran model.Transaction) error
	FindAllTransactionOfUser(userId, status int) ([]model.GettingTransactionDetailResp, error)
	FindAllTransactionOfOwner(ownerId, status int)([]model.GettingTransactionDetailResp, error)
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

func (db *DAO)FindTransactionOfUser(id int) ([]model.Transaction, error) {
	var transactions []model.Transaction
	err := db.connection.Raw("SELECT * FROM transactions WHERE credentialId=?", id).Scan(&transactions).Error
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

func (db *DAO)CreateTransaction(transaction model.Transaction) error{
	err := db.connection.Create(&transaction).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *DAO)FindAllTransactionOfUser(userId, status int) ([]model.GettingTransactionDetailResp, error){
	var output []model.GettingTransactionDetailResp
	sql := `SELECT transactions.id as transactionId, transactions.startTime, transactions.endTime, transactions.licence,
			parkings.address as address, transactions.amount, transactions.status, transactions.created_at,
			transactions.phoneNumber as userPhoneNumber, owners.phoneNumber as hostPhoneNumber, transactions.parkingId
			FROM transactions as tran
			INNER JOIN parkings as p
			ON tran.parkingId = p.id
			INNER JOIN owners as o
			ON p.ownerId = o.credentialId
			WHERE transactions.credentialId=? AND transactions.status=?`
	err := db.connection.Raw(sql, userId, status).Scan(&output).Error
	if err != nil {
		return []model.GettingTransactionDetailResp{}, err
	}
	return output, nil
}

func (db *DAO)FindAllTransactionOfOwner(ownerId, status int)([]model.GettingTransactionDetailResp, error){
	var output []model.GettingTransactionDetailResp
	sql := `SELECT transactions.id as transactionId, transactions.startTime, transactions.endTime, transactions.licence,
			parkings.address as address, transactions.amount, transactions.status, transactions.created_at,
			transactions.phoneNumber as userPhoneNumber, owners.phoneNumber as hostPhoneNumber, transactions.parkingId
			FROM transactions as tran
			INNER JOIN parkings as p
			ON tran.parkingId = p.id
			INNER JOIN owners as o
			ON p.ownerId = o.credentialId
			WHERE o.credentialId=? AND transactions.status=?`
	err := db.connection.Raw(sql, ownerId, status).Scan(&output).Error
	if err != nil {
		return []model.GettingTransactionDetailResp{}, err
	}
	return output, nil
}