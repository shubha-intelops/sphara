package models

import "gorm.io/gorm"

type Signup struct {
	gorm.Model
	Id int64 `gorm:"primaryKey;autoIncrement" json:"ID,omitempty"`

	City string `json:"city,omitempty"`

	Country string `json:"country,omitempty"`

	Dob string `json:"dob,omitempty"`

	EmailId string `json:"emailId,omitempty"`

	Fullname string `json:"fullname,omitempty"`

	Nickname string `json:"nickname,omitempty"`

	Others string `json:"others,omitempty"`

	R_address string `json:"rAddress,omitempty"`

	Zip int `json:"zip,omitempty"`
}
