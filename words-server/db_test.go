package main

import (
	"testing"

	"example.com/Sinezx/words-server/db"
	"example.com/Sinezx/words-server/util"
)

func TestUserTable(t *testing.T) {
	util.InitConfig()
	db.Connt(util.Config.Dsn, "postgres")
	account := "tester"
	password := "tester"
	// insert user that account is tester
	_, err := db.InsertUser(account, password)
	if err != nil {
		t.Error(err.Error())
		return
	}
	// get user that account is tester
	user, err := db.QueryUserByAccount(account)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if user.Password != util.Md5(password) {
		t.Errorf("expectation:%s, actual:%s", util.Md5(password), user.Password)
	}
	// completely delete user that account is tester
	rowCount, err := db.HardDeleteUser(user.ID)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if rowCount != 1 {
		t.Errorf("expectation:1, actual:%d", rowCount)
	}
}

func TestWordTable(t *testing.T) {
	util.InitConfig()
	db.Connt(util.Config.Dsn, "postgres")
	word := db.Word{UserId: 1, SourceText: "abandon", TargetText: "放弃"}
	word_id, err := db.InsertWord(&word)
	if err != nil {
		t.Error(err.Error())
		return
	}
	// get db's word to compare target word
	dbWord, err := db.QueryById(word_id)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if !wordCompare(&word, dbWord) {
		t.Errorf("expectation:%s, actual:%s", util.JsonString(word), util.JsonString(*dbWord))
	}
	// get db's word by userId
	total, dbWords, err := db.QueryWordsByUserId(word.UserId, 0, 10)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if total != 1 || !wordCompare(&word, &dbWords[0]) {
		t.Errorf("[total] expectation:1, actual:%d", total)
		t.Errorf("[dbWords] actual:%s", util.JsonString(dbWords))
	}
	// update word's rate
	err = db.UpdateWordRate(word_id)
	if err != nil {
		t.Error(err.Error())
		return
	}
	// completely delete word that sourcetext is abandon
	rows, err := db.HardDeleteWord(word_id)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if rows != 1 {
		t.Errorf("[total] expectation:1, actual:%d", rows)
	}
}

func wordCompare(a, b *db.Word) bool {
	return a.UserId == b.UserId && a.SourceText == b.SourceText && a.TargetText == b.TargetText
}
