package entity

// Order entity
type Order struct {
	ID          int64   `json:"id,omitempty"`
	CompanyID   int64   `json:"company_id,omitempty"`
	UserID      int64   `json:"user_id,omitempty"`
	ProductID   int64   `json:"product_id,omitempty"`
	Price       float64 `json:"price,omitempty"`
	Freight     float64 `json:"freight,omitempty"`
	Hash        string  `json:"hash,omitempty"`
	QRCodeID    int64   `json:"qr_code_id,omitempty"`
	PaymentType int64   `json:"payment_type,omitempty"`
	AddressID   int64   `json:"address_id,omitempty"`
}

type OrderPartnerData struct {
	Price float64 `json:"price,omitempty"`
}
