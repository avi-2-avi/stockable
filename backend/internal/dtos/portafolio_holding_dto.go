package dtos

import "github.com/google/uuid"

type PortafolioHoldingDTO struct {
	ID            uuid.UUID  `json:"id"`
	PortafolioID  uuid.UUID  `json:"portafolio_id"`
	Company       CompanyDTO `json:"company"`
	Quantity      float64    `json:"quantity"`
	PurchasePrice float64    `json:"purchase_price"`
	PurchaseAt    string     `json:"purchase_at"`
}
