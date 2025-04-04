package dtos

import "github.com/google/uuid"

type CompanyDTO struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Ticker string    `json:"ticker"`
}
