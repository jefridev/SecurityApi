package roles

// RoleRegistration holds a data for saving.
type RoleRegistration struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

// UserRoleRegistration holds data for saving a new relationship between user and role or remove it.
type UserRoleRegistration struct {
	UserID uint64 `json:"userId"`
	RoleID uint64 `json:"roleId"`
	Status bool   `json:"status"`
}
