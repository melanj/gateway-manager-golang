package models

import "time"

type Device struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	UID         uint      `json:"uid"`
	Vendor      string    `json:"vendor"`
	DateCreated time.Time `json:"dateCreated"`
	Status      string    `json:"status"`
	GatewayId     uint   `json:"gateway_id" gorm:"foreignKey:gateway_id"`
}
