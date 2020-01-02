package model

type Config struct {
	Database				Database		`json:"database, omitempty"`
	SecretKey   			SecretKey		`json:"secret_key, omitempty"`
	MaxUploadedFileSize		int64			`json:"max_uploaded_file_size, omitempty"`
	Environment				Environment		`json:"environment, omitempty"`
	Redis					Redis			`json:"redis, omitempty"`
	TokenExpired			int64			`json:"token_expired, omitempty"`
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

type Redis struct {
	Address		string		`json:"Address"`
	Topic 		[]string	`json:"topic"`
	Networks	string		`json:"networks"`
}

