package storage

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

func OpenDB(connStr string) (*DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return &DB{}, fmt.Errorf("db connect error: %v", err)
	}

	if err = db.Ping(); err != nil {
		return &DB{}, fmt.Errorf("db ping error: %v", err)
	}

	return &DB{db}, nil
}

func (db *DB) CreateUserTable() (*UserModel, error) {
	query := `CREATE TABLE IF NOT EXISTS users (
    	id SERIAL PRIMARY KEY,
    	user_name TEXT UNIQUE NOT NULL,
    	first_name TEXT NOT NULL,
    	last_name TEXT NOT NULL,
    	email TEXT UNIQUE NOT NULL,
    	password TEXT NOT NULL,
    	role TEXT NOT NULL,
    	register_at TIMESTAMP NOT NULL,
	)`
	_, err := db.Exec(query)
	if err != nil {
		switch err.(*pq.Error).Code {
		case "42P07":
			return &UserModel{db, "users"}, nil
		default:
			return &UserModel{}, fmt.Errorf("create user table error: %v", err)
		}
	}

	return &UserModel{db, "users"}, nil
}
