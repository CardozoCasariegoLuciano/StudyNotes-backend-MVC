package models

import "time"

type CommonModelFields struct {
	ID        uint       `gorm:"primary_key" mapper:"_id" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
