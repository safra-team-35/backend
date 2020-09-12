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
