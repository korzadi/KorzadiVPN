package models

type User struct {
	Email    string `json:"email"`
	Password string `json:"-"`
	Plan     string `json:"plan"`
	Status   string `json:"status"`
}

func GetDeviceLimit(plan string) int {

	switch plan {

	case "premium":
		return 5

	case "enterprise":
		return 20

	default:
		return 1
	}
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
