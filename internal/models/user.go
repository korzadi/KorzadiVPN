package models

type User struct {
	Email    string `json:"email"`
	Password string `json:"-"`
	Plan     string `json:"plan"`
	Status   string `json:"status"`
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
