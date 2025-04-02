package dtos

import "github.com/google/uuid"

type DataSourceDTO struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	IsVisible bool      `json:"is_visible"`
}
