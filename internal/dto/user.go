package dto

type CreateUserRequest struct {
	RoleID    int    `json:"role_id" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
}

type UpdateUserRequest struct {
	RoleID    *int    `json:"role_id"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Email     *string `json:"email" validate:"email"`
	Password  *string `json:"password"`
}

type UserResponse struct {
	ID        int    `json:"id"`
	RoleID    int    `json:"role_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" validate:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"deleted_at"`
}