package models

import "time"

type Event struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Date        time.Time `json:"date"`
	Status      string    `json:"status"`       // "online", "offline", "waitlist"
	TicketCount int       `json:"ticket_count"` // Total tickets available
	TicketsSold int       `json:"tickets_sold"` // Tickets sold
	OrganizerID uint      `json:"organizer_id"` // FK to User
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Subscription struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id"`  // FK to User
	EventID   uint      `json:"event_id"` // FK to Event
	Status    string    `json:"status"`   // "attending", "waitlist"
	CreatedAt time.Time `json:"created_at"`
}
