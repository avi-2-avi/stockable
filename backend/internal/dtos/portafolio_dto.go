package dtos

import (
	"github.com/google/uuid"
)

type PortafolioDTO struct {
	ID         uuid.UUID     `json:"id"`
	Name       string        `json:"name"`
	Category   string        `json:"category"`
	DataSource DataSourceDTO `json:"data_source"`
}
