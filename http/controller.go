package http

import (
	"encoding/json"
	"fmt"
	"github.com/DoHuy/parking_to_easy/business_logic/auth"
	"github.com/DoHuy/parking_to_easy/config"
	"github.com/DoHuy/parking_to_easy/model"
	"github.com/DoHuy/parking_to_easy/mysql"
	"github.com/DoHuy/parking_to_easy/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"time"
)

type ControllingService struct {
	DAO        	*mysql.DAO
	Middleware 	*MiddleWareService
	Auth		*auth.Auth
}

func NewControllingService(dao *mysql.DAO, middleware *MiddleWareService, auth *auth.Auth) *ControllingService {
	return &ControllingService{
		DAO:        dao,
		Middleware: middleware,
		Auth:		auth,
	}
}

func (con *ControllingService)Options(c *gin.Context) {
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
		c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Content-Type", "application/json")
		c.AbortWithStatus(http.StatusOK)
	}
}
// Bai dau xe
func (con *ControllingService) CreateNewParkingByAdmin(c *gin.Context) {
	// Before create
	var middle model.Middleware
	middle = con.Middleware.BeforeCreateNewParkingByAdmin(c)
	if middle.StatusCode != 0 {
		c.JSON(middle.StatusCode, model.ErrorMessage{Message: middle.Message})
		return
	}
	// implement
	var parkingDAOIface mysql.ParkingDAO
	parkingDAOIface = con.DAO
	//fmt.Println("middle.Datamiddle.Data", middle.Data)
	err := parkingDAOIface.CreateNewParkingByAdmin(middle.Data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorMessage{
			Message:    "Hệ thống có sự cố",
		})
		return
	}
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, model.SuccessMessage{Message: "Thêm mới thành công"})
	return
}

func (con *ControllingService) CreateNewParkingByOwner(c *gin.Context) {
	// before create
	var middle model.Middleware
	middle = con.Middleware.BeforeCreateNewParkingByOwner(c)
	if middle.StatusCode != 0 {
		c.JSON(middle.StatusCode, model.ErrorMessage{Message: middle.Message})
		return
	}
	// implement
	var parkingDAOIface mysql.ParkingDAO
	parkingDAOIface = con.DAO
	err := parkingDAOIface.CreateNewParkingOfOwner(middle.Data)
	if err != nil {
		fmt.Println("ERRRRRR:::::::", err)
		c.JSON(http.StatusInternalServerError, model.ErrorMessage{Message: "Hệ thống có sự cố"})
		return
	}
	c.JSON(http.StatusOK, model.SuccessMessage{Message: "Chia sẻ điểm đậu thành công"})
	return

}

func (con *ControllingService) GetAllApprovedParkings(c *gin.Context) {
	var parkings []model.Parking
	var err error
	// implement
	var parkingDAOIface mysql.ParkingDAO
	parkingDAOIface = con.DAO
	parkings, err = parkingDAOIface.GetAllParking()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorMessage{Message: "Hệ thống có sự cố"})
		return
	}

	if len(parkings) == 0 {
		c.JSON(http.StatusNotFound, model.ErrorMessage{Message: "Không tồn tại bãi đậu nào"})
		return
	}
	var middle model.Middleware
	middle = con.Middleware.AfterGetAllParkings(parkings)
	fmt.Println("middle.Data",middle.Data)
	var rs []model.Parking
	raw, _ := json.Marshal(middle.Data)
	err = json.Unmarshal(raw, &rs)
	if len(rs) == 0 {
		c.JSON(http.StatusOK, model.ErrorMessage{Message: "Không tồn tại bãi đỗ nào"})
		return
	}

	c.JSON(http.StatusOK, middle.Data)
	return

}

func (con *ControllingService) GetAllParkingsOfOwner(c *gin.Context) {
	// Before get all parkings
	var middle model.Middleware
	middle = con.Middleware.BeforeGetDetailUser(c)
	if middle.StatusCode != 0 {
		c.JSON(middle.StatusCode, model.ErrorMessage{Message: middle.Message})
		return
	}
	raw, _ := json.Marshal(middle.Data)
	// implement
	var parkingDAOIface mysql.ParkingDAO
	parkingDAOIface = con.DAO
	owner, err := parkingDAOIface.FindParkingByOwnerId(string(raw))
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, model.ErrorMessage{Message: "Bạn chưa đăng ký chia sẻ điểm đậu xe"})
			return
		}
		fmt.Println("ERR:::::ADASD", err)
		c.JSON(http.StatusInternalServerError, model.ErrorMessage{Message: "Hệ thống có sự cố"})
		return
	}
	c.JSON(http.StatusOK, owner)
	return
}

