package viewmodel

// QRCode - receive data to create a new QRCode
type QRCode struct {
	Price          float64 `json:"price"`
	ExpirationTime int64   `json:"expiration_time" mapper:"-"`
}

// PrivateUser - return a struct with all data. It's need a token to request
type PrivateUser struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}
