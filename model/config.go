package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type Config struct {
	Database				Database		`json:"database,omitempty"`
	SecretKey   			SecretKey		`json:"secret_key,omitempty"`
	MaxUploadedFileSize		int64			`json:"max_uploaded_file_size,omitempty"`
	Environment				Environment		`json:"environment"`
}

type Environment struct {
	Hostname	string
	Port		string
}

type Database struct {
	Username    	 string `json:"username"`
	Password	  	 string `json:"password"`
	DatabaseName 	 string `json:"database_name"`
	Address  	 	 string `json:"address"`
}

type SecretKey string

func getPathTargetedFile(filename string) (string, error){
	path, err := filepath.Abs("./")
	if err != nil {
		return "", fmt.Errorf("Lỗi đọc tệp cấu hình %v", err)
	}
	return filepath.Join(path, filename), nil

}

func GetEnvironmentConfig() Environment {
	var config Config
	configPath, _ := getPathTargetedFile("config.json")
	file, _ := ioutil.ReadFile(configPath)
	json.Unmarshal([]byte(file), &config)
	return config.Environment
}
func getDatabaseConfig() Config{
	var config Config
	configPath, _ := getPathTargetedFile("config.json")
	file, _ := ioutil.ReadFile(configPath)
	json.Unmarshal([]byte(file), &config)
	return config
}

func GetMaxUploadedFileSize() int64 {
	var config Config
	configPath, _ := getPathTargetedFile("config.json")
	file, _ := ioutil.ReadFile(configPath)
	json.Unmarshal([]byte(file), &config)
	return config.MaxUploadedFileSize
}
func GetSecretKey() SecretKey {
	var config Config
	filepath, _ := getPathTargetedFile("config.json")
	file, _ := ioutil.ReadFile(filepath)
	json.Unmarshal([]byte(file), &config)
	return config.SecretKey
}
