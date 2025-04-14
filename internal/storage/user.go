package storage

import (
	"fmt"
	"oop/internal/models"
	"time"
)

type UserModel struct {
	*DB
	table string
}

func (db *UserModel) CreateUser(user models.User) error {
	query := fmt.Sprintf(`INSERT INTO %s (user_name, first_name, last_name, email, password, role, created) VALUES ($1, $2, $3, $4, $5, $6, $7)`, db.table)
	// TODO: hash password
	_, err := db.Exec(query, user.UserName, user.FirstName, user.LastName, user.Email, user.Password, user.Role, time.Now())
	if err != nil {
		return fmt.Errorf("create user error: %v", err)
	}

	return nil
}

func (db *UserModel) GetUserById(id int) (models.User, error) {
	query := fmt.Sprintf(`SELECT user_name, first_name, last_name, email, role, register_at FROM %s WHERE id = $1`, db.table)
	row := db.QueryRow(query, id)
	user := models.User{}
	err := row.Scan(&user.UserName, &user.FirstName, &user.LastName, &user.Email, &user.Role, &user.RegisterAt)
	if err != nil {
		return models.User{}, fmt.Errorf("get user by id error: %v", err)
	}

	return user, nil
}
