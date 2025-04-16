package internal

type UserInput struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	PasswordHashed string `json:"password_hashed"`
}
