package mysql

import (
	"fmt"
	"github.com/DoHuy/parking_to_easy/config"
	"github.com/jinzhu/gorm"
)

type DAO struct {
	connection *gorm.DB
}

func connectDatabase() (*gorm.DB, error) {
	config := config.GetDatabaseConfig()
	var err error
	var connection *gorm.DB
	connectionInfo := fmt.Sprintf(`%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local`, config.Database.Username, config.Database.Password, config.Database.Address, config.Database.DatabaseName)
	connection, err = gorm.Open("mysql", connectionInfo)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("error in connectDatabase(): %s", err.Error())
	}
	connection.LogMode(true)
	return  connection, nil
}
func NewDAO() (*DAO, error){
	connection, err  := connectDatabase()
	if err != nil {
		return &DAO{}, fmt.Errorf("Lỗi kết nối cơ sở đữ liệu")
	}
	return &DAO{connection: connection,}, nil
}

