package models

type Attest struct {
	ID uint `json:"id" gorm:"primary_key"`
	Device Device
}