func (con *ControllingService) RecommendParking(c *gin.Context) {

}
func (con *ControllingService) ModifyParkingByOwner(c *gin.Context) {
	var middle model.Middleware
	middle = con.Middleware.BeforeModifyParkingByOwner(c)
	if middle.StatusCode != 0 {
		c.JSON(middle.StatusCode, model.ErrorMessage{Message: middle.Message})
		return
	}
	var parkingIface mysql.ParkingDAO
	parkingIface = con.DAO
	err := parkingIface.ModifyParking(middle.Data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorMessage{Message:"Hệ thống có sự cố"})
		return
	}
	c.JSON(http.StatusOK, model.SuccessMessage{Message: "Cập nhật thông tin bãi đỗ thành công"})
	return
}
func (con *ControllingService) RemoveParkingOfOwner(c *gin.Context) {
	var middle model.Middleware
	middle = con.Middleware.BeforeDeleteParkingByOwner(c)
	if middle.StatusCode != 0 {
		c.JSON(middle.StatusCode, model.ErrorMessage{Message: middle.Message})
		return
	}
	var parkingIface mysql.ParkingDAO
	parkingIface = con.DAO
	err := parkingIface.DeleteParking(middle.Data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorMessage{Message:"Hệ thống có sự cố"})
		return
	}
	c.JSON(http.StatusOK, model.SuccessMessage{Message: "Xóa bãi đỗ thành công"})
	return
}

func (con *ControllingService) GetAllOwners(c *gin.Context) {
	var middle model.Middleware
	middle = con.Middleware.BeforeGetAllOwners(c)
	if middle.StatusCode != 0 {
		c.JSON(middle.StatusCode, model.ErrorMessage{Message: middle.Message})
		return
	}
	// implement
	var ownerDAOIface mysql.OwnerDAO
	ownerDAOIface = con.DAO
	owners, totalPage, err := ownerDAOIface.GetAllOwners(c.Param("limit"), c.Param("offset"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorMessage{Message: middle.Message})
		return
	}
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, map[string]interface{}{"owners": owners, "totalPage": totalPage})
	return
}

func (con *ControllingService) DisableOwner(c *gin.Context) {
	var middle model.Middleware
	middle = con.Middleware.BeforeDisableOwner(c)
	c.Header("Access-Control-Allow-Origin", "*")
	if middle.StatusCode != 0 {
		c.JSON(middle.StatusCode, model.ErrorMessage{Message: middle.Message})
		return
	}
	var ownerIface = con.DAO
	err := ownerIface.ChangeStatusOwner(middle.Data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorMessage{Message:"Vô hiệu hóa thất bại, hệ thống có sự cố"})
		return
	}
	c.JSON(http.StatusOK, model.SuccessMessage{Message:"Vô hiệu hóa thành công"})
	return
}
func (con *ControllingService) CreateNewTransaction(c *gin.Context) {

}
func (con *ControllingService) GetAllTransactionOfUser(c *gin.Context) {
	var middle model.Middleware
	middle = con.Middleware.BeforeGetAllTransactionOfUser(c)
	if middle.StatusCode != 0 {
		c.JSON(middle.StatusCode, model.ErrorMessage{Message: middle.Message})
		return
	}
	var transactionIface mysql.TransactionDAO
	transactionIface = con.DAO
	transactions, err := transactionIface.FindTransactionOfUser(middle.Data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorMessage{Message: "Hệ thống có sự cố"})
		return
	}
	if len(transactions) == 0 {
		c.JSON(http.StatusNotFound, model.ErrorMessage{Message: "Bạn chưa thực hiện một giao dịch nào"})
		return
	}
	c.JSON(http.StatusOK, transactions)
	return
}
func (con *ControllingService) GetAllTransaction(c *gin.Context) {
	var middle model.Middleware
	middle = con.Middleware.BeforeGetAllTransaction(c)
	if middle.StatusCode != 0 {
		c.JSON(middle.StatusCode, model.ErrorMessage{Message: middle.Message})
		return
	}
	var transactionIface mysql.TransactionDAO
	transactionIface = con.DAO
	transactions, err := transactionIface.FindAllTransaction()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorMessage{Message: "Hệ thống có sự cố"})
		return

	}
	c.JSON(http.StatusOK, transactions)
	return
}
func (con *ControllingService) DeclineTransaction(c *gin.Context) {

}
func (con *ControllingService) AcceptTransaction(c *gin.Context) {

}
func (con *ControllingService) RatingParking(c *gin.Context) {
	var middle model.Middleware
	middle = con.Middleware.BeforeRating(c)
	if middle.StatusCode != 0 {
		c.JSON(middle.StatusCode, model.ErrorMessage{Message: middle.Message})
		return
	}
	var ratingIface mysql.RatingDAO
	ratingIface = con.DAO
	err := ratingIface.CreateVoteOfUser(middle.Data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorMessage{Message: "Hệ thống có sự cố"})
		return
	}
	c.JSON(http.StatusOK, model.SuccessMessage{Message: "Tạo vote thành công"})
	return
}
func (con *ControllingService) UploadFiles(c *gin.Context) {
	// Before Upload
	var midErrorMessage model.Middleware
	token, _ := utils.GetTokenFromHeader(c)
	midErrorMessage = con.Middleware.BeforeUpload(c, token)
	if midErrorMessage.StatusCode != 0 {
		c.JSON(midErrorMessage.StatusCode, model.ErrorMessage{Message: midErrorMessage.Message})
		return
	}
	form, err := c.MultipartForm()
	files := form.File["upload[]"]
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	STATIC_PATH, _ := filepath.Abs("./resource/images")
	envConfig := config.GetEnvironmentConfig()
	var images []string
	for _, file := range files {
		fileID := fmt.Sprintf("(%s)", time.Now().Format(time.RFC3339))
		err := c.SaveUploadedFile(file, filepath.Join(STATIC_PATH, fileID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.ErrorMessage{
				Message: "Hệ thống có sự cố",
			})
			return
		}
		images = append(images, fmt.Sprintf("http://%s:%s/%s", envConfig.Hostname, envConfig.Port, fileID))
	}

	// after upload
	c.JSON(http.StatusOK, map[string][]string{
		"images": images,
	})
	return
}

