package business_logic

import (
	"encoding/json"
	"github.com/DoHuy/parking_to_easy/model"
	"github.com/DoHuy/parking_to_easy/mysql"
	"time"
)

/*
1 xe chua duoc phe duyet dat cho
2 xe da duoc phe duyet dat cho
3 xe da huy dat cho
4 xe dang duoc gui
5 xe ket thuc gui
*/
const BLOCK_TIME = '5' // block time fix cung
type TransactionService struct {
	Dao		*mysql.DAO
}
func NewService(dao *mysql.DAO) *TransactionService{
	return &TransactionService{
		Dao: dao,
	}
}

func (self *TransactionService)CustomTransaction(payload model.Payload, transaction model.Transaction) (model.Transaction, error){
		transaction.CreatedAt    = time.Now().Format(time.RFC3339)
		transaction.CredentialId = payload.UserId
		transaction.Status		 = 1 // chua duoc duyet
		// cal session and amount
		start, _ := time.Parse(time.RFC3339, transaction.StartTime)
		end, _ 	 := time.Parse(time.RFC3339, transaction.EndTime)
		session  := end.Minute() - start.Minute()
		transaction.Session = session
		// cal amount
		var place model.Parking
		var parkingIface mysql.ParkingDAO
		parkingIface = self.Dao
		place, err := parkingIface.FindParkingByID(string(transaction.ParkingId))
		if err != nil {
			return model.Transaction{}, err
		}
		transaction.Amount = int(session/60)*place.BlockAmount

		return transaction, nil

}

func (self *TransactionService)AddNewTicket(data interface{}) error{
	var transaction model.Transaction
	raw,_ := json.Marshal(data)
	err := json.Unmarshal(raw, &transaction)
	var transactionIface mysql.TransactionDAO
	transactionIface = self.Dao
	err = transactionIface.CreateTransaction(transaction)
	if err != nil {
		return err
	}
	return nil
}