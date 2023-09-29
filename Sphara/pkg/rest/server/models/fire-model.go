package models

import "gorm.io/gorm"

type Fire struct {
	gorm.Model
	Id int64 `gorm:"primaryKey;autoIncrement" json:"ID,omitempty"`

	Comment string `json:"comment,omitempty"`

	File string `json:"file,omitempty"`

	Fire_dimension string `json:"fireDimension,omitempty"`

	Flame_smoke string `json:"flameSmoke,omitempty"`

	People_injured_status string `json:"peopleInjuredStatus,omitempty"`
}
