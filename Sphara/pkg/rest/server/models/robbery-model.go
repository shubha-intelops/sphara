package models

import "gorm.io/gorm"

type Robbery struct {
	gorm.Model
	Id int64 `gorm:"primaryKey;autoIncrement" json:"ID,omitempty"`

	Address string `json:"address,omitempty"`

	File string `json:"file,omitempty"`

	Injured bool `json:"injured,omitempty"`

	Place string `json:"place,omitempty"`

	Robbed_things string `json:"robbedThings,omitempty"`

	Value string `json:"value,omitempty"`
}
