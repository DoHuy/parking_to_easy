package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	//"fmt"
	"github.com/DoHuy/parking_to_easy/mysql"
	"regexp"
	"strings"

	//"errors"
	"github.com/DoHuy/parking_to_easy/auth"
	"github.com/DoHuy/parking_to_easy/utils"

	//"fmt"
	"github.com/DoHuy/parking_to_easy/config"
	"github.com/DoHuy/parking_to_easy/model"
	"github.com/gin-gonic/gin"
	//"path/filepath"
	//"time"
)

type Middleware struct {
	DAO		*mysql.DAO
	Auth	*auth.Auth
}
// api Upload
func NewMiddleware(dao *mysql.DAO, auth *auth.Auth) *Middleware {
	return &Middleware{
		DAO: dao,
		Auth: auth,
	}
}

func (mid *Middleware)BeforeUpload(c *gin.Context, token string) model.Middleware {
	// check is True token,  expired token
	checked, err := mid.Auth.CheckTokenIsTrue(token)
	if checked == false && err != nil {
		return model.Middleware{StatusCode: 401, Message: err.Error()}
	}
	form,err := c.MultipartForm()
	files := form.File["upload[]"]
	if err != nil {
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
func (mid *Middleware)BeforeRegister(c *gin.Context) model.Middleware {
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
	credIface = mid.DAO
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
func (mid *Middleware)BeforeLogin(c *gin.Context) model.Middleware {
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
func (mid *Middleware)BeforeGetAllUsers(c *gin.Context) model.Middleware {
	token, err := utils.GetTokenFromHeader(c)
	fmt.Println("token", token)
	c.Header("Access-Control-Allow-Origin", "*")
	if err != nil {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check format token
	checked, _ := mid.Auth.CheckTokenIsTrue(token)
	if checked != true {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check expired token
	checkedExpired, _, _ := mid.Auth.CheckExpiredToken(token)
	if checkedExpired == true {
		return model.Middleware{StatusCode: 400, Message: "Token hết hạn sử dụng"}
	}
	// check role
	role, err,_ := mid.Auth.Authorize(token)

	//fmt.Print("role", role, err, rawerr)
	if role != "admin" {
		return model.Middleware{StatusCode: http.StatusServiceUnavailable, Message: "Service không khả dụng"}
	}
	return model.Middleware{Data: map[string]string{"limit": c.Param("limit"), "offset": c.Param("offset")}}

}

//api get detail user
func (mid *Middleware)BeforeGetDetailUser(c *gin.Context) model.Middleware {
	token, err := utils.GetTokenFromHeader(c)
	if err != nil {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check format token
	checked, _ := mid.Auth.CheckTokenIsTrue(token)
	if checked != true {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check expired token
	checkedExpired, _, _ := mid.Auth.CheckExpiredToken(token)
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
func (mid *Middleware)BeforeCreateNewParkingByAdmin(c *gin.Context) model.Middleware{
	token, err := utils.GetTokenFromHeader(c)
	if err != nil {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check format token
	checked, _ := mid.Auth.CheckTokenIsTrue(token)
	if checked != true {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check expired token
	checkedExpired, _, _ := mid.Auth.CheckExpiredToken(token)
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
	err   = json.Unmarshal(body, &newParking)
	fmt.Println("new PArking :::", newParking)
	if err != nil {
		fmt.Println("ERR:   ", err)
		return model.Middleware{StatusCode: 500, Message: "Hệ thống có sự cố"}
	}
	role, err, _ := mid.Auth.Authorize(token)
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

func (mid *Middleware)BeforeCreateNewOwner(c *gin.Context) model.Middleware{
	token, err := utils.GetTokenFromHeader(c)
	if err != nil {
		fmt.Println("ERR:   ", err)
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check format token
	checked, _ := mid.Auth.CheckTokenIsTrue(token)
	if checked != true {
		fmt.Println("ERR:   ", err)
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check expired token
	checkedExpired, _, _ := mid.Auth.CheckExpiredToken(token)
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

func (mid *Middleware)BeforeCreateNewParkingByOwner(c *gin.Context) model.Middleware{
	token, err := utils.GetTokenFromHeader(c)
	if err != nil {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check format token
	checked, _ := mid.Auth.CheckTokenIsTrue(token)
	if checked != true {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check expired token
	checkedExpired, _, _ := mid.Auth.CheckExpiredToken(token)
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
	fmt.Println("new parking :::", newParking)
	if err != nil {
		fmt.Println("ERR:   ", err)
		return model.Middleware{StatusCode: 500, Message: "Hệ thống có sự cố"}
	}
	// convert data
	newParking.OwnerId = payload.UserId
	newParking.CreatedAt = time.Now().Format(time.RFC3339)
	newParking.Status	 = "PENDING"
	return model.Middleware{Data: newParking}
}

func (mid *Middleware)BeforeGetOwnerById(c *gin.Context) model.Middleware{
	token, err := utils.GetTokenFromHeader(c)
	if err != nil {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check format token
	checked, _ := mid.Auth.CheckTokenIsTrue(token)
	if checked != true {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check expired token
	checkedExpired, _, _ := mid.Auth.CheckExpiredToken(token)
	if checkedExpired == true {
		return model.Middleware{StatusCode: 400, Message: "Token hết hạn sử dụng"}
	}
	if err != nil {
		return model.Middleware{StatusCode: 500, Message: "Hệ thống có sự cố"}
	}
	// check role
	role, err,_ := mid.Auth.Authorize(token)

	//fmt.Print("role", role, err, rawerr)
	if role != "admin" {
		return model.Middleware{StatusCode: http.StatusServiceUnavailable, Message: "Service không khả dụng"}
	}
	// get ownerId
	return model.Middleware{Data: c.Param("id")}
}

func (mid *Middleware)BeforeVerifyParking(c *gin.Context) model.Middleware {
	token, err := utils.GetTokenFromHeader(c)
	if err != nil {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check format token
	checked, _ := mid.Auth.CheckTokenIsTrue(token)
	if checked != true {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check expired token
	checkedExpired, _, _ := mid.Auth.CheckExpiredToken(token)
	if checkedExpired == true {
		return model.Middleware{StatusCode: 400, Message: "Token hết hạn sử dụng"}
	}
	// check role
	role, err,_ := mid.Auth.Authorize(token)
	if role != "admin" {
		return model.Middleware{StatusCode: 503, Message: "Hệ thống không hỗ trợ dịch vụ này"}
	}
	// convert data
	id := c.Param("id")
	return model.Middleware{Data: id}
}

func (mid *Middleware)AfterGetAllParkings(parkings []model.Parking) model.Middleware {
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

func (mid *Middleware)BeforeCalculateAmountParking(c *gin.Context) model.Middleware{
	token, err := utils.GetTokenFromHeader(c)
	if err != nil {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check format token
	checked, _ := mid.Auth.CheckTokenIsTrue(token)
	if checked != true {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check expired token
	checkedExpired, _, _ := mid.Auth.CheckExpiredToken(token)
	if checkedExpired == true {
		return model.Middleware{StatusCode: 400, Message: "Token hết hạn sử dụng"}
	}
	// check chu so huu bai do
	idParking := c.Param("id")
	var parking model.Parking
	var parkingIface mysql.ParkingDAO
	parkingIface = mid.DAO
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

func (mid *Middleware)BeforeGetAllOwners(c *gin.Context) model.Middleware{
	token, err := utils.GetTokenFromHeader(c)
	if err != nil {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check format token
	checked, _ := mid.Auth.CheckTokenIsTrue(token)
	if checked != true {
		return model.Middleware{StatusCode: 400, Message: "Token không khả dụng"}
	}
	// check expired token
	checkedExpired, _, _ := mid.Auth.CheckExpiredToken(token)
	if checkedExpired == true {
		return model.Middleware{StatusCode: 400, Message: "Token hết hạn sử dụng"}
	}
	role, err,_ := mid.Auth.Authorize(token)
	if role != "admin" {
		return model.Middleware{StatusCode: 503, Message: "Hệ thống không hỗ trợ dịch vụ này"}
	}
	return model.Middleware{}
}