package models

import "gorm.io/gorm"

type MedicalCondition struct {
	gorm.Model
	Id int64 `gorm:"primaryKey;autoIncrement" json:"ID,omitempty"`

	File string `json:"file,omitempty"`

	Medical_condition_details string `json:"medicalConditionDetails,omitempty"`
}
