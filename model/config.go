package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type Config struct {
	Database	Database	`json:"database"`
}

type Database struct {
	Username    	 string `json:"username"`
	Password	  	 string `json:"password"`
	DatabaseName 	 string `json:"database_name"`
	Address  	 	 string `json:"address"`
}

func getPathTargetedFile(filename string) (string, error){
	path, err := filepath.Abs("./")
	if err != nil {
		return "", fmt.Errorf("Lỗi đọc tệp cấu hình %v", err)
	}
	return filepath.Join(path, filename), nil

}

func getDatabaseConfig() Config{
	var config Config
	configPath, _ := getPathTargetedFile("config.json")
	file, _ := ioutil.ReadFile(configPath)
	json.Unmarshal([]byte(file), &config)
	return config
}
