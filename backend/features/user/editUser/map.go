package editUser

import (
	. "elytra.com/backend/features/user/database"
	pgx_google_uuid "github.com/vgarvardt/pgx-google-uuid/v5"
)

func mapToParams(r EditUserRequest) EditUserParams {
	return EditUserParams{
		ID:       pgx_google_uuid.UUID(r.ID),
		Username: r.Username,
		Email:    r.Email,
	}
}
