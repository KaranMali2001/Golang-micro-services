package main

import "time"

type Account struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Orders    []Order
	CreatedAt *time.Time `json:"created_at"` // Uppercase first letter, use JSON tag
	UpdatedAt *time.Time `json:"updated_at"` // Uppercase first letter, use JSON tag
}
