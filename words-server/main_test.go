package main

import (
	"testing"
)

/*
func TestDB(t *testing.T) {
	dsn := "host=192.168.2.1 user=postgres password=root dbname=db port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db.Connt(dsn)
	util.Info("first query")
	db.Query()
	util.Info("insert one row")
	db.Insert(&db.Word{Subject: "afagsd", Translation: "sffdf", Rate: 4.364})
	util.Info("query after insert")
	db.Query()
	w := db.First()
	w.Rate = 72.9
	util.Info("update the row")
	db.Update(w)
	util.Info("query after update")
	db.Query()
	util.Info("delete the row")
	db.Delete(w.ID)
	util.Info("last query")
	db.Query()
}
*/

func TestMain(t *testing.T) {
	// dsn := "host=192.168.2.1 user=postgres password=root dbname=db port=5432 sslmode=disable"
	// db.Connt(dsn)
	// tim := time.Now().UTC()
	// word, _ := db.QueryById(26)
	// fmt.Println(word.UpdatedAt)
	// fmt.Println(tim)
}
