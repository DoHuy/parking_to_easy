package controller

import (
	"encoding/json"
	"fmt"
	"github.com/DoHuy/parking_to_easy/auth"
	"github.com/DoHuy/parking_to_easy/config"
	"github.com/DoHuy/parking_to_easy/mysql"
	"github.com/DoHuy/parking_to_easy/redis"
	redisLib "github.com/gomodule/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"github.com/DoHuy/parking_to_easy/model"
	"path/filepath"
	"time"
)

type Controller struct {
	Connection	*gorm.DB
	RedisPool	*redisLib.Pool
}

func getBodyRequest(c *gin.Context) []byte{
	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	return buf[0:num]
}
func NewController(instance *gorm.DB, redisPool *redisLib.Pool) *Controller{
	return &Controller{Connection: instance, RedisPool: redisPool,}
}

func (this *Controller) CreateNewParking(c *gin.Context) {

	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	var parking interface{}
	json.Unmarshal(buf[0:num], &parking)

	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, model.ErrorMessage{
	//		Message: "Server error",
	//	})
	//	return
	//}
	c.JSON(http.StatusOK, parking)
	return
}

func (this *Controller) CreateNewParkingByAdmin(c *gin.Context) {

	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	var newParkingInput interface{}
	json.Unmarshal(buf[0:num], &newParkingInput)
	err, parking := mysql.CreateNewParkingByAdmin(newParkingInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorMessage{
			Message: "Server error",
		})
		return
	}
	c.JSON(http.StatusOK, parking)
	return
}

func (this *Controller) UploadFiles(c *gin.Context) {

	// Before Upload
	// check token valid
	authHeader := c.Request.Header.Get("Authorization")
	_, err := auth.CheckTokenIsTrue(authHeader)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorMessage{Message: err.Error()})
		return
	}
	//
	form,err := c.MultipartForm()
	files := form.File["upload[]"]
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	STATIC_PATH, _ := filepath.Abs("./resource/images")
	envConfig := config.GetEnvironmentConfig()
	var images []string
	for _, file := range files {

		if file.Size >= config.GetMaxUploadedFileSize() {
			c.JSON(http.StatusBadRequest, model.ErrorMessage{Message: "Ảnh được tải lên không được vượt quá 10 MB"})
			return
		}

		fileID := fmt.Sprintf("(%s)%s",time.Now().Format(time.RFC3339), file.Filename)
		err := c.SaveUploadedFile(file, filepath.Join(STATIC_PATH, fileID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.ErrorMessage{
				Message: "Server error",
			})
			return
		}
		images = append(images, fmt.Sprintf("http://%s:%s/%s", envConfig.Hostname, envConfig.Port, fileID))
	}

	c.JSON(http.StatusOK, map[string][]string {
		"images": images,
	})
	return
}

func (this *Controller) FindParkingByID(c *gin.Context) {
	parkingId := c.Param("parkingId")
	parking, err := mysql.FindParkingByID(parkingId)
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

func (this *Controller) GetAllParkings(c *gin.Context) {
	limit  := c.Param("limit")
	offset := c.Param("offset")
	parkings, err := mysql.GetAllParking(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorMessage{
			Message: "Server error",
		})
		return
	}
	c.JSON(http.StatusOK, parkings)
	return
}

func (this *Controller) Register(c *gin.Context) {
	buffer := getBodyRequest(c)
	var credential model.Credential
	json.Unmarshal(buffer, &credential)
	// before create new
	credential.CreatedAt = time.Now().Format(time.RFC3339)
	credential.Expired = time.Now().Format(time.RFC3339)
	//
	//c.JSON(http.StatusOK, parking)
	err, newCredential := mysql.CreateCredential(credential)
	fmt.Println("loi gi vay may: ", err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorMessage{
			Message: "Server error",
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"message": "Tài khoản được tạo thành công", "credential":newCredential})
	return
}

func (this *Controller) Login(c *gin.Context) {
	var credential model.Credential
	//validate user before login
	buffer := getBodyRequest(c)
	json.Unmarshal(buffer, &credential)
	var err error
	fmt.Println("Credential:::::::::::", credential)
	credential, err = mysql.FindCredentialByNameAndPassword(credential.Username, credential.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorMessage{Message: "username hoặc password không đúng"})
		return
	}
	fmt.Println("Credential TOKEN:::::::::::", credential.Token)
	//Tạo mới token lưu vào redis và mysql
	secretKey  := string(config.GetSecretKey())
	if len(credential.Token) <= 0 {
		fmt.Println("SECRET.........", secretKey)
		jwt, err := auth.Encode(model.Payload{UserId: credential.ID, Role: credential.Role, Expired: time.Now().Format(time.RFC3339),}, secretKey)
		fmt.Println("JWT:,,,,,,,,,,,,,", jwt)
		err = redis.SetJWTTokenToRedis(credential.Username, jwt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.ErrorMessage{Message: "Hệ thống có sự cố"})
			return
		}
		c.JSON(http.StatusOK, map[string]string{"token": jwt})
		return
	} else {
		// kiem tra expired token
		var jwt model.JWTOfUser
		jwt, err = redis.GetJWTTokenFromRedis(credential.Username)
		//decodedData, err := auth.Decode(jwt.Jwt, secretKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.ErrorMessage{Message: "Hệ thống có sự cố"})
			return
		}
		//json.Unmarshal(decodedData, &payload)
		if checked, err := auth.CheckExpiredToken(jwt.Jwt); err != nil {
			c.JSON(http.StatusInternalServerError, model.ErrorMessage{Message: err.Error()})
			return
		} else {
			if !checked {
				// tao moi token len redis va tra e token moi
				jwt, err := auth.Encode(model.Payload{UserId: credential.ID, Role: credential.Role, Expired: time.Now().Format(time.RFC3339),}, secretKey)
				err = redis.SetJWTTokenToRedis(credential.Username, jwt)
				if err != nil {
					c.JSON(http.StatusInternalServerError, model.ErrorMessage{Message: "Sự cố hệ thống"})
					return
				}
				c.JSON(http.StatusOK, map[string]string{"token": jwt})
				return
			}
		}
		//
		c.JSON(http.StatusOK, map[string]string{"token": jwt.Jwt})
		return
	}

}

func (this *Controller) CheckParking(c *gin.Context) {

}

func (this *Controller) ModifyParking(c *gin.Context) {

}

func (this *Controller) GetAllUsers(c *gin.Context) {

}

func (this *Controller) GetDetailOwner(c *gin.Context) {

}