package models

import "github.com/jinzhu/gorm"
import "gorm.io/datatypes"

type AttestModuleResult map[string]string

type Attest struct {
	gorm.Model
	DeviceID uint
	Device Device
	Result datatypes.JSON
}
