package dto

import "time"

type Model struct {
	ID        uint      `json:"id" form:"id"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" form:"deleted_at"`
}
