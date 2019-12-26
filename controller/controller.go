package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type Controller struct {
	Connection	*gorm.DB
}

func NewController(instance *gorm.DB) *Controller{
	return &Controller{Connection: instance}
}

func (this *Controller) CreateNewParking(c *gin.Context) {

	b:=map[string]string{
		"a":"b",
	}
	c.JSON(http.StatusOK, b)
	return
}