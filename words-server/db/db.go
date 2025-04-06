package db

import (
	"time"

	"example.com/Sinezx/words-server/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Word struct {
	gorm.Model
	Subject     string    `json:"subject"`
	Translation string    `json:"translation"`
	Rate        float64   `json:"rate"`
	RateUpAt    time.Time `json:"rate_up_at"`
}

func Connt(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		return nil, err
	}
	return db, err
}

func Query(db *gorm.DB, offset int, limit int) (int64, []Word, error) {
	var words []Word
	result := db.Order("rate desc").Offset(offset).Limit(limit).Find(&words)
	return result.RowsAffected, words, result.Error
}

func QueryById(db *gorm.DB, id int) (*Word, error) {
	var word Word
	result := db.First(&word, id)
	return &word, result.Error
}

func Insert(db *gorm.DB, w *Word) (uint, error) {
	result := db.Create(&w)
	return w.ID, result.Error
}

func UpdateWordRate(db *gorm.DB, id uint) (*Word, error) {
	word := Word{Rate: util.TheEbbinghausForgettingCurve(float64(1) / 60), RateUpAt: time.Now().UTC()}
	word.ID = id
	result := db.Model(&word).Updates(&word)
	// result := db.Model(&word).Update("rate", word.Rate)
	return &word, result.Error
}
