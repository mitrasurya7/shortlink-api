package entities

import "time"

type Shortlink_tab struct {
	ID           uint   `gorm:"primaryKey"`
	Shortlink    string `gorm:"unique,not null"`
	Redirectlink string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
