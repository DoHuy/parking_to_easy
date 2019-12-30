package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"github.com/DoHuy/parking_to_easy/model"
)

type Middleware struct {
	Connection	*gorm.DB
}

func NewMiddleware(instance *gorm.DB) *Middleware {
	return &Middleware{instance}
}

func (this *Middleware) ValidateParkingCreating(c *gin.Context) {
	//this.Connection
	//buf := make([]byte, 1024)
	//num, _ := c.Request.Body.Read(buf)
	//reqBody := string(buf[0:num])
	//c.JSON(http.StatusOK, reqBody)
	c.Next()
}

func (this *Middleware) BeforeUploadFiles(c *gin.Context){
	// check token hop le

	// check file size, file type
	form, err := c.MultipartForm()
	files := form.File["upload[]"]
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	for _, file := range files {
		//mimeType := file.Header["Content-Type"][0]
		//fmt.Println(file.Size,  model.GetMaxUploadedFileSize())
		if file.Size <= model.GetMaxUploadedFileSize() {
			c.JSON(http.StatusBadRequest, model.ErrorMessage{Message: "Ảnh được tải lên không được vượt quá 10 MB"})
			return
		}

	}
	fmt.Println("sadgdjajgfafjgwkfjgwer")
	c.Next()
}