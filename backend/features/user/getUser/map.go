package getUser

import (
	"github.com/google/uuid"

	. "elytra.com/backend/features/user/contracts"
	. "elytra.com/backend/features/user/getUser/database"
)

func mapToDto(u User) UserDTO {
	return UserDTO{
		Id:    uuid.UUID(u.ID),
		Name:  u.Username,
		Email: u.Email,
	}
}
