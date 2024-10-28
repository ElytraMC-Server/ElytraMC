package contracts

import "github.com/google/uuid"

type UserDTO struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}
