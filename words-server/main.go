package main

import (
	"time"

	"example.com/Sinezx/words-server/db"
	"example.com/Sinezx/words-server/server"
	"example.com/Sinezx/words-server/util"
)

func main() {
	util.InitConfig()
	err := db.Connt(util.Config.Dsn)
	if err != nil {
		util.Info("db connect failed")
	}
	ticker := time.NewTicker(time.Hour)
	go db.UpdateWordSchedule(ticker)
	err = server.Run()
	if err != nil {
		util.Info("server start failed")
	}
	util.Info("server start success")
}
