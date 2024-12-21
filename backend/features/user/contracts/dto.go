package contracts

import (
	"elytra.com/backend/features/user/database"
	"github.com/google/uuid"
)

type UserDTO struct {
	Id    uuid.UUID `json:"id" example:"ddf6aa7a-625b-4c29-93f0-faf617af5a8e"`
	Name  string    `json:"name" example:"test"`
	Email string    `json:"email" example:"test@gmail.com"`
}

func MapToDto(u database.User) UserDTO {
	return UserDTO{
		Id:    uuid.UUID(u.ID),
		Name:  u.Username,
		Email: u.Email,
	}
}
