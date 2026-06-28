package models

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Category struct {
	ID          int64  `json:"id"`
	Description string `json:"description"`
}

type Transaction struct {
	ID          int64   `json:"id"`
	Date        string  `json:"date"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	Type        string  `json:"type"`
	CategoryID  int64   `json:"category_id"`
	UserID      int64   `json:"user_id"`
}

type TransactionView struct {
	ID                  int64   `json:"id"`
	Date                string  `json:"date"`
	Amount              float64 `json:"amount"`
	Description         string  `json:"description"`
	Type                string  `json:"type"`
	CategoryID          int64   `json:"category_id"`
	CategoryDescription string  `json:"category_description"`
	UserID              int64   `json:"user_id"`
	UserName            string  `json:"user_name"`
}

type Session struct {
	ID        string `json:"id"`
	UserID    int64  `json:"user_id"`
	ExpiresAt string `json:"expires_at"`
	CreatedAt string `json:"created_at"`
}

type UserSession struct {
	UserID    int64  `json:"user_id"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	SessionID string `json:"session_id"`
	ExpiresAt string `json:"expires_at"`
	CreatedAt string `json:"created_at"`
}
