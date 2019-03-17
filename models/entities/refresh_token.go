package entities

import "time"

type RefreshToken struct {
	Jti       string `gorm:"primary_key"`
	UserID    int
	CreatedAt time.Time
	DeletedAt *time.Time
	UpdatedAt time.Time
}
