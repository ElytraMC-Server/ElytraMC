package createUser

import (
	"github.com/google/uuid"

	. "elytra.com/backend/features/user/contracts"
	. "elytra.com/backend/features/user/createUser/database"
	pgx_google_uuid "github.com/vgarvardt/pgx-google-uuid/v5"
)

func mapToDto(u User) UserDTO {
	return UserDTO{
		Id:    uuid.UUID(u.ID),
		Name:  u.Username,
		Email: u.Email,
	}
}

func mapToParams(r CreateUserRequest) CreateUserParams {
	return CreateUserParams{
		ID:       pgx_google_uuid.UUID(uuid.New()),
		Username: r.Username,
		Email:    r.Email,
	}
}
