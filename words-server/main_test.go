package main

import (
	"testing"

	"example.com/Sinezx/words-server/util"
)

func TestConfig(t *testing.T) {
	util.InitConfig()
	if util.Config.Dsn != "host=192.168.2.1 user=postgres password=postgres dbname=db port=5432 sslmode=disable TimeZone=Asia/Shanghai" {
		t.Errorf("dsn: %s", util.Config.Dsn)
	}
	if util.Config.AlarmLine != 30 {
		t.Errorf("alarmLine: %d", util.Config.AlarmLine)
	}
}
