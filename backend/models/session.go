package models

type Session struct {
	ID int `json:"id"`

	Email string `json:"email"`

	Token string `json:"-"`

	IP string `json:"ip"`

	Device string `json:"device"`

	CreatedAt string `json:"created_at"`

	ExpiresAt string `json:"expires_at"`

	Status string `json:"status"`
}
