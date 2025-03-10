package dtos

import "github.com/google/uuid"

type LoginUserDTO struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	FullName string    `json:"full_name"`
}
