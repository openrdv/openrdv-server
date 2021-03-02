package models

import "github.com/jinzhu/gorm"

type Device struct {
	gorm.Model
	UID     string
	UIDType string
	Token   string `json:"-"`
}
