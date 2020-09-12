package viewmodel

// QRCodeRequest - receive data to create a new QRCode
type QRCodeRequest struct {
	CompanyID      int64   `json:"company_id"`
	Price          float64 `json:"price"`
	ProductID      int64   `json:"product_id,omitempty"`
	ExpirationTime int64   `json:"expiration_time,omitempty" mapper:"-"`
}

// QRCodeResponse - response QRCode data
type QRCodeResponse struct {
	CompanyName string  `json:"company_name"`
	Price       float64 `json:"price"`
}
