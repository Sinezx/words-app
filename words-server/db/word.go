package db

import (
	"time"

	"example.com/Sinezx/words-server/util"
	"gorm.io/gorm"
)

type Word struct {
	gorm.Model
	UserId     uint      `json:"user_id"`
	SourceText string    `json:"source_text"`
	TargetText string    `json:"target_text"`
	Rate       float64   `json:"rate"`
	RateUpAt   time.Time `json:"rate_up_at"`
}

func QueryWordsByUserId(userId uint, offset int, limit int) (int64, []Word, error) {
	var words []Word
	result := gormDB.Where("user_id = ?", userId).Order("rate desc").Offset(offset).Limit(limit).Find(&words)
	return result.RowsAffected, words, result.Error
}

func QueryById(id uint) (*Word, error) {
	var word Word
	result := gormDB.First(&word, id)
	return &word, result.Error
}

func InsertWord(w *Word) (uint, error) {
	w.RateUpAt = time.Now().UTC()
	w.Rate = util.TheEbbinghausForgettingCurve(float64(1) / 60)
	result := gormDB.Create(&w)
	return w.ID, result.Error
}

func UpdateWordRate(id uint) error {
	word := Word{Rate: util.TheEbbinghausForgettingCurve(float64(1) / 60), RateUpAt: time.Now().UTC()}
	word.ID = id
	result := gormDB.Model(&word).Updates(&word)
	// result := db.Model(&word).Update("rate", word.Rate)
	return result.Error
}

func HardDeleteWord(id uint) (int64, error) {
	word := Word{}
	word.ID = id
	result := gormDB.Unscoped().Delete(&word)
	return result.RowsAffected, result.Error
}
