package postgres

import (
	"github.com/jmoiron/sqlx"
	"log"
	"vacantr/internal/core"
)

func SaveUser(db *sqlx.DB, u core.User) {
	_, err := db.Exec("INSERT INTO users (telegram_id, username) VALUES ($1, $2) ON CONFLICT (telegram_id) DO NOTHING", u.TelegramID, u.Username)
	if err != nil {
		log.Println("insert user error:", err)
	}
}

func GetSubscribers(db *sqlx.DB) []int64 {
	var ids []int64

	err := db.Select(&ids, `SELECT user_id FROM subscriptions`)
	if err != nil {
		log.Println("get subscribers error:", err)
	}

	return ids
}
