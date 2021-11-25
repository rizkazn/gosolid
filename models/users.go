package models

import "time"

type User struct {
	Id        int       `json:"id,omitempty"`
	Name      string    `json:"name"`
	UserName  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}

type Users []User

func CreateUsers() *User {
	return &User{
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}
}
