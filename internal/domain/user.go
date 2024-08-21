package domain

import "time"

type User struct {
	ID               string    `json:"id"`
	Name             string    `json:"name"`
	Email            string    `json:"email"`
	RegistrationDate time.Time `db:"registration_date"`
	Role             string    `json:"role"`
}
