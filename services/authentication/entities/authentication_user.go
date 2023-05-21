package entities

import (
	"time"
)

type DbAuthenticationUser struct {
	Id        string
	Email     string
	Username  string
	Password  string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
