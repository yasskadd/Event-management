package models

import "time"

type Event struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Date        time.Time `json:"date"`
	OrganizerID uint      `json:"organizer_id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
