package models

import "github.com/jinzhu/gorm"

type Attest struct {
	gorm.Model
	DeviceID uint
	Device Device
}
