package main

import (
	"github.com/gin-gonic/gin"
	"parking_to_esay/router"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	r.Run(":8085")
}