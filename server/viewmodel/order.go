// Order entity
type Order struct {
	Hash        string       `json:"hash,omitempty"`
	PaymentType int64        `json:"payment_type,omitempty"`
	Address     OrderAddress `json:"address,omitempty"`
}

// OrderAddress entity
type OrderAddress struct {
	Country        string `json:"country,omitempty"`
	Street         string `json:"street,omitempty"`
	Number         string `json:"number,omitempty"`
	Complement     string `json:"complement,omitempty"`
	ZipCode        string `json:"zip_code,omitempty"`
	City           string `json:"city,omitempty"`
	FederativeUnit string `json:"federative_unit,omitempty"`
}