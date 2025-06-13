package db

import (
	"time"

	"example.com/Sinezx/words-server/util"
)

func queryAllWordUpdatedAt() (int64, []Word, error) {
	var words []Word
	result := gormDB.Select("id", "rate_up_at").Find(&words)
	if result.Error == nil {
		return result.RowsAffected, words, nil
	} else {
		return 0, nil, result.Error
	}
}

func updateWordRate(word *Word) {
	gormDB.Model(&word).Updates(map[string]any{"rate": word.Rate})
	// db.Exec("UPDATE words SET rate = ? WHERE id = ?", word.Rate, word.ID)
}

func calculateWordRate(w *Word, t *time.Time) {
	sub := t.Sub(w.RateUpAt)
	// rate set zero when update time more than alarmhours
	if sub.Hours() > util.AlarmHours {
		w.Rate = 0
	} else {
		w.Rate = util.TheEbbinghausForgettingCurve(sub.Hours())
	}
}

func UpdateWordSchedule(ticker *time.Ticker) {
	for {
		<-ticker.C
		// update all words rate
		total, words, err := queryAllWordUpdatedAt()
		t := time.Now().UTC()
		if err == nil && total > 0 {
			for _, word := range words {
				calculateWordRate(&word, &t)
				updateWordRate(&word)
			}
		}
		util.Info("[schdule] update word")
	}
}
