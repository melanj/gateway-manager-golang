package models

type Gateway struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Serial      string `json:"serial"`
	Name        string `json:"name"`
	IPv4Address string `json:"ipv4Address"`
}
