package model

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ErrorMessage struct {
	Message		string		`json:"error_message"`
}

type UserDevice struct {
	ID				int 		`gorm:"column:id" json:"id,omitempty"`
	DeviceToken		string		`gorm:"type:text" json:"deviceToken,omitempty"`
	Credential		Credential  `json:"credential,omitempty"`
	CredentialId	int			`gorm:"column:credentialId" json:"credentialId,omitempty"`
}

type Credential struct {
	ID			int			`gorm:"column:id" json:"id,omitempty"`
	Username	string		`gorm:"column:username" json:"username,omitempty"`
	Password	string		`gorm:"column:password" json:"password,omitempty"`
	Email		string		`gorm:"column:email" json:"email,omitempty"`
	Points		int			`json:"points,omitempty"`
	Role		string		`gorm:"column:role" json:"role,omitempty"`
	Token		string		`gorm:"column:token" json:"token,omitempty"`
	Expired		string		`gorm:"column:expired" json:"expired,omitempty"`
	CreatedAt	string		`gorm:"column:created_at" json:"created_at,omitempty"`
	ModifiedAt	string		`gorm:"column:modified_at" json:"modified_at,omitempty"`
	DeletedAt	string		`gorm:"column:deleted_at" json:"deleted_at,omitempty"`
}

type Owner struct {
	CredentialId		int			`gorm:"column:credentialId" json:"credentialId,omitempty"`
	Credential			Credential	`json:"credential,omitempty"`
	FullName			string		`gorm:"column:fullName"  json:"fullName,omitempty"`
	PhoneNumber			string		`gorm:"column:phoneNumber" json:"phoneNumber,omitempty"`
	Address				string		`gorm:"column:address" json:"address,omitempty"`
	CMNDImage			string		`gorm:"column:cmndImage" json:"cmndImage,omitempty"`
	Status				string		`gorm:"column:status" json:"status,omitempty"`
	Parkings			[]Parking	`json:"parkings, omitempty"`
	ParkingId			int			`gorm:"column:parkingId" json:"parkingId,omitempty"`
	CreatedAt			string		`gorm:"column:created_at" json:"created_at,omitempty"`
	ModifiedAt			string		`gorm:"column:modified_at" json:"modified_at,omitempty"`
}

type Parking struct {
	ID					int			`gorm:"column:id" json:"id, omitempty"`
	ParkingName			string		`gorm:"column:parkingName" json:"parkingName, omitempty"`
	Properties			string		`gorm:"column:properties" json:"properties, omitempty"`
	Address				string		`gorm:"column:address" json:"address, omitempty"`
	KindOf				bool		`gorm:"column:kindOf" json:"kindOf, omitempty"`
	ParkingImages 		string		`gorm:"column:parkingImages" json:"parkingImages, omitempty"`
	Payment				string		`gorm:"column:payment" json:"payment, omitempty"`
	Longitude			string		`gorm:"column:longitude" json:"longitude, omitempty"`
	Latitude			string		`gorm:"column:latitude" json:"latitude, omitempty"`
	Capacity			string		`gorm:"column:capacity" json:"capacity, omitempty"`
	BlockAmount			int			`gorm:"column:blockAmount" json:"blockAmount, omitempty"`
	CertificateOfLand	string		`gorm:"column:certificateOfLand" json:"certificateOFLand, omitempty"`
	CreatedAt			string		`gorm:"column:created_at" json:"created_at, omitempty"`
	ModifiedAt 			string		`gorm:"column:modified_at" json:"modified_at, omitempty"`
	DeletedAt 			string		`gorm:"column:deleted_at" json:"deleted_at, omitempty"`
	Describe			string		`gorm:"column:describe" json:"describe, omitempty"`
	Status				string		`gorm:"column:status" json:"status, omitempty"`
}

type Transaction struct {
	CredentialId	int				`gorm:"column:credentialId" json:"credentialId,omitempty"`
	ParkingId		int				`gorm:"column:parkingId" json:"parkingId,omitempty"`
	Credential		Credential		`json:"credential,omitempty,omitempty"`
	Parking			Parking			`json:"parking,omitempty,omitempty"`
	LiencePlate		string			`gorm:"column:liencePlate" json:"liencePlate,omitempty"`
	Session			int64			`gorm:"column:session" json:"session,omitempty"`
	StartTime		string			`gorm:"column:startTime" json:"startTime,omitempty"`
	EndTime			string			`gorm:"column:endTime" json:"endTime,omitempty"`
	Amount			int				`gorm:"column:amount" json:"amount,omitempty"`
	Status			string			`gorm:"column:status" json:"status,omitempty"`
	ReasonMsg		string			`gorm:"column:reasonMsg" json:"reasonMsg,omitempty"`
	CreatedAt		string			`gorm:"column:created_at" json:"created_at,omitempty"`
	ModifiedAt		string			`gorm:"column:modified_at" json:"modified_at,omitempty"`

}

type Payload struct {
	UserId  int     `json:"userId"`
	Role    string  `json:"role"`
	Expired string  `json:"expired"`
}

// Encode generates a jwt.
type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}


//////////////////////////////////////////////////
