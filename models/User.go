package models

import "time"

type User struct {
	ID        uint `json:"id" gorm:"primayKey"`
	CreatedAt time.Time `json:"created_at"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}