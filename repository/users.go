package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/eusebioleite/selfin/database"
	"github.com/eusebioleite/selfin/models"
	"github.com/eusebioleite/selfin/utils"
)

type User = models.User

func GetUsers() ([]User, error) {
	var users []User

	query := `
		SELECT
			id,
			name,
			login,
			password
		FROM users
		ORDER BY id DESC
		`
	rows, err := database.DB.Query(query)
	if err != nil {
		return users, fmt.Errorf("Error obtaining users -> %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Login, &u.Password); err != nil {
			return users, fmt.Errorf("Error parsing users -> %w", err)
		}
		users = append(users, u)
	}

	return users, nil
}

func GetUser(id int64) (User, error) {
	var user User

	query := `
		SELECT
			id,
			name,
			login,
			password
		FROM users
		WHERE id = ?
		ORDER BY id DESC
		`
	err := database.DB.QueryRow(query, id).
		Scan(&user.ID, &user.Name, &user.Login, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("User not found.")
		} else {
			return user, errors.New("Error obtaining user.")
		}
	}

	return user, nil
}

func GetUserByLogin(login string) (User, error) {
	var user User

	query := `
		SELECT
			id,
			name,
			login,
			password
		FROM users
		WHERE login = ?
		ORDER BY id DESC
		`
	err := database.DB.QueryRow(query, login).
		Scan(&user.ID, &user.Name, &user.Login, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("User not found.")
		} else {
			return user, errors.New("Error obtaining user.")
		}
	}

	return user, nil
}

func NewUser(user *User) error {

	query := `
	INSERT INTO users (name, login, password)
	VALUES (?, ?, ?)
	`
	res, err := database.DB.Exec(query, user.Name, user.Login, user.Password)
	if err != nil {
		return fmt.Errorf("Error creating new user -> %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("Error obtaining 'id' generated for user -> %w", err)
	}

	user.ID = id

	return nil
}

func UpdateUser(user *User) error {

	if user.ID == 0 {
		return errors.New("no ID was provided")
	}

	query := "UPDATE users SET "
	var args []any

	if user.Name != "" {
		query += "name = ?, "
		args = append(args, user.Name)
	}

	if user.Login != "" {
		query += "login = ?, "
		args = append(args, user.Login)
	}

	if user.Password != "" {
		query += "password = ?, "
		args = append(args, user.Password)
	}

	if len(args) == 0 {
		return nil
	}

	query = strings.TrimSuffix(query, ", ")

	query += " WHERE id = ?"
	args = append(args, user.ID)

	_, err := database.DB.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("error updating user -> %w", err)
	}

	return nil
}

func DeleteUser(id int64) error {
	if id == 0 {
		return errors.New("no ID was provided")
	}

	query := `
	DELETE FROM users
	WHERE id = ?
	`
	_, err := database.DB.Exec(query, id)
	if err != nil {
		if utils.IsConstraintViolation(err) {
			return errors.New("Cannot delete user -> user is in use by one or more transactions")
		}
		return fmt.Errorf("Error deleting user -> %w", err)
	}

	return nil
}
