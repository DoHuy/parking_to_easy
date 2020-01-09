package model


type CheckingParkingResp struct {
	Message string 	`json:"message"`
}
type CredentialResp struct {
	ID			int			`json:"id"`
	Username	string		`json:"username"`
	Email		string		`json:"email"`
	Points		int			`json:"points"`
	CreatedAt	string		`json:"created_at"`
	ModifiedAt	string		`json:"modified_at"`
	DeletedAt	string		`json:"deleted_at"`
}

type ModifyingParkingResp struct {
	ParkingName		string		`json:"parking_name"`
	Properties		string		`json:"properties,omitempty"`
	Address			string		`json:"address,omitempty"`
	KindOf			bool		`json:"kind_of,omitempty"`
	ParkingImages 	string		`json:"parking_images,omitempty"`
	Payment			string		`json:"payment,omitempty"`
	Longitude		string		`json:"longitude,omitempty"`
	Latitude		string		`json:"latitude,omitempty"`
	Capacity		string		`json:"capacity,omitempty"`
	BlockAmount		int			`json:"block_amount,omitempty"`
	CreatedAt		string		`json:"created_at,omitempty"`
	ModifiedAt 		string		`json:"modified_at,omitempty"`
	Describe		string		`json:"describe,omitempty"`
	Status			string		`json:"status,omitempty"`
}

type ParkingResp struct {
	ParkingName		string		`json:"parking_name"`
	Properties		string		`json:"properties,omitempty"`
	Address			string		`json:"address,omitempty"`
	KindOf			bool		`json:"kind_of,omitempty"`
	ParkingImages 	string		`json:"parking_images,omitempty"`
	Payment			string		`json:"payment,omitempty"`
	Longitude		string		`json:"longitude,omitempty"`
	Latitude		string		`json:"latitude,omitempty"`
	Capacity		string		`json:"capacity,omitempty"`
	BlockAmount		int			`json:"block_amount,omitempty"`
	CreatedAt		string		`json:"created_at,omitempty"`
	ModifiedAt 		string		`json:"modified_at,omitempty"`
	Describe		string		`json:"describe,omitempty"`
	Status			string		`json:"status,omitempty"`
}

type RegisterResp struct {
	Message	string		`json:"message"`
}

type LoginMessageResp struct {
	Token string	`json:"token"`
}

type ErrorMessage struct {
	Message			string		`json:"error_message"`
	RawMessage		string		`json:"raw_message,omitempty"`
}

type VerifyingParkingResp struct {
	Message			string		`json:"message"`
	Parking			ParkingResp	`json:"parking"`
}
type Middleware struct {
	StatusCode	int			`json:"status_code"`
	Message		string		`json:"message"`
	Data		interface{}	`json:"data"`
}

// schema create parking by admin
type NewParkingByAdmin struct {
	ParkingName			string		`json:"parkingName,omitempty"`
	Properties			string		`json:"properties,omitempty"`
	Address				string		`json:"address,omitempty"`
	KindOf				string		`json:"kindOf,omitempty"`
	ParkingImages 		string		`json:"parkingImages,omitempty"`
	Payment				string		`json:"payment,omitempty"`
	Longitude			string		`json:"longitude,omitempty"`
	Latitude			string		`json:"latitude,omitempty"`
	Capacity			string		`json:"capacity,omitempty"`
	BlockAmount			interface{}	`json:"blockAmount,omitempty"`
	CertificateOfLand	string		`json:"certificateOfland,omitempty"`
	Describe			string		`json:"describe,omitempty"`
	Status				string		`json:"status,omitempty"`
	OwnerId				int			`json:"ownerId,omitempty"`
	CreatedAt			string		`json:"created_at,omitempty"`
}
//

type SuccessMessage struct {
	Message	string		`json:"message"`
}

type CalculateAmountParkingResp struct {
	Points	string	`json:"points"`
	Stars	string	`json:"stars"`
}


