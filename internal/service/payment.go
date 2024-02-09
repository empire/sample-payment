package service

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type PaymentService struct {
	db *sql.DB
}

func NewPaymentService(dbPath string) (*PaymentService, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	createTableStmt := `
        CREATE TABLE IF NOT EXISTS payments (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            account_id INTEGER NOT NULL,
            amount DECIMAL(10,2),
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP
        )
    `
	if _, err := db.Exec(createTableStmt); err != nil {
		return nil, err
	}

	return &PaymentService{db: db}, nil
}

func (ps *PaymentService) Deposit(accountID int, amount float64) error {
	stmt, err := ps.db.Prepare("INSERT INTO payments (account_id, amount) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(accountID, amount)
	return err
}

func (ps *PaymentService) Balance(accountID int) (float64, error) {
	var totalBalance float64
	row := ps.db.QueryRow("SELECT SUM(amount) FROM payments WHERE account_id = ?", accountID)
	err := row.Scan(&totalBalance)
	return totalBalance, err
}

func (ps *PaymentService) Withdraw(accountID int, amount float64) error {
	balance, err := ps.Balance(accountID)
	if err != nil {
		return err
	}
	if balance < amount {
		return fmt.Errorf("insufficient funds: Balance is %.2f, cannot withdraw %.2f", balance, amount)
	}

	stmt, err := ps.db.Prepare("INSERT INTO payments (account_id, amount) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(accountID, -amount)
	return err
}
