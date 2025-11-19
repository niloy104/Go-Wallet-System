package domain

type Transaction struct {
	ID         uint64  `json:"id" db:"id"`
	FromUserID *uint64 `json:"from_user_id,omitempty" db:"from_user_id"`
	ToUserID   *uint64 `json:"to_user_id,omitempty" db:"to_user_id"`
	Amount     int64   `json:"amount" db:"amount"`
	Type       string  `json:"type" db:"type"`
	Status     string  `json:"status" db:"status"`
	CreatedAt  string  `json:"created_at" db:"created_at"`
}
