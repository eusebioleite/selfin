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
	CategoryID  int64   `json:"category_id"`
	UserID      int64   `json:"user_id"`
}