func (con *ControllingService) FindParkingByID(c *gin.Context) {
	parkingId := c.Param("parkingId")
	var parkingDAOIface mysql.ParkingDAO
	parkingDAOIface = con.DAO
	parking, err := parkingDAOIface.FindParkingByID(parkingId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorMessage{
			Message: "Server error",
		})
		return
	}
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, parking)
	return
}

func (con *ControllingService) GetAllParkings(c *gin.Context) {
	//limit  := c.Param("limit")
	//offset := c.Param("offset")
	var parkingDAOIface mysql.ParkingDAO
	parkingDAOIface = con.DAO
	parkings, err := parkingDAOIface.GetAllParking()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorMessage{
			Message: "Server error",
		})
		return
	}
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, parkings)

	return
}

func (con *ControllingService) Register(c *gin.Context) {
	// Before register
	var middle model.Middleware
	middle = con.Middleware.BeforeRegister(c)
	if middle.StatusCode != 0 {
		c.JSON(middle.StatusCode, model.ErrorMessage{Message: middle.Message})
		return
	}
	//
	var credential model.Credential
	err := utils.BindRawStructToRespStruct(middle.Data, &credential)
	// before create new
	fmt.Println(credential)
	credential.CreatedAt = time.Now().Format(time.RFC3339)
	credential.Password = utils.EncriptPwd(credential.Password)
	credential.Role = "customer"
	//credential
	var credenIface mysql.CredentialDAO
	credenIface = con.DAO
	err = credenIface.CreateCredential(credential)
	//fmt.Println("loi gi vay may: ", newCredential)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorMessage{
			Message: "Server error",
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"message": "Tài khoản được tạo thành công"})
	return
}

func (con *ControllingService) Login(c *gin.Context) {
	var credential model.Credential
	var middle model.Middleware
	//Before Login
	middle = con.Middleware.BeforeLogin(c)
	if middle.StatusCode != 0 {
		c.JSON(middle.StatusCode, model.ErrorMessage{Message: middle.Message})
		return
	}
	err := utils.BindRawStructToRespStruct(middle.Data, &credential)
	//
	var credIface mysql.CredentialDAO
	credIface = con.DAO
	token, err, rawError := con.Auth.Authenticate(credential, credIface)
	if err != nil && rawError == nil {
		c.JSON(http.StatusNotFound, model.ErrorMessage{Message: err.Error(),})
		return
	}
	c.JSON(http.StatusOK, model.LoginMessageResp{Token: token})
	return
}

func (con *ControllingService) VerifyParking(c *gin.Context) {
	////// Before Verify
	// Check Expired token
	var middle model.Middleware
	middle = con.Middleware.BeforeVerifyParking(c)
	if middle.StatusCode != 0 {
		c.JSON(middle.StatusCode, model.ErrorMessage{Message: middle.Message})
		return
	}
	var id string
	raw, _ := json.Marshal(middle.Data)
	err := json.Unmarshal(raw, &id)
	var parkingIface mysql.ParkingDAO
	parkingIface = con.DAO
	parking, err := parkingIface.FindParkingByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorMessage{Message: "Không tồn tại parking này"})
		return
	}
	parking.Status = "APPROVED"
	parking.ModifiedAt = time.Now().Format(time.RFC3339)
	//////
	// Verify parking
	err = parkingIface.ChangStatusParking(parking)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorMessage{Message: "Không xac thực được parking này"})
		return
	}
	c.JSON(http.StatusOK, model.SuccessMessage{Message: "Xác thực thành công"})
	return
	/////
}

