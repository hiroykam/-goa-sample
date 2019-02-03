package entities

import "time"

type Sample struct {
	ID        int `gorm:"primary_key"`
	Detail    string
	Name      string
	UserID    int
	CreatedAt time.Time
	DeletedAt *time.Time
	UpdatedAt time.Time
}
