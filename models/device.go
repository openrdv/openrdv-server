package models

import (
	_ "github.com/jinzhu/gorm"
)

type Device struct {
	ID uint `json:"id" gorm:"primary_key"`
}
