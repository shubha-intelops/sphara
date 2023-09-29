package models

import "gorm.io/gorm"

type Panic struct {
	gorm.Model
	Id int64 `gorm:"primaryKey;autoIncrement" json:"ID,omitempty"`

	Description string `json:"description,omitempty"`

	File string `json:"file,omitempty"`
}
