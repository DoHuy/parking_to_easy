package business_logic

import (
	"encoding/json"
	"fmt"
	"github.com/DoHuy/parking_to_easy/model"
	"github.com/DoHuy/parking_to_easy/mysql"
	"github.com/DoHuy/parking_to_easy/utils"
	"time"
)

/*
1 xe chua duoc phe duyet dat cho
2 xe da duoc phe duyet dat cho
3 xe da huy dat cho
4 xe dang duoc gui
5 xe ket thuc gui
*/
// những trạng thái khi book bãi đậu
var STAGE = map[string]int{
	"chưa được duyệt gửi":1,
	"được duyệt gửi":2,
	"đang gửi":3,
	"hủy bỏ gửi":4,
	"kết thúc gửi":5,
}

const BLOCK_TIME = '5' // block time fix cung
type TransactionService struct {
	Dao		*mysql.DAO
}
func NewService(dao *mysql.DAO) *TransactionService{
	return &TransactionService{
		Dao: dao,
	}
}

func (self *TransactionService)CheckSelfBooking(parkingIdOfTran, credentialId int) bool {
	var parking model.Parking
	var parkingIface mysql.ParkingDAO
	parkingIface = self.Dao
	fmt.Println("parkingIdOfTran, credentialId", parkingIdOfTran, credentialId)
	parking,_ = parkingIface.FindParkingByID(fmt.Sprintf("%d", parkingIdOfTran))
	fmt.Println("parkingOwnerId, credentialId", parking.OwnerId, credentialId)
	if parking.OwnerId != credentialId {
		return true
	}
	return false

}
func (self *TransactionService)VerifyBookingStartTime(credentialId int, start string, end string) bool{
	var transactionIface mysql.TransactionDAO
	transactionIface = self.Dao
	transaction, err := transactionIface.FindTheLastTransaction(credentialId)
	if err != nil && err.Error() == "record not found" {
		return true
	}
	currentStartTime, err := time.Parse(time.RFC3339, start)
	//currentEndTime, err := time.Parse(time.RFC3339, end)
	endTime, err := time.Parse(time.RFC3339, transaction.EndTime)
	if currentStartTime.Unix() >= endTime.Unix() {
		return true
	}
	return false

}
func (self *TransactionService)CustomTransaction(payload model.Payload, transaction model.Transaction) (model.Transaction, error){
		transaction.CreatedAt    = time.Now().Format(time.RFC3339)
		transaction.CredentialId = payload.UserId
		transaction.Status		 = 1 // chua duoc duyet
		// cal session and amount
		start, _ := time.Parse(time.RFC3339, transaction.StartTime)
		end, _ 	 := time.Parse(time.RFC3339, transaction.EndTime)
		session  := (end.Unix() - start.Unix())/(60*60)
		transaction.Session = int(session)
		// cal amount
		//fmt.Println("transaction::::: id", transaction.ParkingId)
		var place model.Parking
		var parkingIface mysql.ParkingDAO
		parkingIface = self.Dao

		place, err := parkingIface.FindParkingByID(fmt.Sprintf("%d",transaction.ParkingId))
		if err != nil {
			return model.Transaction{}, err
		}
		transaction.Amount = int(session)*place.BlockAmount
	//	fmt.Println("place:::", transaction.Amount)
	//fmt.Println("place::: transaction::::", transaction.Session)
		return transaction, nil

}

func (self *TransactionService)AddNewTicket(data interface{}) error{
	var transaction model.Transaction
	raw,_ := json.Marshal(data)
	fmt.Println("DATAA::::", data)
	err := json.Unmarshal(raw, &transaction)
	var transactionIface mysql.TransactionDAO
	transactionIface = self.Dao
	err = transactionIface.CreateTransaction(transaction)
	if err != nil {
		return err
	}
	return nil
}

func (self *TransactionService)GetTransactionOfOwnerWithStatus(data interface{}) ([]model.GettingTransactionDetailResp, error){
	var input model.GetTransactionOfOwnerWithStatusInput
	err := utils.BindRawStructToRespStruct(data, &input)
	var transactionIface mysql.TransactionDAO
	transactionIface = self.Dao
	transactions, err := transactionIface.FindAllTransactionOfOwner(input.ParkingId,input.Status)
	if err != nil {
		return []model.GettingTransactionDetailResp{}, err
	}
	return transactions, nil
	// get parking table, get
}
func (self *TransactionService)GetTransactionOfUserWithStatus(data interface{}) ([]model.GettingTransactionDetailResp, error){
	var input model.GetTransactionOfUserWithStatusInput
	fmt.Println("before Data::::", data)
	err := utils.BindRawStructToRespStruct(data, &input)
	fmt.Println("After data::::", input)
	var transactionIface mysql.TransactionDAO
	transactionIface = self.Dao
	transactions, err := transactionIface.FindAllTransactionOfUser(input.UserId, input.Status)
	if err != nil {
		return []model.GettingTransactionDetailResp{}, err
	}
	return transactions, nil
	// get parking table, get


}

func (self *TransactionService)NextStepTransaction(data interface{}) error{
	var input model.ChangingStateTransactionInput
	fmt.Println("before Data::::", data)
	err := utils.BindRawStructToRespStruct(data, &input)
	fmt.Println("After data::::", input)
	if err != nil {
		return err
	}
	var transactionIface mysql.TransactionDAO
	transactionIface = self.Dao
	err = transactionIface.ModifyTransaction(input.TransactionId, input.Status)
	if err != nil {
		return err
	}
	return nil
}

func (self *TransactionService)CheckPermissionForTransaction(transactionId, credentialId int) bool {
	var transaction model.Transaction
	var parking model.Parking
	var transactionIface mysql.TransactionDAO
	var parkingIface mysql.ParkingDAO
	transactionIface = self.Dao
	transaction, _ = transactionIface.FindTransactionById(transactionId)
	if credentialId == transaction.CredentialId {
		return true
	}
	parkingIface = self.Dao
	parking,_ = parkingIface.FindParkingByID(fmt.Sprintf("%d",transaction.ParkingId))
	if parking.OwnerId == credentialId {
		return true
	}
	return false

}

func (self *TransactionService)CheckRuleStateTransaction(transactionId, status int) bool {
	var transaction model.Transaction
	var transactionIface mysql.TransactionDAO
	transactionIface = self.Dao
	transaction,_ = transactionIface.FindTransactionById(transactionId)
	if transaction.Status == 1 && (status == 2 || status == 4) {
		return true
	} else if transaction.Status == 2 && (status == 3 || status == 5) {
		return true
	} else if transaction.Status == 3 && status == 5 {
		return true
	}

	return false
}

func (self *TransactionService)CheckParkingOwnerOfTransaction(ownerId, parkingId int) bool {
	var parking model.Parking
	var parkingIface mysql.ParkingDAO
	parkingIface = self.Dao
	parking, _ = parkingIface.FindParkingByID(fmt.Sprintf("%d", parkingId))
	if parking.OwnerId == ownerId {
		return true
	}

	return false

}

// neu trang thai bang 2 thi tru tien cua user
func (self *TransactionService)ExecTransactionBusinesses(status int) error{
	return nil
}