package database

import (
	"korzadivpn/models"
)

func CreateUser(user models.User) error {

	_, err := DB.Exec(
		`
		INSERT INTO users
		(email,password,plan,status)
		VALUES (?,?,?,?)
		`,
		user.Email,
		user.Password,
		user.Plan,
		user.Status,
	)

	return err
}

func GetUser(email string) (models.User, error) {

	var user models.User

	err := DB.QueryRow(
		`
		SELECT email,password,plan,status
		FROM users
		WHERE email=?
		`,
		email,
	).Scan(
		&user.Email,
		&user.Password,
		&user.Plan,
		&user.Status,
	)

	return user, err
}

func UpdateUserPlan(
	email string,
	plan string,
) error {

	_, err := DB.Exec(
		`
		UPDATE users
		SET plan=?
		WHERE email=?
		`,
		plan,
		email,
	)

	return err
}
