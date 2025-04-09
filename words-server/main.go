package main

import (
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
	server.Run()
}
