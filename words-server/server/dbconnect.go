package server

import (
	"context"
	"strings"
	"time"

	"example.com/Sinezx/words-server/db"
	"example.com/Sinezx/words-server/util"
	"github.com/gin-gonic/gin"
)

type ConnectParams struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
	Port     string `json:"port"`
	Conf     string `json:"conf"`
}

func dbconnect(c *gin.Context) {
	session := util.GetSession(c)
	connectParams := ConnectParams{}
	c.BindJSON(&connectParams)
	var builder strings.Builder
	builder.WriteString("host=")
	builder.WriteString(connectParams.Host)
	builder.WriteString(" user=")
	builder.WriteString(connectParams.User)
	builder.WriteString(" password=")
	builder.WriteString(connectParams.Password)
	builder.WriteString(" dbname=")
	builder.WriteString(connectParams.Dbname)
	builder.WriteString(" port=")
	builder.WriteString(connectParams.Port)
	builder.WriteByte(' ')
	builder.WriteString(connectParams.Conf)
	dsn := builder.String()
	util.InfoFormat("[session:%s]->dsn: %s", session.ID(), dsn)
	currentdb, err := db.Connt(dsn)
	if err == nil {

		// save info to session
		session.SetPoint(util.DataBase, currentdb)
		session.Set(util.DataBaseConnectFlag, true)
		session.Save()

		// schedule start
		ctx, cancel := context.WithCancel(context.Background())
		ticker := time.NewTicker(time.Hour)
		session.SetFunc(util.UpdateWordCancelFuncId, cancel)
		go db.UpdateWordSchedule(currentdb, ticker, &ctx)

		StatusOK(c, &gin.H{"message": "success"})
	} else {
		StatusOK(c, &gin.H{"message": "fail"})
	}
}

func disdbconnect(c *gin.Context) {
	clearSession(c)
	StatusOK(c, &gin.H{"message": "success"})
}

func clearSession(c *gin.Context) {
	session := util.GetSession(c)
	// kill schedule ticker goroutine
	cancel := session.GetFunc(util.UpdateWordCancelFuncId)
	if cancel != nil {
		cancel()
	}
	session.Clear()
	session.Save()
	util.InfoFormat("[session:%s]->disconnect", session.ID())
}
