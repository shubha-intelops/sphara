package models

import "gorm.io/gorm"

type EmgContact struct {
	gorm.Model
	Id int64 `gorm:"primaryKey;autoIncrement" json:"ID,omitempty"`

	Contact_number string `json:"contactNumber,omitempty"`

	Name string `json:"name,omitempty"`

	Tag string `json:"tag,omitempty"`
}
