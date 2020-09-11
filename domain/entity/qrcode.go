package entity

// QRCode data information
type QRCode struct {
	CompanyID      int64   `json:"company_id"`
	Price          float64 `json:"price" validate:"required"`
	ExpirationTime string  `json:"expiration_time" validate:"required"`
	Hash           string  `json:"hash"`
}
