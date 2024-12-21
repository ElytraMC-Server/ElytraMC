package editUser

import (
	"github.com/google/uuid"

	. "elytra.com/backend/features/user/contracts"
	. "elytra.com/backend/features/user/database"
	pgx_google_uuid "github.com/vgarvardt/pgx-google-uuid/v5"
)

func mapToDto(u User) UserDTO {
	return UserDTO{
		Id:    uuid.UUID(u.ID),
		Name:  u.Username,
		Email: u.Email,
	}
}

func mapToParams(r EditUserRequest) EditUserParams {
	return EditUserParams{
		ID:       pgx_google_uuid.UUID(r.ID),
		Username: r.Username,
		Email:    r.Email,
	}
}
