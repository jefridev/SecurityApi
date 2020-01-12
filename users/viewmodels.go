package users

// UserRegistration holds data for saving a new user within database.
type UserRegistration struct {
	ID                   uint64 `json:"-"`
	Username             string `json:"username"`
	Email                string `json:"email"`
	Bio                  string `json:"bio"`
	Image                string `json:"image"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
}
