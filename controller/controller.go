package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"github.com/DoHuy/parking_to_easy/model"
	"path/filepath"
	"time"
)

type Controller struct {
	Connection	*gorm.DB
}

func NewController(instance *gorm.DB) *Controller{
	return &Controller{Connection: instance}
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
	err, parking := model.CreateNewParkingByAdmin(newParkingInput)
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

	form,err := c.MultipartForm()
	files := form.File["upload[]"]
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	STATIC_PATH, _ := filepath.Abs("./resource/images")
	envConfig := model.GetEnvironmentConfig()
	var images []string
	for _, file := range files {

		if file.Size >= model.GetMaxUploadedFileSize() {
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
	parking, err := model.FindParkingByID(parkingId)
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
	parkings, err := model.GetAllParking(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorMessage{
			Message: "Server error",
		})
		return
	}
	c.JSON(http.StatusOK, parkings)
	return
}

func (this *Controller) CreateNewCredential(c *gin.Context) {
	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	var credential model.Credential
	json.Unmarshal(buf[0:num], &credential)
	//c.JSON(http.StatusOK, parking)
	err, newCredential := model.CreateCredential(credential)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorMessage{
			Message: "Server error",
		})
		return
	}
	c.JSON(http.StatusOK, newCredential)
	return
}