func (con *ControllingService) GetOwnerById(c *gin.Context) {
	// before get owner
	var middle model.Middleware
	middle = con.Middleware.BeforeGetOwnerById(c)
	if middle.StatusCode != 0 {
		c.JSON(middle.StatusCode, model.ErrorMessage{Message: middle.Message})
		return
	}
	var ownerIface mysql.OwnerDAO
	ownerIface = con.DAO
	owner, err := ownerIface.FindOwnerById(middle.Data)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, model.ErrorMessage{Message: "User chưa đăng ký làm chủ sở hữu"})
			return
		}
		fmt.Println("", err.Error())
		c.JSON(http.StatusInternalServerError, model.ErrorMessage{Message: "Hệ thống có sự cố"})
		return
	}
	c.JSON(http.StatusOK, owner)
	return
}

func (con *ControllingService) GetAllUsers(c *gin.Context) {
	// Before GetAllUser
	var middle model.Middleware
	middle = con.Middleware.BeforeGetAllUsers(c)
	if middle.StatusCode != 0 {
		c.JSON(middle.StatusCode, model.ErrorMessage{Message: middle.Message})
		return
	}
	//
	fmt.Println("test getalllll")
	var credIface mysql.CredentialDAO
	credIface = con.DAO
	credentials, err := credIface.FindAllCredential(c.Param("limit"), c.Param("offset"))
	fmt.Println("test getalllll22222222222")
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorMessage{Message: "Hệ thống có sự cố", RawMessage: err.Error()})
		return
	}
	if len(credentials) == 0 {
		c.JSON(http.StatusNotFound, model.ErrorMessage{Message: "Không có user nào"})
		return
	}
	err = utils.BindRawStructToRespStruct(credentials, &credentials)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorMessage{Message: "Hệ thống có sự cố", RawMessage: err.Error()})
		return
	}
	c.JSON(http.StatusOK, credentials)
	return
}

func (con *ControllingService) GetDetailUser(c *gin.Context) {
	//Before get detail user
	var middle model.Middleware
	middle = con.Middleware.BeforeGetDetailUser(c)
	if middle.StatusCode != 0 {
		c.JSON(middle.StatusCode, model.ErrorMessage{Message: middle.Message})
		return
	}
	//
	id, _ := json.Marshal(middle.Data)
	var credIface mysql.CredentialDAO
	credIface = con.DAO
	credential, err := credIface.FindCredentialByID(string(id))
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, model.ErrorMessage{Message: "Không tồn tại user này trên hệ thống"})
			return
		}
		c.JSON(http.StatusInternalServerError, model.ErrorMessage{Message: "Hệ thống đang bảo trì"})
		return
	}
	c.JSON(http.StatusOK, credential)
	return
}

func (con *ControllingService) CreateNewOwner(c *gin.Context) {
	// before create new owner
	var middle model.Middleware
	middle = con.Middleware.BeforeCreateNewOwner(c)
	if middle.StatusCode != 0 {
		c.JSON(middle.StatusCode, model.ErrorMessage{Message: middle.Message})
		return
	}
	var ownerIface mysql.OwnerDAO
	ownerIface = con.DAO
	err := ownerIface.CreateNewOwner(middle.Data)
	if err != nil {
		if err.Error() == "Error 1062: Duplicate entry '76' for key 'PRIMARY'" {
			c.JSON(http.StatusInternalServerError, model.ErrorMessage{Message: "Bạn đã đăng kí chia sẻ điểm đậu xe"})
			return
		}
		c.JSON(http.StatusInternalServerError, model.ErrorMessage{Message: middle.Message})
		return
	}
	c.JSON(http.StatusOK, model.SuccessMessage{Message: "Đăng ký bãi đỗ thành công"})
	return
}

func (con *ControllingService) CalculateAmountParking(c *gin.Context) {
	var middle model.Middleware
	middle = con.Middleware.BeforeCalculateAmountParking(c)
	if middle.StatusCode != 0 {
		c.JSON(middle.StatusCode, model.ErrorMessage{Message: middle.Message})
		return
	}
	var ratingIface mysql.RatingDAO
	ratingIface = con.DAO
	var transactionIface mysql.TransactionDAO
	transactionIface = con.DAO
	points, err := transactionIface.CalTotalAmountOfParking(c.Param("id"))
	stars, err 	:= ratingIface.AverageStarsOfParking(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorMessage{Message: "Hệ thống có sự cố"})
		return
	}
	starString := fmt.Sprintf("%.2f", stars)
	c.JSON(http.StatusOK, model.CalculateAmountParkingResp{Points: string(points), Stars: starString})
	return
}

func (con *ControllingService)SaveTokenFireBase(c *gin.Context) {

}

func (con *ControllingService)RemoveToken(c *gin.Context) {

}