package core

type User struct {
	ID         int64  `db:"id"`
	TelegramID int64  `db:"telegram_id"`
	Username   string `db:"username"`
}
