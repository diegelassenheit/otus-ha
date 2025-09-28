package domain

import "time"

type User struct {
    ID           string
	FirstName    string
	SecondName   string
	Birthdate    time.Time
	Biography    string
	City         string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
}
