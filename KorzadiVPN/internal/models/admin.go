package models

type Admin struct {
	ID int `json:"id"`

	Email string `json:"email"`

	Role string `json:"role"`

	CreatedAt string `json:"created_at"`
}
