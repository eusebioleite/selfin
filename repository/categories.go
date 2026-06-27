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

type Category = models.Category

func GetCategories() ([]Category, error) {
	var categories []Category

	query := `
	SELECT
		id,
		description
	FROM categories
	ORDER BY id DESC
	`
	rows, err := database.DB.Query(query)
	if err != nil {
		return categories, fmt.Errorf("Error obtaining categories -> %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var c Category
		if err := rows.Scan(&c.ID, &c.Description); err != nil {
			return categories, fmt.Errorf("Error obtaining categories -> %w", err)
		}
		categories = append(categories, c)
	}

	return categories, nil
}

func GetCategory(id int64) (Category, error) {
	var category Category

	query := `
	SELECT
		id,
		description
	FROM categories
	WHERE id = ?
	ORDER BY id DESC
	`
	err := database.DB.QueryRow(query, id).
		Scan(&category.ID, &category.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return category, errors.New("category not found.")
		} else {
			return category, errors.New("Error obtaining category.")
		}
	}

	return category, nil
}

func NewCategory(category *Category) error {

	query := `
	INSERT INTO categories (description)
	VALUES (?)
	`
	res, err := database.DB.Exec(query, category.Description)
	if err != nil {
		return fmt.Errorf("Error creating new category -> %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("Error obtaining 'id' generated for category -> %w", err)
	}

	category.ID = id

	return nil
}

func UpdateCategory(category *Category) error {

	if category.ID == 0 {
		return errors.New("no ID was provided")
	}

	query := "UPDATE categories SET "
	var args []any

	if category.Description != "" {
		query += "description = ?, "
		args = append(args, category.Description)
	}

	if len(args) == 0 {
		return nil
	}

	query = strings.TrimSuffix(query, ", ")

	query += " WHERE id = ?"
	args = append(args, category.ID)

	_, err := database.DB.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("Error updating category -> %w", err)
	}

	return nil
}

func DeleteCategory(id int64) error {
	if id == 0 {
		return errors.New("no ID was provided")
	}

	_, err := database.DB.Exec(
		"DELETE FROM categories where id = ?", id)
	if err != nil {
		if utils.IsConstraintViolation(err) {
			return errors.New("category is in use by one or more transactions")
		}
		return fmt.Errorf("Error deleting category -> %w", err)
	}

	return nil
}
