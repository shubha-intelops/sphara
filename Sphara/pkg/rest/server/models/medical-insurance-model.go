package models

import "gorm.io/gorm"

type MedicalInsurance struct {
	gorm.Model
	Id int64 `gorm:"primaryKey;autoIncrement" json:"ID,omitempty"`

	File string `json:"file,omitempty"`
}
