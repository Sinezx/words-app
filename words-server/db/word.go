package db

import (
	"example.com/Sinezx/words-server/util"
	"gorm.io/gorm"
)

type Word struct {
	gorm.Model
	SourceText string `json:"source_text"`
	TargetText string `json:"target_text"`
	VoicePath  string `json:"voice_path"`
}

func QueryWordById(id uint) (*Word, error) {
	var word Word
	result := gormDB.First(&word, id)
	return &word, result.Error
}

func InsertWord(w *Word) error {
	result := gormDB.Create(&w)
	return result.Error
}

func QueryWordId(sourcetext string) uint {
	var wordId uint
	result := gormDB.Raw("SELECT id FROM words WHERE source_text = ? LIMIT 1", sourcetext).Scan(&wordId)
	if result.Error == nil {
		return wordId
	} else {
		util.Info(result.Error.Error())
		return 0
	}
}

func HardDeleteWord(id uint) (int64, error) {
	word := Word{}
	word.ID = id
	result := gormDB.Unscoped().Delete(&word)
	return result.RowsAffected, result.Error
}
