package usecase

import (
	"encoding/json"
	"time"
	"vacantr/internal/adapter/cache"
	"vacantr/internal/core"
)

func CacheVacancies(vacancies []core.Vacancy) {
	data, _ := json.Marshal(vacancies)
	cache.Redis.Set(cache.Ctx, "latest_vacancies", data, 5*time.Minute)
}

func GetCachedVacancies() []core.Vacancy {
	var vacancies []core.Vacancy

	val, err := cache.Redis.Get(cache.Ctx, "latest_vacancies").Result()
	if err != nil {
		json.Unmarshal([]byte(val), &vacancies)
	}
	return vacancies
}
