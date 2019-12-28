package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var connection *gorm.DB

type ErrorMessage struct {
	Message		string		`json:"error_message"`
}
type UserDevice struct {
	ID				int 		`gorm:"column:id" json:"id"`
	DeviceToken		string		`gorm:"type:text" json:"deviceToken"`
	Credential		Credential  `json:"credential,omitempty"`
	CredentialId	int			`gorm:"column:credentialId" json:"credentialId"`
}

type Credential struct {
	ID			int			`gorm:"column:id" json:"id,omitempty"`
	Username	string		`gorm:"column:username" json:"username"`
	Password	string		`gorm:"column:password" json:"password,omitempty"`
	Email		string		`gorm:"column:email" json:"email"`
	Points		int			`gorm:"column:points" json:"points"`
	Role		string		`gorm:"column:role" json:"role"`
	Token		string		`gorm:"column:token" json:"token"`
	Expired		time.Time	`gorm:"column:expired" json:"expired"`
	CreatedAt	time.Time	`gorm:"column:created_at" json:"created_at"`
	ModifiedAt	time.Time	`gorm:"column:modified_at" json:"modified_at"`
	DeletedAt	time.Time	`gorm:"column:deleted_at" json:"deleted_at"`
}

type Owner struct {
	CredentialId		int			`gorm:"column:credentialId" json:"credentialId"`
	Credential			Credential	`json:"credential,omitempty"`
	FullName			string		`gorm:"column:fullName"  json:"fullName"`
	CMNDImage			string		`gorm:"column:cmndImage" json:"cmndImage"`
	CertificateOfLand	string		`gorm:"column:certificateOfLand" json:"certificateOfLand"`
	Address				string		`gorm:"column:address" json:"address"`
	PhoneNumber			string		`gorm:"column:phoneNumber" json:"phoneNumber"`
	Status				string		`gorm:"column:status" json:"status"`
	CreatedAt			time.Time	`gorm:"column:created_at" json:"created_at"`
	ModifiedAt			time.Time	`gorm:"column:modified_at" json:"modified_at"`
}

type Parking struct {
	ID				int			`gorm:"column:id" json:"id"`
	ParkingName		string		`gorm:"column:parkingName" json:"parkingName"`
	Properties		string		`gorm:"column:properties" json:"properties"`
	Address			string		`gorm:"column:address" json:"address"`
	KindOf			bool		`gorm:"column:kindOf" json:"kindOf"`
	ParkingImages 	string		`gorm:"column:parkingImages" json:"parkingImages"`
	Payment			string		`gorm:"column:payment" json:"payment"`
	Longitude		string		`gorm:"column:longitude" json:"longitude"`
	Latitude		string		`gorm:"column:latitude" json:"latitude"`
	Capacity		string		`gorm:"column:capacity" json:"capacity"`
	BlockAmount		int			`gorm:"column:blockAmount" json:"blockAmount"`
	OwnerId			int			`gorm:"column:ownerId" json:"ownerId"`
	Owner 			Owner		`json:"owners,omitempty"`
	CreatedAt		time.Time	`gorm:"column:created_at" json:"created_at"`
	ModifiedAt 		time.Time	`gorm:"column:modified_at" json:"modified_at"`
	DeletedAt 		time.Time	`gorm:"column:deleted_at" json:"deleted_at"`
	Describe		string		`gorm:"column:describe" json:"describe"`

}

type Transaction struct {
	CredentialId	int				`gorm:"column:credentialId" json:"credentialId"`
	ParkingId		int				`gorm:"column:parkingId" json:"parkingId"`
	Credential		Credential		`json:"credential,omitempty"`
	Parking			Parking			`json:"parking,omitempty"`
	LiencePlate		string			`gorm:"column:liencePlate" json:"liencePlate"`
	Session			time.Duration	`gorm:"column:session" json:"session"`
	StartTime		time.Time		`gorm:"column:startTime" json:"startTime"`
	EndTime			time.Time		`gorm:"column:endTime" json:"endTime"`
	Amount			int				`gorm:"column:amount" json:"amount"`
	Status			string			`gorm:"column:status" json:"status"`
	ReasonMsg		string			`gorm:"column:reasonMsg" json:"reasonMsg"`
	CreatedAt		time.Time		`gorm:"column:created_at" json:"created_at"`
	ModifiedAt		time.Time		`gorm:"column:modified_at" json:"modified_at"`

}

func ConnectDatabase() (*gorm.DB, error) {
	config := getDatabaseConfig()
	var err error
	connectionInfo := fmt.Sprintf(`%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local`, config.Database.Username, config.Database.Password, config.Database.Address, config.Database.DatabaseName)
	connection, err = gorm.Open("mysql", connectionInfo)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("error in connectDatabase(): %s", err.Error())
	}
	return connection, nil
}

// Thao tác với Credential model
func FindCredentialByID(id int) Credential {
	var credential Credential
	connection.Raw("SELECT username, password, email, points, role, token, expired, created_at, modified_at, deleted_at FROM credentials where id=?", id).Scan(&credential)
	return credential
}

func FindAllCredential(limit, offset int) []Credential {
	var credentials []Credential
	rows, _ := connection.Raw("SELECT username, password, email, points, role, token, expired, created_at, modified_at, deleted_at FROM credentials LIMIT ? OFFSET ?", limit, offset).Rows()
	defer rows.Close()
	index := 0
	for rows.Next() {
		rows.Scan(&credentials[index])
		index++
	}
	return credentials
}

func CreateCredential(newUser Credential) error {
	var credential Credential
	err := connection.Table("credentials").Create(&newUser).Scan(&credential).Error

	if err != nil {
		return fmt.Errorf("Loi truy van database: %v", err.Error())
	}
	return nil
}
////////////////////////////////////////////////////

// Thao tac vs Parking Model
func CreateParking(newParking Parking) Parking {
	var parking Parking
	connection.Table("parkings").Create(&newParking).Scan(&parking)
	return parking
}

func ModifyParking(updatedParking Parking) Parking {
	var parking Parking
	connection.Model(&parking).Updates(updatedParking).Scan(&parking)
	return parking
}

func FindParkingByID(id string) ([]Parking, error) {
	var parkings []Parking
	err := connection.Table("parkings").Raw("SELECT * FROM parkings INNER JOIN owners ON owners.credentialId=parkings.ownerId WHERE parkings.id=?", id).Scan(&parkings).Error
	if err != nil {
		return nil, fmt.Errorf("Loi truy van database: %v", err.Error())
	}
	for i, _ := range parkings {
		err := connection.Table("parkings").Raw("SELECT*from owners WHERE credentialId=?", parkings[i].OwnerId).Scan(&parkings[i].Owner).Error

		if err != nil {
			return nil, fmt.Errorf("Loi truy van database: %v", err.Error())
		}
	}

	return parkings, nil
}
//
func GetAllParking(limit, offset string) ([]Parking, error){
	var parkings []Parking
	err := connection.Table("parkings").Raw("SELECT * FROM parkings LIMIT ? OFFSET ?", limit, offset).Scan(&parkings).Error
	if err != nil {
		return nil, fmt.Errorf("Loi truy van database: %v", err.Error())
	}
	for i, _ := range parkings {
		err := connection.Table("parkings").Raw("SELECT*from owners WHERE credentialId=?", parkings[i].OwnerId).Scan(&parkings[i].Owner).Error

		if err != nil {
			return nil, fmt.Errorf("Loi truy van database: %v", err.Error())
		}
	}
	return parkings, nil
}

//////////////////////////////////////////////////
