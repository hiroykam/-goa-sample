package entities

import "time"

type User struct {
	ID             int `gorm:"primary_key"`
	Email          string
	HashedPassword string
	RevokedAt      *time.Time
	LockedAt       *time.Time
	DeletedAt      *time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
