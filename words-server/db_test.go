package main

import (
	"testing"

	"example.com/Sinezx/words-server/db"
	"example.com/Sinezx/words-server/util"
)

func TestUserTable(t *testing.T) {
	util.InitConfig()
	// db.Connt(util.Config.Dsn, "postgres")
	db.Connt(util.Config.DsnSQLite, "sqlite")
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
	// db.Connt(util.Config.Dsn, "postgres")
	db.Connt(util.Config.DsnSQLite, "sqlite")
	word := db.Word{SourceText: "something", TargetText: "anything"}
	db.InsertWord(&word)
	userWord := db.UserWord{UserId: 0, WordId: word.ID}
	uwid, err := db.InsertUserWord(&userWord)
	if err != nil {
		t.Error(err.Error())
		return
	}
	// get db's word by userId
	total, dbUserWords, err := db.QueryUserWordsByUserId(userWord.UserId, 0, 10)
	if err != nil {
		t.Error(err.Error())
		return
	}
	dbWords, err := db.QuerWordById(dbUserWords[0].WordId)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if total != 1 || !wordCompare(&word, dbWords) {
		t.Errorf("[total] expectation:1, actual:%d", total)
		t.Errorf("[dbWords] actual:%s", util.JsonString(dbWords))
	}
	// update word's rate
	err = db.UpdateUserWordRate(uwid)
	if err != nil {
		t.Error(err.Error())
		return
	}
	// completely delete userword
	rows, err := db.HardDeleteUserWord(uwid)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if rows != 1 {
		t.Errorf("[total] expectation:1, actual:%d", rows)
	}
	// completely delete word
	rows, err = db.HardDeleteWord(uwid)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if rows != 1 {
		t.Errorf("[total] expectation:1, actual:%d", rows)
	}

}

func wordCompare(a, b *db.Word) bool {
	return a.SourceText == b.SourceText && a.TargetText == b.TargetText
}
