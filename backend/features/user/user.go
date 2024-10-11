package user

import . "github.com/google/uuid"

type UserDTO struct {
	Id    UUID   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
