package repositories

import (
	"github.com/BrandonWade/enako/api/models"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

type ExpenseRepository interface {
	GetExpenses() ([]models.Expense, error)
	CreateExpense(userID int64, expense *models.Expense) (int64, error)
	UpdateExpense(ID, userID int64, expense *models.Expense) (int64, error)
	DeleteExpense(ID, userID int64) (int64, error)
}

type expenseRepository struct {
	DB *sqlx.DB
}

func NewExpenseRepository(DB *sqlx.DB) ExpenseRepository {
	return &expenseRepository{
		DB,
	}
}

func (e *expenseRepository) GetExpenses() ([]models.Expense, error) {
	userID := 1 // TODO: Hardcoded for testing
	expenses := []models.Expense{}

	err := e.DB.Select(&expenses, `SELECT *
        FROM expenses AS e
        WHERE e.user_id = ?;
    `, userID)
	if err != nil {
		return []models.Expense{}, err
	}

	return expenses, nil
}

func (e *expenseRepository) CreateExpense(userID int64, expense *models.Expense) (int64, error) {
	result, err := e.DB.Exec(`INSERT
		INTO expenses(
			user_id,
			type,
			category,
			description,
			amount,
			date
		) VALUE (
			?,
			?,
			?,
			?,
			?,
			?
		);
	`,
		userID,
		expense.Type,
		expense.Category,
		expense.Description,
		expense.Amount,
		expense.Date,
	)
	if err != nil {
		return 0, err
	}

	ID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return ID, nil
}

func (e *expenseRepository) UpdateExpense(ID, userID int64, expense *models.Expense) (int64, error) {
	result, err := e.DB.Exec(`UPDATE expenses
		SET
	  		user_id = ?,
			type = ?,
			category = ?,
			description = ?,
			amount = ?,
			date = ?
		WHERE id = ?;
	`,
		userID,
		expense.Type,
		expense.Category,
		expense.Description,
		expense.Amount,
		expense.Date,
		ID,
	)
	if err != nil {
		return 0, err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (e *expenseRepository) DeleteExpense(ID, userID int64) (int64, error) {
	result, err := e.DB.Exec(`DELETE
		FROM expenses
		WHERE id = ?
		AND user_id = ?;
	`,
		ID,
		userID,
	)
	if err != nil {
		return 0, err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return count, nil
}
