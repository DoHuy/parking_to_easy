package business_logic

import (
	"encoding/json"
	"fmt"
	"github.com/DoHuy/parking_to_easy/firebase"
	"github.com/DoHuy/parking_to_easy/model"
	"github.com/DoHuy/parking_to_easy/mysql"
	"github.com/DoHuy/parking_to_easy/redis"
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
	Dao			*mysql.DAO
	Redis		*redis.Redis
	FireBase	*firebase.FireBaseService
}
func NewService(dao *mysql.DAO, redis *redis.Redis) *TransactionService{
	firebaseService := firebase.NewFireBaseService(redis)
	return &TransactionService{
		Dao: 	dao,
		Redis:	redis,
		FireBase: firebaseService,
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
	// lay lai id cua transaction vua tao
	createdTransaction, err := transactionIface.FindTransactionByCreatedAt(transaction.CreatedAt)
	// Tạo topic để send thông báo trong redis
	// Lấy thông tin chủ bãi đỗ
	var parkingIface mysql.ParkingDAO
	parkingIface = self.Dao
	parking, err := parkingIface.FindParkingByID(fmt.Sprintf("%d", transaction.ParkingId))
	if err != nil {
		return err
	}
	// lay list token cua khach va cua chu bai do
	var deviceIface mysql.DeviceDAO
	deviceIface = self.Dao
	// lay danh sach token cua khach
	userDevices, err := deviceIface.FindTokensOfUser(transaction.CredentialId)
	if err != nil {
		return err
	}
	// lay danh sach token cua chu bai do
	ownerDevices, err := deviceIface.FindTokensOfUser(parking.OwnerId)
	if err != nil {
		return err
	}
	var ownerTokens []string
	var userTokens 	[]string
	for _, ownerDevice := range ownerDevices {
		ownerTokens = append(ownerTokens, ownerDevice.DeviceToken)
	}
	for _, userDevice := range userDevices {
		userTokens = append(userTokens, userDevice.DeviceToken)
	}
	// luu list token cua khach trong redis
	err = self.Redis.SetTokenListTransactionTopic(createdTransaction.ID, createdTransaction.CredentialId, userTokens)
	// luu list token cua chu bai do trong redis
	err = self.Redis.SetTokenListTransactionTopic(createdTransaction.ID, parking.OwnerId, ownerTokens)
	if err != nil {
		return err
	}
	// ban thong bao cho chu bai do la co bai do vua dat
	// khoi tao 1 instance cua firebase service
	err = self.FireBase.SendNotifyToUserOfTransaction(createdTransaction.ID, parking.OwnerId, "THÔNG BÁO", "Bạn vừa nhận được một lượt đăng ký đậu xe")
	if err != nil {
		return err
	}
	//
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
	err := utils.BindRawStructToRespStruct(data, &input)

	if err != nil {
		return err
	}
	var transactionIface mysql.TransactionDAO
	transactionIface = self.Dao
	err = transactionIface.ModifyTransaction(input.TransactionId, input.Status)
	if err != nil {
		return err
	}
	// lấy thông tin của transaction
	transaction, err := transactionIface.FindTransactionById(input.TransactionId)
	if err != nil {
		return err
	}
	// Lấy thông tin của chủ bãi xe
	var parkingIface mysql.ParkingDAO
	parkingIface = self.Dao
	parking, err := parkingIface.FindParkingByID(fmt.Sprintf("%d", transaction.ParkingId))
	if err != nil {
		return err
	}
	// Lấy thông tin user
	var credentialIface mysql.CredentialDAO
	credentialIface = self.Dao
	credential, err := credentialIface.FindCredentialByID(fmt.Sprintf("%d",transaction.CredentialId))
	// bắn thông báo khi thay đổi trạng thái tương ứng
	if input.Status == 4 {
		err = self.FireBase.SendNotifyToUserOfTransaction(transaction.ID, parking.OwnerId, "THÔNG BÁO", fmt.Sprintf("%s đã hủy đặt bãi %s", credential.Username, parking.ParkingName))
		err = self.Redis.DeleteTransactionTopic(transaction.ID)
		if err != nil {
			return err
		}
	} else if input.Status == 2{
		err = self.FireBase.SendNotifyToUserOfTransaction(transaction.ID, transaction.CredentialId, "THÔNG BÁO", fmt.Sprintf("Yêu cầu đặt bãi %s. đã thành công", parking.ParkingName))
		// tru diem user
		customerService := NewCustomerService(self.Dao)
		if err := customerService.SubPoints(transaction.CredentialId, transaction.Amount/1000); err != nil {
			return err
		}
		//

	} else if input.Status == 3{
		err = self.FireBase.SendNotifyToUserOfTransaction(transaction.ID, parking.OwnerId, "THÔNG BÁO", fmt.Sprintf("%s đã vào bãi %s", credential.Username, parking.ParkingName))
		err = self.FireBase.SendNotifyToUserOfTransaction(transaction.ID, transaction.CredentialId,"THÔNG BÁO", fmt.Sprintf("Xe của bạn đã vào bãi %s", parking.ParkingName))
		// Tao 1 goroutine check session

		//
	} else if input.Status == 5{
		err = self.FireBase.SendNotifyToUserOfTransaction(transaction.ID, parking.OwnerId, "THÔNG BÁO", fmt.Sprintf("%s đã chủ động lấy xe tại bãi %s", credential.Username, parking.ParkingName))
		err = self.Redis.DeleteTransactionTopic(transaction.ID)
		if err != nil {
			return err
		}
		// add point to owner
		ownerService := NewOwnerService(self.Dao, self.Redis)
		if err := ownerService.AddPoints(parking.OwnerId, transaction.Amount); err != nil {
			return err
		}
		//
	}
	//
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

func (self *TransactionService)GetParkingIdFromTransaction(transactionId int) (int, error){
	var transactionIface mysql.TransactionDAO
	transactionIface = self.Dao
	transaction, err := transactionIface.FindTransactionById(transactionId)
	if err != nil {
		return 0, err
	}
	return transaction.ParkingId, nil
}

func (self *TransactionService)AnalysisTransaction(data interface{})([]model.AnalysisOutput, error){
	var transactionIface mysql.TransactionDAO
	transactionIface = self.Dao
	var input model.AnalysisInput
	err := utils.BindRawStructToRespStruct(data, &input)
	if err != nil {
		return []model.AnalysisOutput{}, err
	}
	// block a month
	duration, _	:= time.ParseDuration("720h")
	aMonth  	:= duration.Milliseconds()
	allMonth 	:= (input.End - input.Start)/aMonth
	var results []model.AnalysisOutput
	var month int64
	var startMonth int64
	var endMonth   int64
	for month < allMonth {
		startMonth = startMonth + aMonth
		endMonth = startMonth + aMonth
		var tmp model.AnalysisInput
		tmp.Start = startMonth
		tmp.End = endMonth
		output, err := transactionIface.CountFinishedAndCanceledState(tmp)
		if err != nil {
			return []model.AnalysisOutput{}, err
		}
		results = append(results, output)
		month++
	}

	return results, nil
}
// neu trang thai bang 2 thi tru tien cua user
func (self *TransactionService)ExecTransactionBusinesses(status int) error{
	return nil
}