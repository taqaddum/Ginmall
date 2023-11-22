package model

type Address struct {
	Preset
	UserID      int    `json:"user_id"`
	Username    string `json:"username"`
	Phone       string `json:"phone"`
	IsDefault   bool   `json:"is_default"`
	Province    string `json:"province"`
	City        string `json:"city"`
	Region      string `json:"region"`
	FullAddress string `json:"full_address"`
}
