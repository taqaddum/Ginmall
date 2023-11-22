package model

type Order struct {
	ID            int    `json:"id"`
	Username      string `json:"username"`
	Phone         string `json:"phone"`
	Province      string `json:"province"`
	City          string `json:"city"`
	Region        string `json:"region"`
	DetailAddress string `json:"detail_address"`
}
