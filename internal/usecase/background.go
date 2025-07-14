package usecase

import "time"

func StartBackgroundParser(vacancyUC *VacancyUseCase) {
	ticker := time.NewTicker(5 * time.Minute)
	go func() {
		for {
			<-ticker.C
			vacancyUC.GetTopVacancies()
		}
	}()
}
