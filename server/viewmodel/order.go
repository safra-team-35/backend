package viewmodel

// Order entity
type Order struct {
	Hash        string `json:"hash,omitempty"`
	PaymentType int64  `json:"payment_type,omitempty"`
	AddressID   int64  `json:"address_id,omitempty"`
}
