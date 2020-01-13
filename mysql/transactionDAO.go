package mysql

import (
	"fmt"
	"github.com/DoHuy/parking_to_easy/model"
	"time"
)

type TransactionDAO interface {
	CalTotalAmountOfParking(id string) (int, error)
	FindTransactionOfUser(id int) ([]model.Transaction, error)
	FindAllTransaction() ([]model.Transaction, error)
	CreateTransaction(tran model.Transaction) error
	FindAllTransactionOfUser(userId, status int) ([]model.GettingTransactionDetailResp, error)
	FindAllTransactionOfOwner(parkingId, status int)([]model.GettingTransactionDetailResp, error)
	ModifyTransaction(transactionId, status int) error
	FindTransactionById(id int) (model.Transaction, error)
	FindTheLastTransaction(credentialId int) (model.Transaction, error)
	CountFinishedAndCanceledState(input model.AnalysisInput) (model.AnalysisOutput, error)
}

func (db *DAO)CalTotalAmountOfParking(id string) (int, error) {
	const FINISHED = 5
	type Data struct {
		Total	int		`json:"total"`
	}
	var data Data
	err := db.connection.Raw("SELECT SUM(amount) AS total FROM transactions WHERE parkingId=? AND `status`=?", id, FINISHED).Scan(&data).Error
	if err != nil {
		return 0, nil
	}
	fmt.Println("DATA:::",data)
	return data.Total/1000, nil
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

	sql := `SELECT id, startTime, endTime, licence, address, amount, status, created_at, userPhoneNumber, hostPhoneNumber  
			FROM (SELECT tran.id, tran.startTime startTime, tran.endTime endTime, tran.licence, p.address as address, 
					tran.amount, tran.status status, tran.created_at created_at, tran.phoneNumber userPhoneNumber, 
					o.phoneNumber hostPhoneNumber 
					FROM transactions as tran 
					INNER JOIN parkings as p ON tran.parkingId = p.id 
					INNER JOIN owners as o ON p.ownerId = o.credentialId 
					WHERE tran.credentialId=? AND tran.status=?) AS tb`
	err := db.connection.Raw(sql, userId, status).Scan(&output).Error
	if err != nil {
		return []model.GettingTransactionDetailResp{}, err
	}
	return output, nil
}

func (db *DAO)FindAllTransactionOfOwner(parkingId, status int)([]model.GettingTransactionDetailResp, error){
	var output []model.GettingTransactionDetailResp
	sql := `SELECT id, startTime, endTime, licence, address, amount, status, created_at, userPhoneNumber, hostPhoneNumber  
			FROM (SELECT tran.id, tran.startTime startTime, tran.endTime endTime, tran.licence, p.address as address, 
					tran.amount, tran.status status, tran.created_at created_at, tran.phoneNumber userPhoneNumber, 
					o.phoneNumber hostPhoneNumber 
					FROM transactions as tran 
					INNER JOIN parkings as p ON tran.parkingId = p.id 
					INNER JOIN owners as o ON p.ownerId = o.credentialId 
					WHERE tran.parkingId=? AND tran.status=?) AS tb`
	err := db.connection.Raw(sql,parkingId, status).Scan(&output).Error
	if err != nil {
		return []model.GettingTransactionDetailResp{}, err
	}
	fmt.Println("OUTPUSADADSAD:::", output)
	return output, nil
}

func (db *DAO)ModifyTransaction(transactionId, status int) error{
	err := db.connection.Exec("UPDATE transactions SET `status`=? WHERE id=?", status, transactionId).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *DAO)FindTransactionById(id int) (model.Transaction, error) {
	var transaction model.Transaction
	err := db.connection.Raw("SELECT* FROM transactions WHERE id=?", id).Scan(&transaction).Error
	if err != nil {
		return model.Transaction{}, err
	}

	return transaction, nil
}

func (db *DAO)FindTheLastTransaction(credentialId int) (model.Transaction, error) {
	var transaction model.Transaction
	err := db.connection.Raw("SELECT* FROM transactions WHERE credentialId=? AND status BETWEEN 1 AND 3  ORDER BY id desc", credentialId).Scan(&transaction).Error
	if err != nil {
		return model.Transaction{}, err
	}
	return transaction, nil
}

func (db *DAO)CountFinishedAndCanceledState(input model.AnalysisInput) (model.AnalysisOutput, error) {
		var finishedTransactions []model.Transaction
		var canceledTransactions []model.Transaction
		var output model.AnalysisOutput

		err := db.connection.Raw(`
			SELECT * FROM transactions WHERE status=4`).Scan(&canceledTransactions).Error
		err = db.connection.Raw(`
			SELECT * FROM transactions WHERE status=5`).Scan(&finishedTransactions).Error
		if err != nil {
			return model.AnalysisOutput{}, err
		}
		//fmt.Println("finishedTransactions", start.Unix())
		//fmt.Println("canceledTransactions", end.Unix())
		// count cancel
		for i:=0; i<len(canceledTransactions); i++ {
			created, _ := time.Parse(time.RFC3339, canceledTransactions[i].CreatedAt)
			fmt.Println("created::::::::", created.Unix())
			if input.Start >= created.Unix() && created.Unix()<=input.End{
				output.Canceled++
			}
		}
		// count finished
	for i:=0; i<len(finishedTransactions); i++ {
		created, _ := time.Parse(time.RFC3339, finishedTransactions[i].CreatedAt)
		if input.Start>=created.Unix() && created.Unix()<=input.End{
			output.Finished++
		}
	}
	return output, nil
}