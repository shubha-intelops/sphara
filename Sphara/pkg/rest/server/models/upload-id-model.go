package models

import "gorm.io/gorm"

type UploadId struct {
	gorm.Model
	Id int64 `gorm:"primaryKey;autoIncrement" json:"ID,omitempty"`

	File string `json:"file,omitempty"`

	Id_number string `json:"idNumber,omitempty"`

	Id_type string `json:"idType,omitempty"`
}
