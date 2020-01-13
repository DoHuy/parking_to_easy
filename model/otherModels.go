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
	BlockAmount			int			`json:"blockAmount,omitempty"`
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
	Points	int		`json:"points"`
	Stars	float64	`json:"stars"`
}

type CreatingTransactionInput struct {
	Payload		Payload		`json:"payload"`
	Transaction Transaction	`json:"transaction"`
}

type GettingTransactionDetailResp struct {
	ID				int			`json:"transactionId"`
	StartTime		string		`gorm:"column:startTime" json:"startTime,omitempty"`
	EndTime			string		`gorm:"column:endTime" json:"endTime,omitempty"`
	Licence			string		`json:"licence"`
	Address			string		`json:"address"`
	Amount			int			`json:"amount"`
	Status			int			`json:"status"`
	CreatedAt		string		`json:"created_at"`
	UserPhoneNumber	string		`gorm:"column:userPhoneNumber" json:"userPhoneNumber,omitempty"`
	HostPhoneNumber	string		`gorm:"column:hostPhoneNumber" json:"hostPhoneNumber,omitempty"`

}

type GetTransactionOfUserWithStatusInput struct {
	Status	int	`json:"status"`
	UserId	int	`json:"userId"`
}

type GetTransactionOfOwnerWithStatusInput struct {
	ParkingId	int `json:"parkingId"`
	Status		int	`json:"status"`
}

type ChangingStateTransactionInput struct {
	Status			int		`json:"status"`
	TransactionId	int		`json:"transactionId"`
	CredentialId	int		`json:"credentialId"`
}


type DataStruct struct {
	CredentialId string		`json:"credentialId"`
	Status		 string		`json:"status"`
	ModifiedAt	 string		`json:"modified_at"`

}

type VerifyingParkingInput struct {
	Status		string		`json:"status"`
	ID			string		`json:"id"`
	ModifiedAt	string		`json:"modified_at"`
}

type GettingAllOwnersOutput struct {
	Owner	Owner		`json:"owner"`
	Stars	float64		`json:"stars"`
	Votes	int			`json:"votes"`
}

type TransactionTopicInput struct {
	TransactionId	int			`json:"transactionId"`
	CredentialId	int			`json:"credentialId"`
	OwnerId			int			`json:"ownerId"`
	OwnerTokensList []string	`json:"ownerTokensList"`
	UserTokensList  []string	`json:"UserTokensList"`
}

type AnalysisInput struct {
	Start	int64	`json:"start"`
	End		int64	`json:"end"`
}

type AnalysisOutput struct {
	Finished 	 int	`json:"finished"`
	Canceled	 int	`json:"canceled"`
}