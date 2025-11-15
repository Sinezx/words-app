package db

import (
	"gorm.io/gorm"
)

type Word struct {
	gorm.Model
	SourceText string `json:"source_text"`
	TargetText string `json:"target_text"`
	VoicePath  string `json:"voice_path"`
}

func QuerWordById(id uint) (*Word, error) {
	var word Word
	result := gormDB.First(&word, id)
	return &word, result.Error
}

func InsertWord(w *Word) error {
	result := gormDB.Create(&w)
	return result.Error
}

func HardDeleteWord(id uint) (int64, error) {
	word := Word{}
	word.ID = id
	result := gormDB.Unscoped().Delete(&word)
	return result.RowsAffected, result.Error
}
