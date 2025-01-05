package models

type User struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Salary     int64  `json:"salary"`
	Department string `json:"department"`
	AddressId  string `json:"address_id"`
}

type Address struct {
	ID    string `json:"id"`
	City  string `json:"city"`
	State string `json:"state"`
}