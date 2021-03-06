package http

import (
	"encoding/json"
	"fmt"
	"github.com/DoHuy/parking_to_easy/business_logic"
	"net/http"
	"strconv"
	"time"

	//"fmt"
	"github.com/DoHuy/parking_to_easy/mysql"
	"regexp"
	"strings"

	//"errors"
	"github.com/DoHuy/parking_to_easy/business_logic/auth"
	"github.com/DoHuy/parking_to_easy/utils"

	//"fmt"
	"github.com/DoHuy/parking_to_easy/config"
	"github.com/DoHuy/parking_to_easy/model"
	"github.com/gin-gonic/gin"
	//"path/filepath"
	//"time"
)

type MiddleWareService struct {
	Factory		*business_logic.ServiceFactory
}
// api Upload
func NewMiddleware(factory	*business_logic.ServiceFactory) *MiddleWareService {

	return &MiddleWareService{
		Factory: factory,
	}
}

func (mid *MiddleWareService)BeforeUpload(c *gin.Context, token string) model.Middleware {
	// check is True token,  expired token
	service := mid.Factory.GetAuthService()
	checked, err := service.CheckTokenIsTrue(token)
	if checked == false && err != nil {
		return model.Middleware{StatusCode: 401, Message: err.Error()}
	}
	form,err := c.MultipartForm()
	files := form.File["upload[]"]
	if form == nil {
		return model.Middleware{StatusCode: 400, Message: "form rỗng"}
	}
	if err != nil {
		fmt.Println("ererererer", err)
		return model.Middleware{StatusCode: 400, Message: err.Error()}
	}
	//var images []string
	for _, file := range files {
		if file.Size >= config.GetMaxUploadedFileSize() {
			return model.Middleware{StatusCode: 400, Message: "File upload không được vượt quá 10 mb"}
		}

	}
	return model.Middleware{}
}

