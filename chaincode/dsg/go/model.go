package main


type GoldBarData struct {
	BarLocation                 string    `json:"barLocation"`
	BarPurity                   float64   `json:"barPurity"`
	BarSerialNumber             int       `json:"barSerialNumber"`
	BarHallmarkVerified         bool      `json:"barHallmarkVerfied"`
	BarRefiner                  string    `json:"barRefiner"`
	BarWeightInGms              float64   `json:"barWeightInGms"`
	BarID                       string    `json:"barId"`
}


type GoldReserve struct {
	TotalGoldAmount       float64    `json:"totalGoldAmount"`
	BarCount              int        `json:"barCount"`
	BarRefiner            string     `json:"barRefiner"`
}


type PaymentCardDetail struct {
	PaymentCardID       string     `json:"paymentCardId"`
	PaymentProviderID   string     `json:"paymentProviderId"`
}


type PurchaseGoldBar struct {
	Username     				 string     		 `json:"username"`
	CreateAt            		 string      		 `json:"createAt"`
	OrderID             		 string     		 `json:"orderId"`
	UserID            			 string     		 `json:"userId"`
	OrderDate         		     string      		 `json:"orderDate"`
	WeightInGms                  float64    		 `json:"weightInGms"`
	Amount                       int         		 `json:"amount"`
	Symbol                       string    		     `json:"symbol"`
	PaymentStatus                string      		 `json:"paymentStatus"`
	KYCID                        string      		 `json:"kycId"`
	KYCVerified                  bool       		 `json:"kycVerified"`
	KYCLastVerified              string      		 `json:"kycLastVerfied"`
	PaymentProviderID            string      		 `json:"paymentProviderId"`
	PaymentCardID                string     		 `json:"paymentCardId"`
	PaymentCardDetails           []PaymentCardDetail `json:"paymentCardDetails"`
	AuthenticatedShopperID       string   			 `json:"authenticatedShopperId"`
}

type SellGoldBar struct {
	SenderUserID                 string     		 `json:"senderUserId"`
	CreateAt            		 string      		 `json:"createAt"`
	OrderID             		 string     		 `json:"orderId"`
	OrderDate         		     string      		 `json:"orderDate"`
	WeightInGms                  float64    		 `json:"weightInGms"`
	Amount                       int         		 `json:"amount"`
	Symbol                       string    		     `json:"symbol"`
	PaymentStatus                string      		 `json:"paymentStatus"`
	KYCID                        string      		 `json:"kycId"`
	KYCVerified                  bool       		 `json:"kycVerified"`
	KYCLastVerified              string      		 `json:"kycLastVerfied"`
	PaymentProviderID            string      		 `json:"paymentProviderId"`
	PaymentCardID                string     		 `json:"paymentCardId"`
	PaymentCardDetails           []PaymentCardDetail `json:"paymentCardDetails"`
	AuthenticatedShopperID       string   			 `json:"authenticatedShopperId"`
}

type SendGoldBar struct {
	SenderUserID                 string     		 `json:"senderUserId"`
	RecieverUserID               string              `json:"recieverUserId"`
	CreateAt            		 string      		 `json:"createAt"`
	OrderID             		 string     		 `json:"orderId"`
	OrderDate         		     string      		 `json:"orderDate"`
	WeightInGms                  float64    		 `json:"weightInGms"`
	Amount                       int         		 `json:"amount"`
	PaymentStatus                string      		 `json:"paymentStatus"`
	KYCID                        string      		 `json:"kycId"`
	KYCVerified                  bool       		 `json:"kycVerified"`
	KYCLastVerified              string      		 `json:"kycLastVerfied"`
	KYCStatus                    string      		 `json:"kycStatus"`
	KYCProviderID                string              `json:"kycProviderID"`
	KYCEnabledFlag               bool                `json:"kycEnabledFlag"`
}

type AllGoldBarDetails struct {
	DocType             string             `json:"docType"`
	GoldBarArray        []GoldBarData      `json:"goldBarArray"`
}

type PurchaseConfirmation struct {
	GoldBarID  		        string     `json:"goldBarId"`
	GoldBarWeightAssigned   float64    `json:"goldBarWeightAssigned"`
	UserID                  string     `json:"userId"`
	OrderID                 string     `json:"orderID"`
}

type DWR struct {
	DWRID             string    `json:"dwrId"`
	DateIssued        string    `json:"dateIssued"`
	UserID            string    `json:"userID"`
	WeightInGms       float64   `json:"weightInGms"`
	Purity            float64    `json:"purity"`
	BarLocation       string    `json:"barLocation"`
	HalmarkVerified   bool      `json:"halmarkVerified"`
	GoldBarNumber     int       `json:"goldBarNumber"`
	GoldBarID         string    `json:"goldBarId"`
	OrderID           string    `json:"orderId"`
}

type UserWallet struct {
	UserID      			string      `json:"userID"`
	GoldBarID           	string      `json:"goldBarId"`
	GoldOwnsInGms           float64     `json:"goldOwnsInGms"`
	BarPurity               float64     `json:"barPurity"`
	BarLocation             string    `json:"barLocation"`
	BarSerialNumber         int       `json:"barSerialNumber"`
	BarHallmarkVerified     bool      `json:"barHallmarkVerfied"`
	BarRefiner              string    `json:"barRefiner"`
}

