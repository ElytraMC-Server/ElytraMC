package loginUser

import (
	. "elytra.com/backend/features/user/database"
)

func mapToParams(r LoginUserRequest) LoginUserParams {
	return LoginUserParams{
		Username: r.Username,
		Email:    r.Email,
	}
}
