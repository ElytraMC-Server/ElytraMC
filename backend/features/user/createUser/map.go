package createUser

import (
	"github.com/google/uuid"

	. "elytra.com/backend/features/user/database"
	pgx_google_uuid "github.com/vgarvardt/pgx-google-uuid/v5"
)

func mapToParams(r CreateUserRequest) CreateUserParams {
	return CreateUserParams{
		ID:       pgx_google_uuid.UUID(uuid.New()),
		Username: r.Username,
		Email:    r.Email,
	}
}
