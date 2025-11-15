package db

import (
	"time"

	"example.com/Sinezx/words-server/util"
	"gorm.io/gorm"
)

type UserWord struct {
	gorm.Model
	UserId   uint      `json:"user_id"`
	WordId   uint      `json:"word_id"`
	Rate     float64   `json:"rate"`
	RateUpAt time.Time `json:"rate_up_at"`
}

func QueryUserWordsByUserId(userId uint, offset int, limit int) (int64, []UserWord, error) {
	var userwords []UserWord
	result := gormDB.Where("user_id = ?", userId).Order("rate desc").Offset(offset).Limit(limit).Find(&userwords)
	return result.RowsAffected, userwords, result.Error
}

func InsertUserWord(w *UserWord) (uint, error) {
	w.RateUpAt = time.Now().UTC()
	w.Rate = util.TheEbbinghausForgettingCurve(float64(1) / 60)
	result := gormDB.Create(&w)
	return w.ID, result.Error
}

func UpdateUserWordRate(id uint) error {
	userWord := UserWord{Rate: util.TheEbbinghausForgettingCurve(float64(1) / 60), RateUpAt: time.Now().UTC()}
	userWord.ID = id
	result := gormDB.Model(&userWord).Updates(&userWord)
	// result := db.Model(&word).Update("rate", word.Rate)
	return result.Error
}

func HardDeleteUserWord(id uint) (int64, error) {
	word := UserWord{}
	word.ID = id
	result := gormDB.Unscoped().Delete(&word)
	return result.RowsAffected, result.Error
}
