package entity

// QRCode data information
type QRCode struct {
	ID             int64   `json:"id"`
	CompanyID      int64   `json:"company_id"`
	CompanyName    string  `json:"company_name"`
	ProductID      int64   `json:"product_id"`
	Price          float64 `json:"price" validate:"required"`
	ExpirationTime string  `json:"expiration_time" validate:"required"`
	Hash           string  `json:"hash"`
}
