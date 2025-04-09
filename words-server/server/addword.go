package server

import (
	"example.com/Sinezx/words-server/db"
	"example.com/Sinezx/words-server/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AddWord struct {
	SourceText string `json:"source_text"`
	TargetText string `json:"target_text"`
}

func addword(c *gin.Context) {
	session := sessions.Default(c)
	var addWord AddWord
	c.BindJSON(&addWord)
	var word db.Word
	word.UserId = session.Get(util.SessionUserIdKey).(int)
	word.SourceText = addWord.SourceText
	word.TargetText = addWord.TargetText
	id, err := db.InsertWord(&word)
	if err == nil {
		util.InfoFormat("[session:%s]->word insert success, id: %d", session.ID(), id)
	} else {
		util.InfoFormat("[session:%s]->word insert fail: %s", session.ID(), err.Error())
	}
	if err == nil {
		StatusOK(c, &gin.H{
			"id": id,
		})
	} else {
		StatusBadRequest(c, &gin.H{
			"message": err.Error(),
		})
	}
}
