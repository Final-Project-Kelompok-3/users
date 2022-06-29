package dto

type CreateUserRequest struct {
	RoleID    int    `json:"roleId" validate:"required"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"nullable"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	Verified  bool   `json:"verified" validate:"required"`
}

type UpdateUserRequest struct {
	Name     *string `json:"name"`
	Email    *string `json:"email" validate:"email"`
	Password *string `json:"password"`
}

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}