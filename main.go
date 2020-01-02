package main

import (
	"github.com/DoHuy/parking_to_easy/mysql"
	"github.com/DoHuy/parking_to_easy/redis"
	"github.com/DoHuy/parking_to_easy/router"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"path/filepath"
)

func main() {
	r := gin.Default()
	static_path, _ := filepath.Abs("./")
	r.Use(static.Serve("/", static.LocalFile(filepath.Join(static_path, "resource/images"), false)))
	connection, err := mysql.ConnectDatabase()
	redisPool := redis.InitPoolConnectionRedis()
	if err != nil {
		panic("Lỗi kết nối database !")
	}
	router.InitRouter(r, connection, redisPool)
	r.Run(":8085")
}