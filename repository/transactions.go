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

type Transaction = models.Transaction

func GetTransactions() ([]Transaction, error) {
	var transactions []Transaction

	query := `
	SELECT
		id,
		date,
		amount,
		type,
		description,
		category_id
		user_id
	FROM transactions
	ORDER BY id DESC
	`
	rows, err := database.DB.Query(query)
	if err != nil {
		return transactions, fmt.Errorf("Error obtaining transactions -> %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var t Transaction
		if err := rows.Scan(&t.ID, &t.Date, &t.Amount, &t.Type, &t.Description, &t.CategoryID, &t.UserID); err != nil {
			return transactions, fmt.Errorf("Error parsing users -> %w", err)
		}
		transactions = append(transactions, t)
	}

	return transactions, nil
}

func GetTransaction(id int64) (Transaction, error) {
	var transaction Transaction

	query := `
	SELECT
		id,
		date,
		amount,
		type,
		description,
		category_id,
		user_id
	FROM transactions
	WHERE id = ?
	ORDER BY id DESC
	`
	err := database.DB.QueryRow(query, id).
		Scan(&transaction.ID, &transaction.Date, &transaction.Amount, &transaction.Type, &transaction.Description, &transaction.CategoryID, &transaction.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return transaction, errors.New("transaction not found.")
		} else {
			return transaction, errors.New("Error obtaining transaction.")
		}
	}

	return transaction, nil
}

func NewTransaction(transaction *Transaction) error {

	query := `
	INSERT INTO transactions (date, amount, type, description, category_id, user_id)
	VALUES (?, ?, ?, ?, ?, ?)
	`
	res, err := database.DB.Exec(query, transaction.Date, transaction.Amount, transaction.Type, transaction.Description, transaction.CategoryID, transaction.UserID)
	if err != nil {
		return fmt.Errorf("Error creating transaction -> %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("Error getting generated transaction id-> %w", err)
	}

	transaction.ID = id

	return nil
}

func UpdateTransaction(transaction *Transaction) error {

	if transaction.ID == 0 {
		return errors.New("no ID was provided")
	}

	query := "UPDATE transactions SET "
	var args []any

	if transaction.Date != "" {
		query += "date = ?, "
		args = append(args, transaction.Date)
	}

	if transaction.Amount > 0 {
		query += "amount= ?, "
		args = append(args, transaction.Amount)
	}

	if transaction.Type != "" {
		query += "type = ?, "
		args = append(args, transaction.Type)
	}

	if transaction.CategoryID > 0 {
		query += "category_id = ?, "
		args = append(args, transaction.CategoryID)
	}

	if transaction.UserID > 0 {
		query += "user_id= ?, "
		args = append(args, transaction.UserID)
	}

	if len(args) == 0 {
		return nil
	}

	query = strings.TrimSuffix(query, ", ")

	query += " WHERE id = ?"
	args = append(args, transaction.ID)

	_, err := database.DB.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("Error updating transaction -> %w", err)
	}

	return nil
}

func DeleteTransaction(id int64) error {
	if id == 0 {
		return errors.New("no ID was provided")
	}

	query := `
	DELETE FROM transactions
	WHERE id = ?
	`
	_, err := database.DB.Exec(query, id)
	if err != nil {
		if utils.IsConstraintViolation(err) {
			return errors.New("Cannot delete transaction -> transaction is in use by one or more transactions")
		}
		return fmt.Errorf("Error deleting transaction -> %w", err)
	}

	return nil
}