//api Register
func (mid *MiddleWareService)BeforeRegister(c *gin.Context) model.Middleware {
	var credential model.Credential
	cred := utils.GetBodyRequest(c)
	err := json.Unmarshal(cred, &credential)
	if err != nil {
		return model.Middleware{StatusCode: 500, Message: "He thong co su co"}
	}
	//check username
	//var rs model.Credential
	fmt.Println(credential.Username, credential.Email)
	var credIface mysql.CredentialDAO
	credIface = mid.Factory.Dao
	rs, err := credIface.FindCredentialByName(credential.Username)
	fmt.Println("ERRRRR", err)
	//err = utils.BindRawStructToRespStruct(raw, &rs)
	fmt.Println("credential Email", credential.Email)
	fmt.Println("rs Email", rs)
	if rs.Username == credential.Username {
		return model.Middleware{StatusCode: 400, Message: "username đã tồn tại"}
	}
	rs, _ = credIface.FindCredentialByMail(credential.Email)
	fmt.Println("cred: rs", credential, rs)
	if rs.Email == credential.Email {
		return model.Middleware{StatusCode: 400, Message: "email đã tồn tại"}
	}
	// check email
	matched, _ := regexp.Match(`^[a-z][a-z0-9_\.]{5,32}@[a-z0-9]{2,}(\.[a-z0-9]{2,4}){1,2}$`, []byte(credential.Email))
	if matched == false {
		return model.Middleware{StatusCode: 400, Message: "Email không hợp lệ"}
	}
	return model.Middleware{Data: credential}
}
// Before login
func (mid *MiddleWareService)BeforeLogin(c *gin.Context) model.Middleware {
	raw := utils.GetBodyRequest(c)
	var credential model.Credential
	err := json.Unmarshal(raw, &credential)
	if err != nil {
		return model.Middleware{StatusCode: 500, Message: "Hệ thống có sự cố"}
	}
	credential.Username = strings.Trim(credential.Username, " ")
	credential.Password = strings.Trim(credential.Password, " ")
	if len(credential.Username) == 0 || len(credential.Password) == 0 {
		return model.Middleware{StatusCode: 400, Message: "Username và Password không được trống"}
	}
	credential.Password = utils.EncriptPwd(credential.Password)
	return model.Middleware{Data: credential}
}
// api get all users
func (mid *MiddleWareService)BeforeGetAllUsers(c *gin.Context) model.Middleware {
	token, err := utils.GetTokenFromHeader(c)
	fmt.Println("token", token)
	//c.Header("Access-Control-Allow-Origin", "*")
	if err != nil {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check format token
	service := mid.Factory.GetAuthService()
	checked, _ := service.CheckTokenIsTrue(token)
	if checked != true {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check expired token
	checkedExpired, _, _ := service.CheckExpiredToken(token)
	if checkedExpired == true {
		return model.Middleware{StatusCode: 400, Message: "Token hết hạn sử dụng"}
	}
	// check role
	role, err,_ := service.Authorize(token)

	//fmt.Print("role", role, err, rawerr)
	if role != "admin" {
		return model.Middleware{StatusCode: http.StatusServiceUnavailable, Message: "Service không khả dụng"}
	}
	return model.Middleware{Data: map[string]string{"limit": c.Param("limit"), "offset": c.Param("offset")}}

}

//api get detail user
func (mid *MiddleWareService)BeforeGetDetailUser(c *gin.Context) model.Middleware {
	token, err := utils.GetTokenFromHeader(c)
	if err != nil {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	service := mid.Factory.GetAuthService()
	// check format token
	checked, _ := service.CheckTokenIsTrue(token)
	if checked != true {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check expired token
	checkedExpired, _, _ := service.CheckExpiredToken(token)
	if checkedExpired == true {
		return model.Middleware{StatusCode: 400, Message: "Token hết hạn sử dụng"}
	}
	// check accessible
	var payload model.Payload
	secretKey  := string(config.GetSecretKey())
	raw, _ := auth.Decode(token, secretKey)
	err = json.Unmarshal(raw, &payload)
	if err != nil {
		return model.Middleware{StatusCode: 500, Message: "Hệ thống có sự cố"}
	}
	return model.Middleware{Data: payload.UserId}
}
// before create new parking by admin
func (mid *MiddleWareService)BeforeCreateNewParkingByAdmin(c *gin.Context) model.Middleware{
	token, err := utils.GetTokenFromHeader(c)
	if err != nil {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	service := mid.Factory.GetAuthService()
	// check format token
	checked, _ := service.CheckTokenIsTrue(token)
	if checked != true {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check expired token
	checkedExpired, _, _ := service.CheckExpiredToken(token)
	if checkedExpired == true {
		return model.Middleware{StatusCode: 400, Message: "Token hết hạn sử dụng"}
	}
	//
	var payload model.Payload
	secretKey  := string(config.GetSecretKey())
	raw, _ := auth.Decode(token, secretKey)
	err = json.Unmarshal(raw, &payload)
	if err != nil {
		return model.Middleware{StatusCode: 500, Message: "Hệ thống có sự cố"}
	}
	// get body
	var newParking model.NewParkingByAdmin
	body := utils.GetBodyRequest(c)
	fmt.Println(string(body))
	err   = json.Unmarshal(body, &newParking)
	fmt.Println("new PArking :::", newParking)
	if err != nil {
		fmt.Println("ERR:   ", err)
		return model.Middleware{StatusCode: 500, Message: "Hệ thống có sự cố"}
	}
	role, err, _ := service.Authorize(token)
	if err != nil {
		return model.Middleware{StatusCode: 500, Message: "Hệ thống có sự cố"}
	}
	if role != "admin" {
		return model.Middleware{StatusCode: 503, Message: "Dịch vụ không hỗ trợ"}
	}
	// convert data
	newParking.Status = "APPROVED"
	newParking.CreatedAt = time.Now().Format(time.RFC3339)
	newParking.OwnerId = payload.UserId

	return model.Middleware{Data: newParking}

}

func (mid *MiddleWareService)BeforeCreateNewOwner(c *gin.Context) model.Middleware{
	token, err := utils.GetTokenFromHeader(c)
	if err != nil {
		fmt.Println("ERR:   ", err)
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	service := mid.Factory.GetAuthService()
	// check format token
	checked, _ := service.CheckTokenIsTrue(token)
	if checked != true {
		fmt.Println("ERR:   ", err)
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check expired token
	checkedExpired, _, _ := service.CheckExpiredToken(token)
	if checkedExpired == true {
		return model.Middleware{StatusCode: 400, Message: "Token hết hạn sử dụng"}
	}
	//
	var payload model.Payload
	secretKey  := string(config.GetSecretKey())
	raw, _ := auth.Decode(token, secretKey)
	err = json.Unmarshal(raw, &payload)
	if err != nil {
		return model.Middleware{StatusCode: 500, Message: "Hệ thống có sự cố"}
	}
	// get body
	var newOwner model.Owner
	body := utils.GetBodyRequest(c)
	err   = json.Unmarshal(body, &newOwner)
	fmt.Println("new Owner :::", newOwner)
	if err != nil {
		fmt.Println("ERR:   ", err)
		return model.Middleware{StatusCode: 500, Message: "Hệ thống có sự cố"}
	}
	//convert data
	fmt.Println("PAyload::::", payload)
	newOwner.CredentialId = payload.UserId
	newOwner.Status		  = "ENABLE"
	newOwner.CreatedAt 	  = time.Now().Format(time.RFC3339)
	return model.Middleware{Data: newOwner}
}

func (mid *MiddleWareService)BeforeCreateNewParkingByOwner(c *gin.Context) model.Middleware{
	token, err := utils.GetTokenFromHeader(c)
	if err != nil {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	service := mid.Factory.GetAuthService()
	// check format token
	checked, _ := service.CheckTokenIsTrue(token)
	if checked != true {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check expired token
	checkedExpired, _, _ := service.CheckExpiredToken(token)
	if checkedExpired == true {
		return model.Middleware{StatusCode: 400, Message: "Token hết hạn sử dụng"}
	}
	//
	var payload model.Payload
	secretKey  := string(config.GetSecretKey())
	raw, _ := auth.Decode(token, secretKey)
	err = json.Unmarshal(raw, &payload)
	if err != nil {
		return model.Middleware{StatusCode: 500, Message: "Hệ thống có sự cố"}
	}
	// get body
	var newParking model.Parking
	body := utils.GetBodyRequest(c)
	err   = json.Unmarshal(body, &newParking)
	if err != nil {
		return model.Middleware{StatusCode: 500, Message: "Hệ thống có sự cố"}
	}
	// convert data
	newParking.OwnerId = payload.UserId
	newParking.CreatedAt = time.Now().Format(time.RFC3339)
	newParking.Status	 = "PENDING"
	parkingService := mid.Factory.GetParkingService()
	checkedLocation := parkingService.CheckExistedLocation(newParking.Longitude, newParking.Latitude)
	if checkedLocation == false {
		return model.Middleware{StatusCode:400, Message: "Bạn đã chia sẻ vị trí này rồi"}
	}
	return model.Middleware{Data: newParking}
}

func (mid *MiddleWareService)BeforeGetOwnerById(c *gin.Context) model.Middleware{
	token, err := utils.GetTokenFromHeader(c)
	if err != nil {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check format token
	service := mid.Factory.GetAuthService()
	checked, _ := service.CheckTokenIsTrue(token)
	if checked != true {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check expired token
	checkedExpired, _, _ := service.CheckExpiredToken(token)
	if checkedExpired == true {
		return model.Middleware{StatusCode: 400, Message: "Token hết hạn sử dụng"}
	}
	if err != nil {
		return model.Middleware{StatusCode: 500, Message: "Hệ thống có sự cố"}
	}
	// check role
	role, err,_ := service.Authorize(token)

	//fmt.Print("role", role, err, rawerr)
	if role != "admin" {
		return model.Middleware{StatusCode: http.StatusServiceUnavailable, Message: "Service không khả dụng"}
	}
	// get ownerId
	return model.Middleware{Data: c.Param("id")}
}

func (mid *MiddleWareService)BeforeVerifyParking(c *gin.Context) model.Middleware {
	token, err := utils.GetTokenFromHeader(c)
	if err != nil {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check format token
	service := mid.Factory.GetAuthService()
	checked, _ := service.CheckTokenIsTrue(token)
	if checked != true {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check expired token
	checkedExpired, _, _ := service.CheckExpiredToken(token)
	if checkedExpired == true {
		return model.Middleware{StatusCode: 400, Message: "Token hết hạn sử dụng"}
	}
	// check role
	role, err,_ := service.Authorize(token)
	if role != "admin" {
		return model.Middleware{StatusCode: 503, Message: "Hệ thống không hỗ trợ dịch vụ này"}
	}
	// convert data
	var input model.VerifyingParkingInput
	rawBody := utils.GetBodyRequest(c)
	err = json.Unmarshal(rawBody, &input)
	id := c.Param("id")
	input.ID = id
	input.ModifiedAt = time.Now().Format(time.RFC3339)
	// kiem tra su ton tai cua parking
	parkingService := mid.Factory.GetParkingService()
	if checked := parkingService.CheckExistedParking(id); checked != true{
		return model.Middleware{StatusCode: 404, Message: "Bãi đỗ không tồn tại"}
	}
	if input.Status != "REJECTED" && input.Status != "APPROVED" {
		return model.Middleware{StatusCode: 400, Message: "Trạng thái cập nhật không đúng"}
	}

	return model.Middleware{Data: input}
}

func (mid *MiddleWareService)AfterGetAllParkings(parkings []model.Parking) model.Middleware {
	var rs []model.Parking
	for i := 0 ; i < len(parkings) ; i++ {
		if parkings[i].Status == "APPROVED" {
			//fmt.Println()
			rs = append(rs, parkings[i])
			//fmt.Println("rsiii", rs[i])
		}
	}
	return model.Middleware{Data: rs}
}

func (mid *MiddleWareService)BeforeCalculateAmountParking(c *gin.Context) model.Middleware{
	token, err := utils.GetTokenFromHeader(c)
	if err != nil {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check format token
	authService := mid.Factory.GetAuthService()
	checked, _ := authService.CheckTokenIsTrue(token)
	if checked != true {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check expired token
	checkedExpired, _, _ := authService.CheckExpiredToken(token)
	if checkedExpired == true {
		return model.Middleware{StatusCode: 400, Message: "Token hết hạn sử dụng"}
	}
	// check chu so huu bai do
	idParking := c.Param("id")
	var parking model.Parking
	var parkingIface mysql.ParkingDAO
	parkingIface = mid.Factory.Dao
	parking, err = parkingIface.FindParkingByID(idParking)
	if err != nil {
		if err.Error() == "record not found"{
			return model.Middleware{StatusCode: 404, Message: "Bãi đỗ này không tồn tại"}

		}
		return model.Middleware{StatusCode: 500, Message: "Hệ thống này có sự cố"}
	}
	//
	var payload model.Payload
	secretKey  := string(config.GetSecretKey())
	raw, _ := auth.Decode(token, secretKey)
	err = json.Unmarshal(raw, &payload)
	if err != nil {
		return model.Middleware{StatusCode: 500, Message: "Hệ thống có sự cố"}
	}
	//
	if payload.UserId != parking.OwnerId {
		return model.Middleware{StatusCode: 403, Message: "Bạn không có quyền truy cập"}
	}
	return model.Middleware{Data: idParking}
}

func (mid *MiddleWareService)BeforeGetAllOwners(c *gin.Context) model.Middleware{
	token, err := utils.GetTokenFromHeader(c)
	if err != nil {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check format token
	authService := mid.Factory.GetAuthService()
	checked, _ := authService.CheckTokenIsTrue(token)
	if checked != true {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check expired token
	checkedExpired, _, _ := authService.CheckExpiredToken(token)
	if checkedExpired == true {
		return model.Middleware{StatusCode: 400, Message: "Token hết hạn sử dụng"}
	}
	role, err,_ := authService.Authorize(token)
	if role != "admin" {
		return model.Middleware{StatusCode: 503, Message: "Hệ thống không hỗ trợ dịch vụ này"}
	}
	return model.Middleware{}
}

func (mid *MiddleWareService)BeforeDisableOwner(c *gin.Context) model.Middleware{
	token, err := utils.GetTokenFromHeader(c)
	if err != nil {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check format token
	authService := mid.Factory.GetAuthService()
	checked, _ := authService.CheckTokenIsTrue(token)
	if checked != true {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check expired token
	checkedExpired, _, _ := authService.CheckExpiredToken(token)
	if checkedExpired == true {
		return model.Middleware{StatusCode: 400, Message: "Token hết hạn sử dụng"}
	}
	role, err,_ := authService.Authorize(token)
	if role != "admin" {
		return model.Middleware{StatusCode: 503, Message: "Hệ thống không hỗ trợ dịch vụ này"}
	}
	fmt.Println("DATAAAAA::::::", c.Param("id"))
	return model.Middleware{Data:model.DataStruct{CredentialId: c.Param("id"), Status: "DISABLED", ModifiedAt:time.Now().Format(time.RFC3339)}}
}

func (mid *MiddleWareService)BeforeModifyParkingByOwner(c *gin.Context) model.Middleware{
	token, err := utils.GetTokenFromHeader(c)
	if err != nil {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check format token
	authService := mid.Factory.GetAuthService()
	checked, _ := authService.CheckTokenIsTrue(token)
	if checked != true {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check expired token
	checkedExpired, _, _ := authService.CheckExpiredToken(token)
	if checkedExpired == true {
		return model.Middleware{StatusCode: 400, Message: "Token hết hạn sử dụng"}
	}
	var payload model.Payload
	secretKey  := string(config.GetSecretKey())
	raw, _ := auth.Decode(token, secretKey)
	err = json.Unmarshal(raw, &payload)
	if err != nil {
		return model.Middleware{StatusCode: 500, Message: "Hệ thống có sự cố"}
	}
	// check xem owner co so huu bai nay hay khong
	var parking model.Parking
	var parkingIface mysql.ParkingDAO
	parkingIface = mid.Factory.Dao
	parking, err = parkingIface.FindParkingByID(c.Param("id"))
	if err != nil {
		if err.Error() == "record not found" {
			return model.Middleware{StatusCode: 404, Message: "Bãi đỗ này không tồn tại"}
		}
		return model.Middleware{StatusCode: 500, Message: "Hệ thống có sự cố"}
	}
	if parking.OwnerId != payload.UserId {
		return model.Middleware{StatusCode:403, Message: "Bạn không sở hữu bãi đỗ này"}
	}

	// get body
	var updatedParking model.Parking
	body := utils.GetBodyRequest(c)
	err   = json.Unmarshal(body, &updatedParking)
	fmt.Println("updated parking :::", updatedParking)
	// convert data
	type UpdatedData struct {
		ID			string		`json:"id"`
		Capacity	string		`json:"capacity"`
		ModifiedAt	string		`json:"modified_at"`
	}
	return model.Middleware{Data:UpdatedData{ID: c.Param("id"), Capacity: updatedParking.Capacity, ModifiedAt: time.Now().Format(time.RFC3339)}}
}

func (mid *MiddleWareService)BeforeDeleteParkingByOwner(c *gin.Context) model.Middleware{
	token, err := utils.GetTokenFromHeader(c)
	if err != nil {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check format token
	authService := mid.Factory.GetAuthService()
	checked, _  := authService.CheckTokenIsTrue(token)
	if checked != true {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check expired token
	checkedExpired, _, _ := authService.CheckExpiredToken(token)
	if checkedExpired == true {
		return model.Middleware{StatusCode: 400, Message: "Token hết hạn sử dụng"}
	}
	var payload model.Payload
	secretKey  := string(config.GetSecretKey())
	raw, _ := auth.Decode(token, secretKey)
	err = json.Unmarshal(raw, &payload)
	if err != nil {
		return model.Middleware{StatusCode: 500, Message: "Hệ thống có sự cố"}
	}
	// check xem owner co so huu bai nay hay khong
	var parking model.Parking
	var parkingIface mysql.ParkingDAO
	parkingIface = mid.Factory.Dao
	parking, err = parkingIface.FindParkingByID(c.Param("id"))
	if err != nil {
		if err.Error() == "record not found" {
			return model.Middleware{StatusCode: 404, Message: "Bãi đỗ này không tồn tại"}
		}
		return model.Middleware{StatusCode: 500, Message: "Hệ thống có sự cố"}
	}
	if parking.OwnerId != payload.UserId {
		return model.Middleware{StatusCode:403, Message: "Bạn không sở hữu bãi đỗ này"}
	}

	// convert data
	type DeletedData struct {
		ID			string		`json:"id"`
		DeletedAt	string		`json:"deleted_at"`
	}
	return model.Middleware{Data:DeletedData{ID: c.Param("id"), DeletedAt: time.Now().Format(time.RFC3339)}}
}

func (mid *MiddleWareService)BeforeCalculateAmountAndVote(c *gin.Context) model.Middleware{
	token, err := utils.GetTokenFromHeader(c)
	if err != nil {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check format token
	authService := mid.Factory.GetAuthService()
	checked, _ := authService.CheckTokenIsTrue(token)
	if checked != true {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check expired token
	checkedExpired, _, _ := authService.CheckExpiredToken(token)
	if checkedExpired == true {
		return model.Middleware{StatusCode: 400, Message: "Token hết hạn sử dụng"}
	}
	var payload model.Payload
	secretKey  := string(config.GetSecretKey())
	raw, _ := auth.Decode(token, secretKey)
	err = json.Unmarshal(raw, &payload)
	if err != nil {
		return model.Middleware{StatusCode: 500, Message: "Hệ thống có sự cố"}
	}
	// check xem owner co so huu bai nay hay khong
	var parking model.Parking
	var parkingIface mysql.ParkingDAO
	parkingIface = mid.Factory.Dao
	parking, err = parkingIface.FindParkingByID(c.Param("id"))
	if err != nil {
		if err.Error() == "record not found" {
			return model.Middleware{StatusCode: 404, Message: "Bãi đỗ này không tồn tại"}
		}
		return model.Middleware{StatusCode: 500, Message: "Hệ thống có sự cố"}
	}
	if parking.OwnerId != payload.UserId {
		return model.Middleware{StatusCode:403, Message: "Bạn không sở hữu bãi đỗ này"}
	}
	return model.Middleware{Data: c.Param("id")}
}

func (mid *MiddleWareService)BeforeGetTransactionOfUser(c *gin.Context) model.Middleware{
	token, err := utils.GetTokenFromHeader(c)
	if err != nil {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check format token
	authService := mid.Factory.GetAuthService()
	checked, _ := authService.CheckTokenIsTrue(token)
	if checked != true {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check expired token
	checkedExpired, _, _ := authService.CheckExpiredToken(token)
	if checkedExpired == true {
		return model.Middleware{StatusCode: 400, Message: "Token hết hạn sử dụng"}
	}
	var payload model.Payload
	secretKey  := string(config.GetSecretKey())
	raw, _ := auth.Decode(token, secretKey)
	err = json.Unmarshal(raw, &payload)
	if err != nil {
		return model.Middleware{StatusCode: 500, Message: "Hệ thống có sự cố"}
	}
	// convert data
	status,_ := strconv.Atoi(c.Param("status"))

	return model.Middleware{Data: model.GetTransactionOfUserWithStatusInput{Status: status, UserId: payload.UserId}}
}

func (mid *MiddleWareService)BeforeGetAllTransaction(c *gin.Context)model.Middleware{
	token, err := utils.GetTokenFromHeader(c)
	if err != nil {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check format token
	authService := mid.Factory.GetAuthService()
	checked, _ := authService.CheckTokenIsTrue(token)
	if checked != true {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check expired token
	checkedExpired, _, _ := authService.CheckExpiredToken(token)
	if checkedExpired == true {
		return model.Middleware{StatusCode: 400, Message: "Token hết hạn sử dụng"}
	}

	var payload model.Payload
	secretKey  := string(config.GetSecretKey())
	raw, _ := auth.Decode(token, secretKey)
	err = json.Unmarshal(raw, &payload)
	if err != nil {
		return model.Middleware{StatusCode: 500, Message: "Hệ thống có sự cố"}
	}
	if payload.Role != "admin" {
		return model.Middleware{StatusCode: 503, Message: "Dịch vụ không sẵn có"}
	}
	return model.Middleware{}
}

func (mid *MiddleWareService)BeforeRating(c *gin.Context) model.Middleware{
	token, err := utils.GetTokenFromHeader(c)
	if err != nil {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check format token
	authService := mid.Factory.GetAuthService()
	checked, _ := authService.CheckTokenIsTrue(token)
	if checked != true {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check expired token
	checkedExpired, _, _ := authService.CheckExpiredToken(token)
	if checkedExpired == true {
		return model.Middleware{StatusCode: 400, Message: "Token hết hạn sử dụng"}
	}
	//convert data
	var payload model.Payload
	secretKey  := string(config.GetSecretKey())
	raw, _ := auth.Decode(token, secretKey)
	err = json.Unmarshal(raw, &payload)
	if err != nil {
		return model.Middleware{StatusCode: 500, Message: "Hệ thống có sự cố"}
	}
	rawBody := utils.GetBodyRequest(c)
	var rating model.Rating
	var votingInput model.VotingInput
	err = json.Unmarshal(rawBody, &votingInput)
	tranService := mid.Factory.GetTransactionService()
	parkingId, err := tranService.GetParkingIdFromTransaction(votingInput.TransactionId)
	rating.CredentialId = payload.UserId
	rating.Stars = votingInput.Stars
	rating.ParkingId = parkingId
	if err != nil {
		return model.Middleware{StatusCode: 400, Message: "Hệ thống có sự cố"}
	}

	return model.Middleware{Data: rating}

}

func (mid *MiddleWareService)BeforeRecommendParking(c *gin.Context) model.Middleware{
	token, err := utils.GetTokenFromHeader(c)
	if err != nil {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check format token
	authService := mid.Factory.GetAuthService()
	checked, _ := authService.CheckTokenIsTrue(token)
	if checked != true {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check expired token
	checkedExpired, _, _ := authService.CheckExpiredToken(token)
	if checkedExpired == true {
		return model.Middleware{StatusCode: 400, Message: "Token hết hạn sử dụng"}
	}
	return model.Middleware{}
}

func (mid *MiddleWareService)BeforeCreateNewTransaction(c *gin.Context) model.Middleware{
	token, err := utils.GetTokenFromHeader(c)
	if err != nil {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check format token
	authService := mid.Factory.GetAuthService()
	checked, _ := authService.CheckTokenIsTrue(token)
	if checked != true {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check expired token
	checkedExpired, _, _ := authService.CheckExpiredToken(token)
	if checkedExpired == true {
		return model.Middleware{StatusCode: 400, Message: "Token hết hạn sử dụng"}
	}
	// convert data
	var transaction	model.Transaction
	var payload model.Payload
	secretKey  := string(config.GetSecretKey())
	raw, _ := auth.Decode(token, secretKey)
	err = json.Unmarshal(raw, &payload)
	rawBody := utils.GetBodyRequest(c)
	err = json.Unmarshal(rawBody, &transaction)
	tranService := mid.Factory.GetTransactionService()
	//
	flag := tranService.CheckSelfBooking(transaction.ParkingId, payload.UserId)
	if flag != true {
		return model.Middleware{StatusCode: 403, Message: "Bạn không được tự đặt chỗ cho bãi của chính mình"}
	}
	converted, err := tranService.CustomTransaction(payload, transaction)
	// check wallet
	customerService := mid.Factory.GetCustomerService()
	checkedWallet := customerService.CheckWallet(converted.Amount, converted.CredentialId)
	if checkedWallet == false {
		return model.Middleware{StatusCode: 400, Message: "Bạn không còn đủ điểm để thực hiện giao dịch"}
	}
	flagCheckTime := tranService.VerifyBookingStartTime(converted.CredentialId, converted.StartTime, converted.EndTime)
	if flagCheckTime != true {
		return model.Middleware{StatusCode: 400, Message: "Ngày bắt đầu của session mới không được trước ngày kết thúc của session trc đó"}
	}
	return model.Middleware{Data: converted}
}

func (mid *MiddleWareService)BeforeGetAllTransactionOfOwner(c *gin.Context) model.Middleware{
	token, err := utils.GetTokenFromHeader(c)
	if err != nil {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check format token
	authService := mid.Factory.GetAuthService()
	checked, _ := authService.CheckTokenIsTrue(token)
	if checked != true {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check expired token
	checkedExpired, _, _ := authService.CheckExpiredToken(token)
	if checkedExpired == true {
		return model.Middleware{StatusCode: 400, Message: "Token hết hạn sử dụng"}
	}

	var payload model.Payload
	secretKey  := string(config.GetSecretKey())
	raw, _ := auth.Decode(token, secretKey)
	err = json.Unmarshal(raw, &payload)
	if err != nil {
		return model.Middleware{StatusCode: 500, Message: "Hệ thống có sự cố"}
	}
	if payload.Role == "admin" {
		return model.Middleware{StatusCode: 503, Message: "Dịch vụ không sẵn có"}
	}
	// check bai do thuoc chu so huu
	status,_ := strconv.Atoi(c.Param("status"))
	parkingId,_ := strconv.Atoi(c.Param("parkingId"))
	tranService := mid.Factory.GetTransactionService()
	if flag := tranService.CheckParkingOwnerOfTransaction(payload.UserId, parkingId); flag != true {
		return model.Middleware{StatusCode: 403, Message: "Bạn không có quyền truy cập tới bãi đỗ này"}
	}
	// convert data
	return model.Middleware{Data: model.GetTransactionOfOwnerWithStatusInput{ParkingId: parkingId, Status: status}}
}

func (mid *MiddleWareService)BeforeDeclineTransaction(c *gin.Context) model.Middleware{
	token, err := utils.GetTokenFromHeader(c)
	if err != nil {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check format token
	authService := mid.Factory.GetAuthService()
	checked, _ := authService.CheckTokenIsTrue(token)
	if checked != true {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check expired token
	checkedExpired, _, _ := authService.CheckExpiredToken(token)
	if checkedExpired == true {
		return model.Middleware{StatusCode: 400, Message: "Token hết hạn sử dụng"}
	}
	var payload model.Payload
	secretKey  := string(config.GetSecretKey())
	raw, _ := auth.Decode(token, secretKey)
	err = json.Unmarshal(raw, &payload)
	if err != nil {
		return model.Middleware{StatusCode: 500, Message: "Hệ thống có sự cố"}
	}
	return model.Middleware{Data: payload.UserId}
}

func (mid *MiddleWareService)BeforeChangeStateTransaction(c *gin.Context) model.Middleware{
	token, err := utils.GetTokenFromHeader(c)
	if err != nil {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check format token
	authService := mid.Factory.GetAuthService()
	checked, _ := authService.CheckTokenIsTrue(token)
	if checked != true {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check expired token
	checkedExpired, _, _ := authService.CheckExpiredToken(token)
	if checkedExpired == true {
		return model.Middleware{StatusCode: 400, Message: "Token hết hạn sử dụng"}
	}
	var payload model.Payload
	secretKey  := string(config.GetSecretKey())
	raw, _ := auth.Decode(token, secretKey)
	err = json.Unmarshal(raw, &payload)
	if err != nil {
		return model.Middleware{StatusCode: 500, Message: "Hệ thống có sự cố"}
	}
	// convert data
	var input model.ChangingStateTransactionInput
	rawBody  := utils.GetBodyRequest(c)
	err = json.Unmarshal(rawBody, &input)
	input.CredentialId = payload.UserId
	// check xem co phai chu cua bai, hay la khach book
	tranService := mid.Factory.GetTransactionService()
	flag := tranService.CheckPermissionForTransaction(input.TransactionId, input.CredentialId)
	if flag != true {
		return model.Middleware{StatusCode: 403, Message:"Không có quyền thao tác với giao dịch này"}
	}
	// check changing status
	flag = tranService.CheckRuleStateTransaction(input.TransactionId, input.Status)
	if flag == false {
		return model.Middleware{StatusCode:400, Message: "Cập nhật trạng thái khống đúng luật"}
	}
	return model.Middleware{Data: input}
}

func (mid *MiddleWareService)BeforeCreateAndRemoveDeviceToken(c *gin.Context) model.Middleware{
	token, err := utils.GetTokenFromHeader(c)
	if err != nil {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check format token
	authService := mid.Factory.GetAuthService()
	checked, _ := authService.CheckTokenIsTrue(token)
	if checked != true {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check expired token
	checkedExpired, _, _ := authService.CheckExpiredToken(token)
	if checkedExpired == true {
		return model.Middleware{StatusCode: 400, Message: "Token hết hạn sử dụng"}
	}
	var input model.UserDevice
	rawBody  := utils.GetBodyRequest(c)
	err = json.Unmarshal(rawBody, &input)
	var payload model.Payload
	secretKey  := string(config.GetSecretKey())
	raw, _ := auth.Decode(token, secretKey)
	err = json.Unmarshal(raw, &payload)
	input.CredentialId = payload.UserId
	return model.Middleware{Data: input}
}

func (mid *MiddleWareService)BeforeAnalysisTransaction(c *gin.Context) model.Middleware{
	token, err := utils.GetTokenFromHeader(c)
	if err != nil {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check format token
	authService := mid.Factory.GetAuthService()
	checked, _ := authService.CheckTokenIsTrue(token)
	if checked != true {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check expired token
	checkedExpired, _, _ := authService.CheckExpiredToken(token)
	if checkedExpired == true {
		return model.Middleware{StatusCode: 400, Message: "Token hết hạn sử dụng"}
	}
	var payload model.Payload
	secretKey  := string(config.GetSecretKey())
	raw, _ := auth.Decode(token, secretKey)
	err = json.Unmarshal(raw, &payload)
	if payload.Role != "admin" {
		return model.Middleware{StatusCode: 503, Message: "Dịch vụ không được hỗ trợ"}
	}
	start, err := strconv.Atoi(c.Param("start"))
	end, err := strconv.Atoi(c.Param("end"))
	return model.Middleware{Data: model.AnalysisInput{Start: int64(start), End: int64(end)}}
}
