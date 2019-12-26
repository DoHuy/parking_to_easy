package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Middleware struct {
	Connection	*gorm.DB
}

func NewMiddleware(instance *gorm.DB) *Middleware {
	return &Middleware{instance}
}

func (this *Middleware) ValidateParkingCreating(c *gin.Context) {
	c.Next()
}
