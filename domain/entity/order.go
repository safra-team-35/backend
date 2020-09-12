package entity

// Order entity
type Order struct {
	ID          int64   `json:"id,omitempty"`
	CompanyID   int64   `json:"company_id,omitempty"`
	Price       float64 `json:"price,omitempty"`
	Hash        string  `json:"hash,omitempty"`
	PaymentType int64   `json:"payment_type,omitempty"`
	AddressID   int64   `json:"address_id,omitempty"`
}
