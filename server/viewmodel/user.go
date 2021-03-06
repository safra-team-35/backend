package viewmodel

// Address model
type Address struct {
	ID             int64  `json:"id,omitempty"`
	Country        string `json:"country,omitempty"`
	Street         string `json:"street,omitempty"`
	Number         string `json:"number,omitempty"`
	Complement     string `json:"complement,omitempty"`
	ZipCode        string `json:"zip_code,omitempty"`
	City           string `json:"city,omitempty"`
	FederativeUnit string `json:"federative_unit,omitempty"`
}

// Order model
type Order struct {
	Hash        string  `json:"hash,omitempty"`
	PaymentType int64   `json:"payment_type,omitempty"`
	AddressID   int64   `json:"address_id,omitempty"`
	Freight     float64 `json:"freight,omitempty"`
}

// OrderSummary model
type OrderSummary struct {
	CompanyName string  `json:"company_name,omitempty"`
	Price       float64 `json:"price,omitempty"`
	Freight     float64 `json:"freight,omitempty"`
}
