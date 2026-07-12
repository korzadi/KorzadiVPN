package models

type Activity struct {
	ID int `json:"id"`

	Email string `json:"email"`

	Server string `json:"server"`

	Action string `json:"action"`

	Device string `json:"device"`

	IP string `json:"ip"`

	CreatedAt string `json:"created_at"`
}
