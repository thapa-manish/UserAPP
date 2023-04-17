package model

type User struct {
	ID         int64  `json:"id"`
	Email      string `json:"email"`
	UserName   string `json:"user_name"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	UserStatus string `json:"user_status"`
	Department string `json:"department"`
}
