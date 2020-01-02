package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"github.com/DoHuy/parking_to_easy/model"
)

func getPathTargetedFile(filename string) (string, error){
	path, err := filepath.Abs("./")
	if err != nil {
		return "", fmt.Errorf("Lỗi đọc tệp cấu hình %v", err)
	}
	return filepath.Join(path, filename), nil

}

func GetEnvironmentConfig() model.Environment {
	var config model.Config
	configPath, _ := getPathTargetedFile("config.json")
	file, _ := ioutil.ReadFile(configPath)
	json.Unmarshal([]byte(file), &config)
	return config.Environment
}
func GetDatabaseConfig() model.Config{
	var config model.Config
	configPath, _ := getPathTargetedFile("config.json")
	file, _ := ioutil.ReadFile(configPath)
	json.Unmarshal([]byte(file), &config)
	return config
}

func GetMaxUploadedFileSize() int64 {
	var config model.Config
	configPath, _ := getPathTargetedFile("config.json")
	file, _ := ioutil.ReadFile(configPath)
	json.Unmarshal([]byte(file), &config)
	return config.MaxUploadedFileSize
}
func GetSecretKey() model.SecretKey {
	var config model.Config
	filepath, _ := getPathTargetedFile("config.json")
	file, _ := ioutil.ReadFile(filepath)
	json.Unmarshal([]byte(file), &config)
	return config.SecretKey
}

func GetConfigRedis() model.Redis {
	var config model.Config
	filepath, _ := getPathTargetedFile("config.json")
	file, _ := ioutil.ReadFile(filepath)
	json.Unmarshal([]byte(file), &config)
	return config.Redis
}

func GetTokenExpired() int64 {
	var config model.Config
	filepath, _ := getPathTargetedFile("config.json")
	file, _ := ioutil.ReadFile(filepath)
	json.Unmarshal([]byte(file), &config)
	return config.TokenExpired
}