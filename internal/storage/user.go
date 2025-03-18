package storage

import (
	"fmt"
	"time"
)

type UserModel struct {
	*DB
	table string
}

type User struct {
	UserName  string `json:"user_name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (db UserModel) CreateUser(user User) error {
	query := fmt.Sprintf(`INSERT INTO %s (user_name, first_name, last_name, created) VALUES ($1, $2, $3, $4)`, db.table)
	_, err := db.Exec(query, user.UserName, user.FirstName, user.LastName, time.Now())
	if err != nil {
		return fmt.Errorf("create user error: %v", err)
	}

	return nil
}

func (db UserModel) GetUserById(id int) (User, error) {
	query := fmt.Sprintf(`SELECT user_name, first_name, last_name FROM %s WHERE id = $1`, db.table)
	row := db.QueryRow(query, id)
	user := User{}
	err := row.Scan(&user.UserName, &user.FirstName, &user.LastName)
	if err != nil {
		return User{}, fmt.Errorf("get user by id error: %v", err)
	}

	return user, nil
}
