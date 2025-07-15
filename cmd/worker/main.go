package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"gopkg.in/telebot.v3"
	"log"
	"os"
	"time"
	"vacantr/internal/adapter/storage/postgres"
	"vacantr/internal/core"
)

func main() {
	pref := telebot.Settings{
		Token:  os.Getenv("TELEGRAM_TOKEN"),
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}
	bot, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{os.Getenv("KAFKA_ADDR")},
		Topic:   "vacancy_created",
		GroupID: "vacancy-consumer-group",
	})
	defer r.Close()

	db := postgres.Connect()

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Println("kafka read error", err)
		}

		var vac core.Vacancy
		if err := json.Unmarshal(m.Value, &vac); err != nil {
			log.Println("json unmarshal error", err)
			continue
		}

		subscribers := postgres.GetSubscribers(db)
		for _, userID := range subscribers {
			bot.Send(&telebot.User{ID: userID}, fmt.Sprintf("%s\n%s", vac.Title, vac.URL))
			postgres.MarkVacancySeen(db, userID, vac.ID)
		}
	}
}
