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

var db *gorm.DB

func Connt(dsn string) error {
	util.InfoFormat("dsn: %s", dsn)
	if db == nil {
		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
			NowFunc: func() time.Time {
				return time.Now().UTC()
			},
		})
		if err != nil {
			db = nil
			return err
		}
		util.Info("db ready")
	} else {
		util.Info("db already exists")
	}
	return nil
}

func Clear() {
	db = nil
}

func Query(offset int, limit int) (int64, []Word, error) {
	var words []Word
	result := db.Order("rate desc").Offset(offset).Limit(limit).Find(&words)
	if result.Error == nil {
		util.InfoFormat("query Total: %d", result.RowsAffected)
		return result.RowsAffected, words, nil
	} else {
		return 0, nil, result.Error
	}
}

func QueryById(id int) (*Word, error) {
	var word Word
	result := db.First(&word, id)
	return &word, result.Error
}

func Insert(w *Word) (uint, error) {
	result := db.Create(&w)
	if result.Error == nil {
		util.InfoFormat("word insert success, id: %d", w.ID)
		return w.ID, nil
	} else {
		util.Info(result.Error.Error())
		return w.ID, result.Error
	}
}

func Delete(id uint) {
	// soft delete
	// db.Delete(&Word{}, id)
	db.Exec("DELETE FROM words WHERE id = ?", id)
}

func UpdateWordRate(id uint) (*Word, error) {
	word := Word{Rate: util.TheEbbinghausForgettingCurve(float64(1) / 60), RateUpAt: time.Now().UTC()}
	word.ID = id
	result := db.Model(&word).Updates(&word)
	// result := db.Model(&word).Update("rate", word.Rate)
	if result.Error == nil {
		util.InfoFormat("word: %d rate update", id)
		return &word, nil
	} else {
		util.Info(result.Error.Error())
		return &word, result.Error
	}
}

func Update(w *Word) {
	db.Save(&w)
}
