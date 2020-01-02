package model

type JWTOfUser struct {
	Key 	string		`json:"key"`
	Jwt	    string		`json:"jwt"`
}

type DevicesOfUser struct {
	Key		string			`json:"key"`
	Devices []string 		`json:"devices"`
}