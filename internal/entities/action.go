package entities

import "time"

type Actions struct {
	ID            int
	UserID        int
	User          User
	ActionTime    time.Time
	RequestResult int
	TempC         float64 // Use pointer for nullable fields
}
