package user

type UsersDto struct {
	Name string
}

// UpdateUserRequest : request to update a user
type UpdateUserRequest struct {
	Name *string
	Role *RoleDto
}

type RoleDto struct {
	Name string
}
