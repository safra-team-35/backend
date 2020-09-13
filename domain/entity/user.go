package entity

// Address entity
type Address struct {
	ID             int64  `json:"id,omitempty"`
	UserID         int64  `json:"order_id,omitempty"`
	Country        string `json:"country,omitempty"`
	Street         string `json:"street,omitempty"`
	Number         string `json:"number,omitempty"`
	Complement     string `json:"complement,omitempty"`
	ZipCode        string `json:"zip_code,omitempty"`
	City           string `json:"city,omitempty"`
	FederativeUnit string `json:"federative_unit,omitempty"`
}

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

// OrderSummary entity
type OrderSummary struct {
	CompanyName string  `json:"company_name,omitempty"`
	Price       float64 `json:"price,omitempty"`
	Freight     float64 `json:"freight,omitempty"`
}
