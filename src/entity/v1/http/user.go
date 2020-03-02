package http

import "time"

// UserRequest Scheme From Request
type UserRequest struct {
	Name           string  `json:"name" form:"name"`
	Email          string  `json:"email" form:"email"`
	Password       string  `json:"password" form:"password"`
	DateOfBirth    string  `json:"date_of_birth" form:"date_of_birth"`
	PhoneNumber    string  `json:"phone_number" form:"phone_number"`
	CurrentAddress string  `json:"current_address" form:"current_address"`
	City           string  `json:"city" form:"city"`
	Province       string  `json:"province" form:"province"`
	District       string  `json:"district" form:"district"`
	Lat            float64 `json:"lat" form:"lat"`
	Lng            float64 `json:"lng" form:"lng"`
	Job            string  `json:"job" form:"job"`
	Image          string  `json:"image" form:"image"`
	StartIncome    int     `json:"start_income" form:"start_income"`
	EndIncome      int     `json:"end_income" form:"end_income"`
	Interest       string  `json:"interests" form:"interests"`
	Habit          string  `json:"habit" form:"habit"`
	IsOrganizer    bool    `json:"is_organizer" form:"is_organizer"`
}

// UserDetailResponse scheme to detail respons
type UserDetailResponse struct {
	ID             uint       `json:"id"`
	Name           string     `json:"name"`
	Email          string     `json:"email"`
	Password       string     `json:"password"`
	DateOfBirth    string     `json:"date_of_birth"`
	PhoneNumber    string     `json:"phone_number"`
	CurrentAddress string     `json:"current_address"`
	City           string     `json:"city"`
	Province       string     `json:"province"`
	District       string     `json:"district"`
	Lat            float64    `json:"lat"`
	Lng            float64    `json:"lng"`
	Job            string     `json:"job"`
	Image          string     `json:"image"`
	StartIncome    int        `json:"start_income"`
	EndIncome      int        `json:"end_income"`
	Interest       string     `json:"interests"`
	Habit          string     `json:"habit"`
	IsOrganizer    bool       `json:"is_organizer"`
	CreatedAt      *time.Time `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
}

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
