package models

import "time"

type Product struct {
	Id        int       `json:"id,omitempty"`
	Name      string    `json:"name"`
	Price     string    `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}

type Products []Product

func CreateProds() *Product {
	return &Product{
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}
}
