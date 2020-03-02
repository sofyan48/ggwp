package http

import "time"

// UserResponse scheme to all respons
type UserResponse struct {
	ID             uint       `json:"id"`
	Name           string     `json:"name"`
	Email          string     `json:"email"`
	Password       string     `json:"password"`
	DateOfBirth    string     `json:"date_of_birt"`
	PhoneNumber    string     `json:"phone_number"`
	CurrentAddress string     `json:"current_address"`
	City           string     `json:"city"`
	Province       string     `json:"province"`
	District       string     `json:"district"`
	CreatedAt      *time.Time `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
}
