package models

import "gorm.io/gorm"

type Ambulance struct {
	gorm.Model
	Id int64 `gorm:"primaryKey;autoIncrement" json:"ID,omitempty"`

	File string `json:"file,omitempty"`

	Injuredcount int `json:"injuredcount,omitempty"`

	Reason string `json:"reason,omitempty"`
}
