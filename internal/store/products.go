package store

import (
	"database/sql"
)

type Product struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
	CreatedAt   string `json:"created_at"`
	ExpiredDate string `json:"expired_date,omitempty"`
}

type ProductStore struct {
	db *sql.DB
}
