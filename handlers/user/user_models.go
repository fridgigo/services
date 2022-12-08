package user

import "time"

type User struct {
	Email          string     `json:"email"`
	FirstName      string     `json:"first_name"`
	LastName       string     `json:"last_name"`
	Password       string     `json:"password"`
	RepeatPassword string     `json:"repeat_password"`
	CreatedAt      *time.Time `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
	Deleted        bool       `json:"deleted"`
	Confirmed      bool       `json:"confirmed"`
}
